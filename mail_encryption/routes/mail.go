package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
)

type Data struct {
	Jto    string
	Jcc    string
	Jtitle string
	Jinput string
}

type Mail2 struct {
	Mto      string
	Mfrom    string
	Mtitle   string
	Mcontent template.HTML
	Mdate    string
}

type Todo struct {
	Title  string
	Sender string
	To     string
	Date   string
	Done   bool
	Uid    uint32
}

type TodoPageData struct {
	Who   string
	Num   string
	Todos []Todo
}

func compose(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	//_, v1 := AllSessions(w, r)
	//mailadd := fmt.Sprintf("%v", v1)
	//fmt.Println(mailadd)

	bo := r.FormValue("bo")
	var reply Mail2
	reply.Mto = bo

	t, err := template.ParseFiles("public/mail_compose.html", "public/_head.html", "public/_menu.html", "public/_header.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, reply)
}

func mailRead(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)

	fmt.Println("::::::::::", mailadd)
	if mailadd != "<nil>" {
		//t.Execute(w, imapPage(mailadd))
	} else {
		//t.Execute(w, nil)
	}
	uid1, ok := r.URL.Query()["uid"]
	num, ok := r.URL.Query()["num"]

	fmt.Println(uid1[0], ok)
	uid, err := strconv.Atoi(uid1[0])

	m := imapRead(mailadd, uid, num[0])
	t, err := template.ParseFiles("public/mail_read.html", "public/_head.html", "public/_menu.html", "public/_header.html")
	if err != nil {
		panic(err)
	}

	t.Execute(w, m)
}

func imapRead(s string, uid int, num string) Mail2 {

	var m2 Mail2
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	imap.CharsetReader = charset.Reader

	mailadd := s

	//[1 : len(s)-1]
	//log.Println(mailadd)

	// Connect to db
	db := connectDB()
	defer disconnectDB(db)
	p := mailboxPointer(num, mailadd)
	accountinfo := getAccountInfo(db, mailadd)

	hostaddr := (accountinfo.imap_add + ":" + accountinfo.imap_port)

	log.Println("host: ", hostaddr)

	c, err := client.DialTLS(hostaddr, nil)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(mailadd, accountinfo.mail_passwd); err != nil {
		log.Panic(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select(p, false)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("메일 개수: ", mbox.Messages)

	// Get the last message
	if uid == 0 {
		log.Panic("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)

	seqSet.AddNum((uint32)(uid))

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)

	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Panic(err)
			//Error in IMAP command FETCH: Invalid messageset
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Panic("Server didn't returned message")
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Panic("Server didn't returned message body")
	}
	// Create a new mail reader
	mr, err := mail.CreateReader(r)

	if err != nil {
		log.Panic(err, " Error Skip")
	}

	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		t := date.UTC()
		t = t.In(time.FixedZone("KST", 9*60*60)) // 한국시간으로 바꾸기
		m2.Mdate = (t.Format("2006-01-02 15:04:05"))
	}

	if num != "2" {
		if from, err := header.AddressList("From"); err == nil {
			m2.Mfrom = from[0].Address
			m2.Mto = "보낸 사람"
		}
	} else {
		if to, err := header.AddressList("To"); err == nil && len(to) != 0 {
			m2.Mfrom = to[0].Address
			m2.Mto = "받는 사람"
		}
	}

	if subject, err := header.Subject(); err == nil {
		m2.Mtitle = subject
	}

	var emailbody string
	// Process each message's part
	var i = 0

	for {

		i++
		p, err := mr.NextPart()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err, "Error Skip")
		}
		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			if string(b) != "" { //다음 내용이 빈칸인 경우 존재 제외
				emailbody = string(b)
				//fmt.Println("imapread 내용: ", emailbody)
			}
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Println("Got attachment: ", filename)
		}
	}

	m2.Mcontent = (template.HTML)(emailbody)
	return m2
}

func write(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	fmt.Println("------------------")

	var u Data
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("받는 사람: ", u.Jto)
	fmt.Println("참조: ", u.Jcc)
	fmt.Println("제목: ", u.Jtitle)
	fmt.Println("내용: ", u.Jinput)

	//defer http.Redirect(w, r, "/index", http.StatusMovedPermanently) // 페이지 이동 안됨

	// DB 연동		//enable
	db := connectDB()
	defer disconnectDB(db)

	var (
		addr      string
		to        string
		receivers []string
	)
	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)
	//mailadd = mailadd[1 : len(mailadd)-1]
	accountinfo := getAccountInfo(db, mailadd)

	toMulti := strings.Split(u.Jto, ";")
	for i := range toMulti {
		to = toMulti[i]
		receivers = append(receivers, to)
	}

	// to = u.Jto
	// fmt.Println("receivers: ", receivers)
	// receivers = append(receivers, to)

	auth := smtp.PlainAuth("", mailadd, accountinfo.mail_passwd, accountinfo.smtp_add)

	//메일 내용
	/*
		headerSubject := "Subject: " + u.Jtitle + "\n"

		headerBlank := "From" + sender + "\r\n"
			body := u.Jinput + "\r\n"

			fmt.Println("------------")
	*/
	//date := time.Now().Format("2006-01-02 15:04:05")
	headers := make(map[string]string)
	headers["From"] = mailadd
	headers["To"] = u.Jto
	headers["Subject"] = u.Jtitle
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	//headers["Date"] = date
	// headers["Time"] = date
	// fmt.Println("시간: ", date)
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + u.Jinput + "\r\n"

	msg := []byte(message)

	addr = accountinfo.smtp_add + ":" + accountinfo.smtp_port

	fmt.Println("accountinfo.mail_passwd: ", accountinfo.mail_passwd, "mailadd: ", mailadd, " accountinfo.smtp_add: ", accountinfo.smtp_add, "accountinfo.smtp_port: ", accountinfo.smtp_port)
	fmt.Println("7 addr: ", addr, "auth: ", auth, "mailadd: ", mailadd, "receivers: ", receivers, "msg: ", msg)
	err1 := smtp.SendMail(addr, auth, mailadd, receivers, msg)
	//[negroni] PANIC: read tcp 192.168.1.49:3815->211.231.108.14:465: wsarecv: An existing connection was forcibly closed by the remote host.
	fmt.Println("8")
	if err1 != nil {
		fmt.Println("9")
		panic(err1)
	}
	fmt.Println("10")
	fmt.Print("Successfully Sent Mail!")
}
