package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func mailboxPointer(num string, mailadd string) string {

	mailPlatform := mailadd[strings.IndexByte(mailadd, '@')+1:]
	var selectedBox string

	if mailPlatform == "gmail.com" {
		switch num {
		case "1":
			selectedBox = "INBOX"
		case "2":
			selectedBox = "[Gmail]/보낸편지함"
		case "3":
			selectedBox = "[Gmail]/임시보관함"
		case "4":
			selectedBox = "[Gmail]/스팸함"
		case "5":
			selectedBox = "[Gmail]/휴지통"
		}
	} else if mailPlatform == "naver.com" {
		switch num {
		case "1":
			selectedBox = "INBOX"
		case "2":
			selectedBox = "Sent Messages"
		case "3":
			selectedBox = "임시보관함"
		case "4":
			selectedBox = "스팸메일함"
		case "5":
			selectedBox = "Deleted Messages"
		}
	} else if mailPlatform == "jintech2ng.co.kr" {
		switch num {
		case "1":
			selectedBox = "INBOX"
		case "2":
			selectedBox = "Sent"
		case "3":
			selectedBox = "Draft"
		case "4":
			selectedBox = "Spam"
		case "5":
			selectedBox = "Trash"
		}
	} else if mailPlatform == "daum.net" {
		switch num {
		case "1":
			selectedBox = "INBOX"
		case "2":
			selectedBox = "Sent Messages"
		case "3":
			selectedBox = "Drafts"
		case "4":
			selectedBox = "스팸편지함"
		case "5":
			selectedBox = "Deleted Messages"
		}
	}
	return selectedBox
}

func sent(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_sent.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html")
	if err != nil {
		panic(err)
	}

	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)
	fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(mailadd, "2"))
	} else {
		t.Execute(w, nil)
	}
}

func spam(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_spam.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html")
	if err != nil {
		panic(err)
	}
	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)
	fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(mailadd, "4"))
	} else {
		t.Execute(w, nil)
	}
}

func draft(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_draft.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html")
	if err != nil {
		panic(err)
	}
	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)
	fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(mailadd, "3"))
	} else {
		t.Execute(w, nil)
	}
}

func trash(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/menu_trash.html", "public/_head.html", "public/_menu.html", "public/_header.html", "public/_table.html")
	if err != nil {
		panic(err)
	}
	_, v1 := AllSessions(w, r)
	mailadd := fmt.Sprintf("%v", v1)
	//fmt.Println(":::::::", mailadd)
	if mailadd != "<nil>" && mailadd != "" && mailadd != "[]" {
		t.Execute(w, imapPage(mailadd, "5"))
	} else {
		t.Execute(w, nil)
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/setting.html", "public/_head.html", "public/_menu.html", "public/_header.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func read(w http.ResponseWriter, r *http.Request) {
	LoggedIn(w, r, "/login")
	t, err := template.ParseFiles("public/mail_read.html", "public/_head.html", "public/_menu.html", "public/_header.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
