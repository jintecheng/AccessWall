package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey  = "oauth2_current_user"
	sessionDuration = time.Hour
)

type User struct {
	Uid      string    `json:"uid" bson:"uid"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	AvataUrl string    `json:"avatarurl" bson:"avataurl"`
	Password string    `json:"password" bson:"password"`
	Token    string    `json:"token" bson:"token"`
	Expired  time.Time `json:"expired" bson:"expired"`
}

func (u *User) Valid() bool {
	//현재시간 기준으로 만료시간 확인
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	//만료 시간 연장
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) *User {
	//세션에서 CurrentUser 정보를 가져옴
	s := sessions.GetSession(r)
	t := s.Get("sotauth_cookie")
	fmt.Println(t)

	if s.Get(currentUserKey) == nil {
		return nil
	}

	data := s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(data, &u)
	return &u
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		u.Refresh()
	}

	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
