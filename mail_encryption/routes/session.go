package routes

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

// GetSession 세션 추출
func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session-jintechemail")
	if err != nil {
		//panic(err)
	}
	return session
}

// AllSessions 겟 세션
func AllSessions(w http.ResponseWriter, r *http.Request) (interface{}, interface{}) {
	session := GetSession(w, r)
	id := session.Values["id"]
	eId := session.Values["eId"]

	return id, eId
}

// LoggedIn 로그인 되었는지 판단
func LoggedIn(w http.ResponseWriter, r *http.Request, urlRedirect string) {
	var URL string
	if urlRedirect == "" {
		URL = "/login"
	} else {
		URL = urlRedirect
	}
	id, _ := AllSessions(w, r)

	if id == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

// NotLoggedIn 로그인이 안되었는지 확인
func NotLoggedIn(w http.ResponseWriter, r *http.Request) {
	id, _ := AllSessions(w, r)
	if id != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func Invalid(w http.ResponseWriter, r *http.Request, what int) {
	if what == 0 {
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}
