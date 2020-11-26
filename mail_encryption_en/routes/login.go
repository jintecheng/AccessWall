package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func forgot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/forgot.html", "public/_head.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/login.html", "public/_head.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session := GetSession(w, r)
	session.Values["id"] = nil
	session.Values["eId"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func join(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/join.html", "public/_head.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func loginTest(w http.ResponseWriter, r *http.Request) {

	db := connectDB()
	defer disconnectDB(db)

	id := r.FormValue("id")
	pw := r.PostFormValue("pw")
	fmt.Println("id,pw: ", id, pw)

	if id == "" || pw == "" {
		w.WriteHeader(401)
	}

	var password string
	err := db.QueryRow("Select password FROM member WHERE id=$1", id).Scan(&password)

	if err != nil {
		fmt.Print(w, "wrong ID")
	}
	fmt.Println(password, " / ", pw)
	if EqualPassword(password, pw) == true {
		fmt.Print(w, "Login Success \n")
		session := GetSession(w, r)

		session.Values["id"] = id
		session.Save(r, w)
		var email string
		db.QueryRow("SELECT mail FROM mail_info WHERE id=$1 AND default_mail=true", id).Scan(&email)
		fmt.Println("email: ", email)
		session.Values["eId"] = email
		session.Save(r, w)

		fmt.Println("성공!")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

	} else {
		fmt.Print("wrong PW")
		http.Error(w, "Please send a request body", 400)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := r.FormValue("id")
	password1 := r.FormValue("password1")
	//password2 := r.FormValue("password2")

	email := r.FormValue("email")
	smtp_add := r.FormValue("smtp_add")
	smtp_port := r.FormValue("smtp_port")
	imap_add := r.FormValue("imap_add")
	imap_port := r.FormValue("imap_port")
	mail_passwd := r.FormValue("mail_passwd")
	ssl_tls_use := false
	default_mail := true

	db := connectDB()
	defer disconnectDB(db)

	var flag string
	db.QueryRow("SELECT id FROM member WHERE id=$1", id).Scan(&flag)

	if flag != "" {
		http.Error(w, "Already existed ID.", 400)
		return
	}

	// if name == "" || id == "" || email == "" || password1 == "" || password2 == "" {
	// 	log.Println("가입 정보를 모두 입력해주세요.")
	// 	http.Error(w, "Already existed ID.", 400)
	// 	return
	// }

	// if password1 != password2 {
	// 	log.Println("비밀번호가 다릅니다.")
	// 	http.Error(w, "Already existed ID.", 400)
	// 	return
	// }

	password := GeneratePassword(password1)
	_, err := db.Exec("INSERT INTO member VALUES ($1, $2, $3)", id, password, name)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "INSERT member error", 400)
		return
	}

	//db query: CREATE EXTENSION pgcrypto
	_, err = db.Exec("INSERT INTO mail_info(mail, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd, default_mail) VALUES($1, $2, $3, $4, $5, $6, $7, encode(encrypt(convert_to($8,'utf8'),'epjtwihnsasdc','aes'),'hex'), $9)",
		email, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd, default_mail)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "INSERT mail_info error", 403)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func idFind(w http.ResponseWriter, r *http.Request) {
	var id string
	var receivers []string

	type id3 struct {
		Name  string
		Email string
	}

	var u3 id3

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&u3)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("이름: ", u3.Name, " 이메일: ", u3.Email)

	db := connectDB()
	defer disconnectDB(db)

	db.QueryRow("SELECT member.id FROM mail_info, member WHERE member.name=$1 and mail_info.mail=$2", u3.Name, u3.Email).Scan(&id)

	if id == "" {
		http.Error(w, "Please check your name and email adress.", 400)
		return
	}

	receivers = append(receivers, u3.Email)

	auth := smtp.PlainAuth("", "jh5022@jintech2ng.co.kr", "Wkqtps!2", "smtp.Whoisg.kr")

	headers := make(map[string]string)
	headers["From"] = "jintech@jintecg.com"
	headers["To"] = u3.Email
	headers["Subject"] = "EES find ID"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + "Your ID: " + id + "\r\n"

	msg := []byte(message)

	smtp.SendMail("smtp.Whoisg.kr:587", auth, "jh5022@jintech2ng.co.kr", receivers, []byte(msg))
	w.WriteHeader(200)
}

func pwFind(w http.ResponseWriter, r *http.Request) {

	var id string
	var receivers []string

	name := r.FormValue("userName2")
	email := r.PostFormValue("userAddress2")

	fmt.Println("이름: ", name, " 이메일: ", email)

	db := connectDB()
	defer disconnectDB(db)

	newPw := newPassword(10)

	db.QueryRow("SELECT member.id FROM mail_info, member WHERE member.name=$1 and mail_info.mail=$2", name, email).Scan(&id)
	if id == "" {
		http.Error(w, "Please check your name and email adress.", 400)
		return
	}

	// newPw = GeneratePassword(newPw)
	if newPw == "Error" {
		http.Error(w, "Password Encryption Fail", 401)
	}

	_, err := db.Exec("UPDATE member SET password=$1 WHERE id=$2 ", GeneratePassword(newPw), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Update Error.", 401)
		return
	}

	receivers = append(receivers, email)

	auth := smtp.PlainAuth("", "jh5022@jintech2ng.co.kr", "Wkqtps!2", "smtp.Whoisg.kr")

	headers := make(map[string]string)
	headers["From"] = "jintech@jintecg.com"
	headers["To"] = email
	headers["Subject"] = "EES find Password"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + "Your new password: " + newPw + "\r\n"

	msg := []byte(message)

	smtp.SendMail("smtp.Whoisg.kr:587", auth, "jh5022@jintech2ng.co.kr", receivers, []byte(msg))
	w.WriteHeader(200)
}

const charset3 = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset3[seededRand.Intn(len(charset3))]
	}
	return string(b)
}

func newPassword(length int) string {
	return StringWithCharset(length, charset3)
}

func GeneratePassword(password string) string {
	pass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Password Encryption Fail", err)
		return "error"
	}
	return string(hashed)
}

func EqualPassword(hashedPassword string, password string) bool {
	fmt.Println(hashedPassword, " -- ", password)
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
