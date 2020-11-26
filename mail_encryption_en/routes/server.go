package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Server() {

	router := mux.NewRouter()

	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("public/js/"))))

	router.HandleFunc("/", login).Methods("GET")
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/index", index).Methods("GET")
	router.HandleFunc("/join", join).Methods("GET")
	//router.HandleFunc("/inbox", inbox).Methods("GET")
	router.HandleFunc("/sent", sent).Methods("GET")
	//router.HandleFunc("/imp", imp).Methods("GET")
	router.HandleFunc("/spam", spam).Methods("GET")
	router.HandleFunc("/draft", draft).Methods("GET")
	router.HandleFunc("/trash", trash).Methods("GET")
	router.HandleFunc("/compose", compose).Methods("GET")
	router.HandleFunc("/compose", compose).Methods("post")
	router.HandleFunc("/logout", logout).Methods("GET")
	router.HandleFunc("/set", set).Methods("GET")
	router.HandleFunc("/mail_read", read).Methods("GET")
	router.HandleFunc("/mailChange", mailChange).Methods("GET")
	router.HandleFunc("/mailRead", mailRead).Methods("GET") // ?
	router.HandleFunc("/modMailServer", modMailServer).Methods("GET")
	router.HandleFunc("/pop", pop).Methods("GET")
	router.HandleFunc("/forgot", forgot).Methods("GET")

	router.HandleFunc("/deleteAccount", deleteAccount).Methods("POST")
	router.HandleFunc("/changeAccountPassword", changeAccountPassword).Methods("POST")
	router.HandleFunc("/openModalPwCheck", changeAccountPasswordCheck).Methods("POST")
	router.HandleFunc("/idFind", idFind).Methods("POST")
	router.HandleFunc("/pwFind", pwFind).Methods("POST")
	router.HandleFunc("/mailDelete", mailDelete).Methods("POST")
	router.HandleFunc("/defaultmailChange", defaultmailChange).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/mailPreset", mailPreset).Methods("POST")
	router.HandleFunc("/mailserverUpdate", mailserverUpdate).Methods("POST")
	router.HandleFunc("/mailserverDelete", mailserverDelete).Methods("POST")
	router.HandleFunc("/mailserverInsert", mailserverInsert).Methods("POST")
	router.HandleFunc("/index", index).Methods("POST")
	router.HandleFunc("/mailList", mailList).Methods("POST")
	router.HandleFunc("/loginTest", loginTest).Methods("POST")
	router.HandleFunc("/write", write).Methods("POST") //메일 쓰기 - 보내기 버튼

	server := negroni.Classic() // router 변수를 모니터링 하면서 로그를 출력할 middle ware
	server.UseHandler(router)   // negroni 패키지 사용 router가 실행 될때 로그를 가져오는 server 변수 선언

	PORT := os.Getenv("SERVER_PORT")
	// fmt.Println("PORT: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, server))
	//C:\Windows\System32\drivers\etc
	// openssl 사용 root ca 발급 -> 보안 인증 브라우저에서 안됨
	//log.Fatal(http.ListenAndServeTLS(":"+PORT, "server.crt", "server.key", server))
}
