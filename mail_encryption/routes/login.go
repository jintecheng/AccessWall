package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
	//fmt.Println(":::::::", session.Values["id"])
	http.Redirect(w, r, "/login", http.StatusFound)
}

func join(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/join.html", "public/_head.html")
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
		http.Redirect(w, r, "/login", http.StatusFound)

	}

	var password string
	err := db.QueryRow("Select password FROM member WHERE id=$1", id).Scan(&password)

	if err != nil {
		fmt.Print(w, "wrong ID")
	}

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

		http.Redirect(w, r, "/index", http.StatusFound)

	} else {
		fmt.Print("wrong PW")
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func GeneratePassword(password string) (string, error) {
	pass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	return string(hashed), err
}

func EqualPassword(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := r.FormValue("id")
	password1 := r.FormValue("password1")
	password2 := r.FormValue("password2")

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
	db.QueryRow("SELECT id FROM members WHERE id=$1", id).Scan(&flag)

	if flag != "" {
		log.Println("이미 존재하는 아이디입니다.")
		http.Redirect(w, r, "/join", http.StatusFound)
		return
	}

	if name == "" || id == "" || email == "" || password1 == "" || password2 == "" {
		log.Println("가입 정보를 모두 입력해주세요.")
		http.Redirect(w, r, "/join", http.StatusFound)
		return
	}

	if password1 != password2 {
		log.Println("비밀번호가 다릅니다.")
		http.Redirect(w, r, "/join", http.StatusFound)
		return
	}

	password, _ := GeneratePassword(password1)

	_, err := db.Exec("INSERT INTO member VALUES ($1, $2, $3)", id, password, name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("INSERT INTO mail_info(mail, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd, default_mail) VALUES($1, $2, $3, $4, $5, $6, $7, encode(encrypt(convert_to($8,'utf8'),'epjtwihnsasdc','aes'),'hex'), $9)",
		email, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd, default_mail)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
