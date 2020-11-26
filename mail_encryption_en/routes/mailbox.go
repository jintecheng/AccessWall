package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func mailboxPointer(num string, mailadd string) string {
	mailPlatform := mailadd[strings.IndexByte(mailadd, '@')+1:]
	var selectedBox string
	var mailboxes string
	var arr []string
	n, _ := strconv.Atoi(num)

	//goDotEnvVariable(mailPlatform)

	mailboxes = goDotEnvVariable(mailPlatform)
	arr = strings.Split(mailboxes, ";")

	selectedBox = arr[n-1]
	fmt.Println("메일: ", selectedBox)
	return selectedBox
}

func sent(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_sent.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}

	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)
	//	fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(id, mailadd, "2"))
	} else {
		t.Execute(w, nil)
	}
}

func spam(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_spam.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)
	//	fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(id, mailadd, "4"))
	} else {
		t.Execute(w, nil)
	}
}

func draft(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_draft.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)
	//fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(id, mailadd, "3"))
	} else {
		t.Execute(w, nil)
	}
}

func trash(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_trash.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	v, v1 := AllSessions(w, r)
	id := fmt.Sprintf("%v", v)
	mailadd := fmt.Sprintf("%v", v1)
	//fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(id, mailadd, "5"))
	} else {
		t.Execute(w, nil)
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/setting.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func read(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/mail_read.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_headjs.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
