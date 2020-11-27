package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
)

func index(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_index.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}

	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)

	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(id, mailadd, "1"))
	} else {
		t.Execute(w, nil)
	}
}

func imapPage(id string, s string, num string) TodoPageData {
	log.Println("Connecting to server...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	imap.CharsetReader = charset.Reader
	mailadd := s

	// Connect to db
	db := connectDB()
	defer disconnectDB(db)

	accountinfo := getAccountInfo(db, mailadd, id)

	hostaddr := (accountinfo.imap_add + ":" + accountinfo.imap_port)
	log.Println("host: ", hostaddr)

	// Connect to server
	c, err := client.DialTLS(hostaddr, nil)
	if err != nil {
		log.Panic(err)

	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(mailadd, accountinfo.mail_passwd); err != nil {
		fmt.Println("로그인 에러")
		log.Panic(err)
		//로그인 에러
		//2020/06/10 10:21:48 Sorry, IMAP UserAuth Server Is Checking..Please, Try later.
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Panic("folder fetch error", err.Error())
	}
	log.Println("folder fetched...")
	p := mailboxPointer(num, mailadd)

	// Select INBOX
	mbox, err := c.Select(p, false)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	tonum := mbox.Messages // 메일 총개수
	if mbox.Messages > 50 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 49
	}
	//	fmt.Println("박스 개수: ", mbox.Messages)

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, tonum)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	//log.Println("Last 4 messages:")

	data := TodoPageData{
		Todos: make([]Todo, 55),
	}
	if tonum == 0 {
		return data
	}

	data.Num = num
	if num == "2" {
		data.Who = "To"
	} else {
		data.Who = "From"
	}
	var i = 0
	for msg := range messages {
		r := msg.GetBody(&section)
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Panic(err)
		}
		header := mr.Header

		if date, err := header.Date(); err == nil {
			//fmt.Println("date: ", date.UTC())
			t := date.UTC()
			t = t.In(time.FixedZone("SST", 8*60*60)) // 한국시간으로 바꾸기
			data.Todos[i].Date = (t.Format("2006-01-02 15:04:05"))

		}
		if num != "2" { //"_table" 보낸메일함은 받는 사람으로
			if from, err := header.AddressList("From"); err == nil {
				if from != nil {
					data.Todos[i].Sender = from[0].Address
				} else {
					data.Todos[i].Sender = "<null>"
				}
			}
		} else {
			if to, err := header.AddressList("To"); err == nil && len(to) != 0 { // to가 존재하지 않으면 오류 발생 - 제외
				//runtime error: index out of range [0] with length 0
				data.Todos[i].Sender = to[0].Address
			} else {
				data.Todos[i].Sender = ""
			}
		}

		if subject, err := header.Subject(); err == nil {
			if subject == "" {
				data.Todos[i].Title = "<제목없음>"
			} else {
				data.Todos[i].Title = subject
			}
		}
		data.Todos[i].Done = true
		data.Todos[i].Uid = msg.SeqNum
		//fmt.Println("UID: ", msg.SeqNum)
		i++
	}

	if err := <-done; err != nil {
		log.Panic(err)
	}
	log.Println("Done!")
	//메일 순서 변경
	for i, j := 0, len(data.Todos)-1; i < j; i, j = i+1, j-1 {
		data.Todos[i], data.Todos[j] = data.Todos[j], data.Todos[i]
	}
	return data
}

func mailList(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	var indivserver string
	var mailservers []byte
	var defaultMail string
	sessid1, _ := AllSessions(w, r)
	sessid := fmt.Sprintf("%v", sessid1)

	db := connectDB()
	defer disconnectDB(db)

	db.QueryRow("SELECT mail FROM mail_info WHERE id=$1 AND default_mail=true", sessid).Scan(&defaultMail)

	rows, err := db.Query("SELECT mail FROM mail_info WHERE id=$1 AND default_mail=false", sessid)
	if err != nil {
		log.Panic("mailList: ", err)
	}
	defer rows.Close()
	c := []byte(defaultMail + ";")
	mailservers = append(mailservers, c...)

	for rows.Next() {
		err := rows.Scan(&indivserver)
		if err != nil {
			log.Panic("mailList: ", err)
		}
		b := []byte(indivserver + ";")
		mailservers = append(mailservers, b...)
	}
	//	_, v1 := AllSessions(w, r)
	//c, ok := v1.(*[]byte)
	//	eid := fmt.Sprintf("%v", v1)
	//mailservers = append(mailservers, **c)
	//fmt.Print(mailservers)

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	//	b, _ := ioutil.ReadAll(r.Body)
	//s := string(b)
	//fmt.Println("받은 검색 내용: ", s)

	w.Write(mailservers)
}

func mailChange(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	email, _ := r.URL.Query()["id"]
	url, _ := r.URL.Query()["url"]
	url1 := strings.Join(url, "")
	url1 = "/" + url1
	//fmt.Println("url1: ", url1)
	//fmt.Println(email[0], ok)
	session := GetSession(w, r)

	session.Values["eId"] = email[0]
	session.Save(r, w)

	if strings.Contains(url1, "mailRead") {
		fmt.Println("in")
		w.Write([]byte("1"))
	}

	//fmt.Println("1: ", url)
	//w.Write([]byte("1"))
}
