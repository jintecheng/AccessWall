package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//설정 이메일 서버 리스트
func mailPreset(w http.ResponseWriter, r *http.Request) {
	//LoggedIn(w, r, "/login")

	var mailAdd string
	var mailservers []byte

	db := connectDB()
	defer disconnectDB(db)

	rows, err := db.Query("SELECT mail_add FROM mail_preset")
	if err != nil {
		log.Panic("mailPreset: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&mailAdd)
		if err != nil {
			log.Panic("mailPreset: ", err)
		}
		b := []byte(mailAdd + ";")
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
	fmt.Println("mailservers: ", mailservers)
	w.Write(mailservers)
}

//추가 버튼
func mailserverInsert(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer disconnectDB(db)

	v1, _ := AllSessions(w, r)
	id := fmt.Sprintf("%v", v1)
	//id = id[1 : len(id)-1]

	mail := r.FormValue("mail")
	smtp_add := r.FormValue("smtp_add")
	smtp_port := r.FormValue("smtp_port")
	imap_add := r.FormValue("imap_add")
	imap_port := r.FormValue("imap_port")
	mail_passwd := r.FormValue("mail_passwd")
	ssl_tls_use := false

	//imap_port1, _ := strconv.Atoi(imap_port)
	fmt.Println("1", id, "7")
	fmt.Println("1", mail, "7")
	fmt.Println("2", smtp_add, "7")
	fmt.Println("3", smtp_port, "7")
	fmt.Println("4", imap_add, "7")
	fmt.Println("5", imap_port, "7")
	fmt.Println("6", mail, "7")
	fmt.Println("7", mail_passwd, "7")

	_, err := db.Exec("INSERT INTO mail_info(mail, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd,default_mail) VALUES($1, $2, $3, $4, $5, $6, $7, encode(encrypt(convert_to($8,'utf8'),'epjtwihnsasdc','aes'),'hex'),$9)",
		mail, id, smtp_add, smtp_port, imap_add, imap_port, ssl_tls_use, mail_passwd, false)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/set", http.StatusFound)
}

//삭제 버튼
func mailserverDelete(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer disconnectDB(db)

	v1, _ := AllSessions(w, r)
	id := fmt.Sprintf("%v", v1)
	mail := r.FormValue("mail")

	_, err := db.Exec("DELETE FROM mail_info WHERE mail=$1 AND id=$2", mail, id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/set", http.StatusFound)
}

// 설정 - 이메일 리스트 선택
func modMailServer(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer disconnectDB(db)

	mail, _ := r.URL.Query()["mail"]
	chk, _ := r.URL.Query()["chk"]
	email := mail[0]
	if chk != nil { //이메일 서버 셀렉트 선택
		check := chk[0]
		info := getPresetInfo(db, email)
		data := make(map[string]interface{})
		fmt.Println("email|", email, "|chk|", check, "|")
		data["smtp_add"] = info.smtp_add
		data["smtp_port"] = info.smtp_port

		if check == "pop" {
			data["pop_add"] = info.pop3_add
			data["pop_port"] = info.pop3_port

		} else if check == "imap" {
			data["imap_add"] = info.imap_add
			data["imap_port"] = info.imap_port

		}

		doc, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(doc)

	} else { // 이메일 리스트 선택
		fmt.Println(email)
		info := getAccountInfo(db, email)
		data := make(map[string]interface{})

		data["imap_add"] = info.imap_add
		data["imap_port"] = info.imap_port
		data["smtp_add"] = info.smtp_add
		data["smtp_port"] = info.smtp_port
		data["mail"] = email
		data["mail_passwd"] = info.mail_passwd

		doc, _ := json.Marshal(data)

		//modMail := []string{info.imap_add, info.imap_port, info.smtp_add, info.smtp_port, email, info.mail_passwd}

		w.Header().Set("Content-Type", "application/json")
		//	mod := []byte(modMail)
		w.Write(doc)
		//	json.NewEncoder(w).Encode(doc)
	}
}

//수정 버튼
func mailserverUpdate(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer disconnectDB(db)

	v1, _ := AllSessions(w, r)
	id := fmt.Sprintf("%v", v1)

	mail := r.FormValue("mail")
	tempmail := r.FormValue("tempmail")
	fmt.Println("temp", tempmail)
	smtp_add := r.FormValue("smtp_add")
	smtp_port := r.FormValue("smtp_port")
	imap_add := r.FormValue("imap_add")
	imap_port := r.FormValue("imap_port")
	mail_passwd := r.FormValue("mail_passwd")
	//ssl_tls_use := false

	_, err := db.Exec("UPDATE mail_info SET mail=$1, smtp_add=$2, smtp_port=$3, imap_add=$4, imap_port=$5, mail_passwd=encode(encrypt(convert_to($6,'utf8'),'epjtwihnsasdc','aes'),'hex') WHERE mail=$7 AND id=$8", mail, smtp_add, smtp_port, imap_add, imap_port, mail_passwd, tempmail, id)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/set", http.StatusFound)
}

func defaultmailChange(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer disconnectDB(db)

	v1, _ := AllSessions(w, r)
	id := fmt.Sprintf("%v", v1)
	mail := r.PostFormValue("opt")

	_, err := db.Exec("UPDATE mail_info SET default_mail=false WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("UPDATE mail_info SET default_mail=true WHERE id=$1 AND mail=$2", id, mail)
	if err != nil {
		panic(err)
	}

	session := GetSession(w, r)
	session.Values["eId"] = mail
	session.Save(r, w)

	http.Redirect(w, r, "/set", http.StatusFound)
}
