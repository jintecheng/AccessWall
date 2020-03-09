package routes

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/urfave/negroni"
)

const (
	sessionKey    = "hqbfs_session"
	sessionSecret = "hqbfs_session"
	rootDir       = "./Users"
)

func Server() {

	r := http.NewServeMux()
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", index)
	r.HandleFunc("/mail", mailPage)

	//user
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/hqblogin", hqbLogin)
	r.HandleFunc("/auth/", loginHandler)
	r.HandleFunc("/upload", upload)
	r.HandleFunc("/delete", delete)
	r.HandleFunc("/makedir", makeDir)
	r.HandleFunc("/register", register)

	//
	r.HandleFunc("/favicon.ico", favicon)

	//file
	r.HandleFunc("/filesperm", filesPerm)
	r.HandleFunc("/searchfile", searchFile)
	r.HandleFunc("/link", linkFile)

	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))
	n.Use(LoginRequired("/login", "/auth", "/signup", "/hqblogin", "/register"))
	n.UseHandler(r)
	n.Run(":8000")
}
