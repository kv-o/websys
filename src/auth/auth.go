package auth

import (
	"encoding/base64"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"strings"
	"time"
)

type user struct {
	Name     string
	Username string
	Password string
	Sessions []string
	Mutex    sync.Mutex
}

func (u *user) Lock() {
	u.Mutex.Lock()
}

func (u *user) Unlock() {
	u.Mutex.Unlock()
}

var User user

func LookupUser(cookie string) (user, error) {
	// TODO: replace token extraction with regexp code
	token := ""
	start := strings.Index(cookie, "token=")
	if start == -1 {
		return user{}, errors.New("invalid cookie")
	}
	start += 6
	end := strings.Index(cookie[start:], ";")
	if end == -1 {
		token = cookie[start:]
	} else {
		token = cookie[start:end]
	}
	for _, s := range User.Sessions {
		if s == token {
			return User, nil
		}
	}
	return user{}, errors.New("no user with given session token")
}

func expire(u *user, token string, expiry time.Time) {
	time.Sleep(time.Until(expiry))
	u.Lock()
	for i, v := range u.Sessions {
		if v == token {
			u.Sessions = append(u.Sessions[:i], u.Sessions[i+1:]...)
		}
	}
	u.Unlock()
}

func auth(query url.Values) (string, error) {
	username := query.Get("username")
	password := query.Get("password")
	if username != User.Username || password != User.Password {
		return "", errors.New("incorrect username or password")
	}

	buf := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())

	for i := range buf {
		buf[i] = byte(rand.Intn(255))
	}

	token := base64.StdEncoding.EncodeToString(buf)
	cookie := "token=" + token + "; Expires="
	expiry := time.Now().UTC().AddDate(0, 0, 1)
	cookie += expiry.Format(time.RFC1123)

	User.Lock()
	User.Sessions = append(User.Sessions, token)
	User.Unlock()

	go expire(&User, token, expiry)
	return cookie, nil
}

func Handle(w http.ResponseWriter, r *http.Request) {
	_, err := LookupUser(r.Header.Get("Cookie"))
	if err != nil {
		var cookie string
		err := r.ParseForm()
		if err == nil {
			cookie, err = auth(r.PostForm)
		}
		if err == nil {
			w.Header().Set("Set-Cookie", cookie)
			goto redirect
		} else {
			log.Println(err)
			w.Header().Set("Location", "/signin?auth=failed")
			w.WriteHeader(302)
		}
	}
redirect:
	path := r.URL.Query().Get("redirect")
	if !strings.HasPrefix(path, "/") {
		path = "/launch"
	}
	w.Header().Set("Location", path)
	w.WriteHeader(302)
}
