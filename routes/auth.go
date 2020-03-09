package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/urfave/negroni"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

const (
	nextPageKey     = "next_page"
	authSecurityKey = "auth_security_key"
)

type LoginResult struct {
	Account User   `json:"account"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func init() {
	gomniauth.SetSecurityKey(authSecurityKey)
	/*
	gomniauth.WithProviders(
		google.New("796891389935-g76o7dgg29gv1gq3o6eu8t3olm9o5g2s.apps.googleusercontent.com", "M-WJyrmgPqZ6EORfu7_evu8K", "http://127.0.0.1:8000/auth/callback/google"),
	)
	*/
		gomniauth.WithProviders(
			google.New("796891389935-3fim0c62nl1aup73s6tk0vgf3e3bt098.apps.googleusercontent.com", "ULi-40O_Lei7uiGJLxm87aOi", "http://aws.sot.is:8000/auth/callback/google"),
		)
}

func sotLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	action := params[2]
	provider := params[3]
	s := sessions.GetSession(r)

	switch action {
	case "login":
		if r.Method == "POST" {
			return
		}
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}

		loginUrl, err := p.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, loginUrl, http.StatusFound)
	case "callback":
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}

		creds, err := p.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln(err)
		}

		user, err := p.GetUser(creds)
		if err != nil {
			log.Fatalln(err)
		}

		u := &User{
			Uid:      user.Data().Get("id").MustStr(),
			Name:     user.Name(),
			Email:    user.Email(),
			AvataUrl: user.AvatarURL(),
		}

		gUser = u

		SetCurrentUser(r, u)
		// user email의 아이디로 디렉토리를 만듦
		userEmailId := strings.Split(u.Email, "@")
		createDir(userEmailId[0])

		// cookie 삭제후에 s.Get()이 널을 리턴 하는 버그를 수정
		var sessionPath string
		tInterface := s.Get(nextPageKey)
		if tInterface != nil {
			sessionPath = tInterface.(string)
		} else {
			sessionPath = "/"
		}

		if linkQuery != "" && gUser != nil {
			sessionPath = "/link?file=" + linkQuery
		}
		http.Redirect(w, r, sessionPath, http.StatusFound)
		/*
			case "sotlogin":
				defer r.Body.Close()

				body, _ := ioutil.ReadAll(r.Body)
				rrr := string(body)
				fmt.Println(rrr)
				c := &LoginResult{}

				err := json.Unmarshal(body, &c)
				if !c.Status || err != nil {
					http.Redirect(w, r, "/", http.StatusFound)
				}

				gUser = &c.Account
				userEmailId := strings.Split(gUser.Email, "@")
				createDir(userEmailId[0])

				var sessionPath string
				tInterface := s.Get(nextPageKey)
				if tInterface != nil {
					sessionPath = tInterface.(string)
				} else {
					sessionPath = "/?dir="
				}

				if linkQuery != "" && gUser != nil {
					sessionPath = "/link?file=" + linkQuery
				}
				co, err := r.Cookie("cookie-name")
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(co)
				http.Redirect(w, r, sessionPath, http.StatusFound)
		*/
	default:
		http.Error(w, "Auth action '"+action+"' is not supported", http.StatusNotFound)
	}
}

func LoginRequired(ignore ...string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		//링크 파일 다운로드를 위한 추가

		for _, s := range ignore {
			if strings.HasPrefix(r.URL.Path, s) {
				next(w, r)
				return
			}
		}

		if strings.HasPrefix(r.URL.Path, "/link") {
			linkQuery = r.URL.Query().Get("file")
		} else {
			linkQuery = ""
		}

		u := GetCurrentUser(r)

		cookie, err := r.Cookie("hqbfs-session")
		if cookie != nil && err == nil {
			fmt.Println("Hello")
		}

		if u != nil && u.Valid() {
			SetCurrentUser(r, u)
			next(w, r)
			return
		}

		SetCurrentUser(r, nil)

		sessions.GetSession(r).Set(nextPageKey, r.URL.RequestURI())
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}
	}

	return true, params

}
