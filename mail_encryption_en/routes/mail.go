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

type Id2 struct {
	name  string
	email string
}

type Mail2 struct {
	Mto      string
	Mfrom    string
	Mtitle   string
	Mcontent template.HTML
	Mdate    string
	Mid      uint32
	Mnum     string
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

	bo := r.FormValue("bo")
	var reply Mail2
	reply.Mto = bo

	t, err := template.ParseFiles("public/mail_compose.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, reply)
}

func mailRead(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)

	if mailadd != "<nil>" {
		//t.Execute(w, imapPage(mailadd))
	} else {
		//t.Execute(w, nil)
	}
	uid1, ok := r.URL.Query()["uid"]
	num, ok := r.URL.Query()["num"] //페이지 번호
	fmt.Println(uid1[0], ok)

	uid, err := strconv.Atoi(uid1[0])

	m := imapRead(id, mailadd, uid, num[0])
	t, err := template.ParseFiles("public/mail_read.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	//w.Header
	t.Execute(w, m)
}

func imapRead(id string, s string, uid int, num string) Mail2 {

	var m2 Mail2
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	imap.CharsetReader = charset.Reader

	mailadd := s

	db := connectDB()
	defer disconnectDB(db)
	p := mailboxPointer(num, mailadd)
	accountinfo := getAccountInfo(db, mailadd, id)

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

	if uid == 0 {
		log.Panic("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)

	seqSet.AddNum((uint32)(uid))

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
		t = t.In(time.FixedZone("SST", 8*60*60)) // Singapol region
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
			b, _ := ioutil.ReadAll(p.Body)
			if string(b) != "" {
				emailbody = string(b)
			}
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Println("Got attachment: ", filename)
		}
	}
	m2.Mid = msg.SeqNum
	m2.Mnum = num
	m2.Mcontent = (template.HTML)(emailbody)
	return m2
}

func mailDelete(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	type a struct {
		Mail []string
		Num  string
	}
	var getMail a
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&getMail)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	len := len(getMail.Mail)
	fmt.Println("길이: ", len)
	fmt.Println("게시판번호: ", getMail.Num) //1:index 2:sent 3:spam 4:draft 5: trash
	fmt.Println("게시판번호: ", getMail.Mail)

	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)

	// Connect to db
	db := connectDB()
	defer disconnectDB(db)
	p := mailboxPointer(getMail.Num, mailadd)

	accountinfo := getAccountInfo(db, mailadd, id)
	hostaddr := (accountinfo.imap_add + ":" + accountinfo.imap_port)

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

	// We will delete the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqset := new(imap.SeqSet)

	for i := 0; i < len; i++ {
		num6, _ := strconv.ParseUint(getMail.Mail[i], 10, 32)
		num := uint32(num6)
		seqset.AddNum(num)

		// First mark the message as deleted
		item := imap.FormatFlagsOp(imap.AddFlags, true)
		flags := []interface{}{imap.DeletedFlag}
		if err := c.Store(seqset, item, flags, nil); err != nil {
			log.Fatal("2", err)
		}

		// Then delete it
		if err := c.Expunge(nil); err != nil {
			log.Fatal("3", err)
		}
	}

	log.Println("Checked message has been deleted")
}

func write(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")

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

	// DB 연동
	db := connectDB()
	defer disconnectDB(db)

	var (
		addr      string
		to        string
		receivers []string
	)
	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)
	//mailadd = mailadd[1 : len(mailadd)-1]
	accountinfo := getAccountInfo(db, mailadd, id)

	toMulti := strings.Split(u.Jto, ";")
	for i := range toMulti {
		to = toMulti[i]
		receivers = append(receivers, to)
	}

	auth := smtp.PlainAuth("", mailadd, accountinfo.mail_passwd, accountinfo.smtp_add)

	headers := make(map[string]string)
	headers["From"] = mailadd
	headers["To"] = u.Jto
	headers["Subject"] = u.Jtitle
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + u.Jinput + "\r\n"

	msg := []byte(message)
	addr = accountinfo.smtp_add + ":" + accountinfo.smtp_port
	err1 := smtp.SendMail(addr, auth, mailadd, receivers, msg)

	if err1 != nil {
		panic(err1)
	}
	fmt.Print("Successfully Sent Mail!")
}
