package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type JsonUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateResult struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/signup.html")
	t.Execute(w, nil)
}

func hqbLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Method error")
		return
	}

	err := r.ParseForm()

	if err != nil {
		fmt.Fprint(w, "Login error")
	}

	hlemail := r.PostFormValue("email")
	hlpassword := r.PostFormValue("password")

	rUser := User{
		Email:    hlemail,
		Password: hlpassword,
	}

	jsonBytes, err := json.Marshal(&rUser)

	if err != nil {
		fmt.Fprint(w, "json parse error")
	}

	_, _ = SendJsonPost(jsonBytes, "http://localhost:8001/api/user/login")
	http.Redirect(w, r, "/", http.StatusFound)
	/*
		if err != nil {
			panic(err)
		}
		fmt.Println(r.RemoteAddr)
		http.Post("http://localhost:8000/auth/sotlogin/sotauth", "application/json", bytes.NewBuffer(resultBytes))
	*/
}

//여기서 수신을 받아서 sotauth로 전송 하고 리턴 값을 기다린다.
func register(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("reg_user")
	email := r.PostFormValue("reg_email")
	password := r.PostFormValue("password")

	resp, err := http.PostForm("http://localhost:8001/api/user/new", url.Values{"name": {name}, "email": {email}, "password": {password}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
	/*
		if r.Method != "POST" {
			fmt.Fprint(w, "Method error")
		}

		err := r.ParseForm()
		if err != nil {
			fmt.Fprint(w, "Data error")
		}

		regUsername := r.PostFormValue("reg_user")
		regEmail := r.PostFormValue("reg_email")
		regPasswd := r.PostFormValue("passwordf")
		rUser := User{
			Name:     regUsername,
			Email:    regEmail,
			Password: regPasswd,
		}

		//rUser := JsonUser{regUsername, regEmail, regPasswd}

		jsonBytes, err := json.Marshal(&rUser)

		if err != nil {
			fmt.Fprint(w, "Json parse error")
		}

		resultBytes, err := SendJsonPost(jsonBytes, "http://localhost:8001/api/user/new")
	*/

	createResult := &CreateResult{}
	err = json.Unmarshal(respBody, &createResult)

	if err != nil {
		fmt.Fprint(w, err)
	}

	if createResult.Status {
		dir := strings.Split(email, "@")[0]
		createDir(dir)
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/signup", http.StatusFound)
	}
}
