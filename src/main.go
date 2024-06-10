package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/term"

	"main/auth"
	"main/filesys"
	"main/launch"
	"main/res"
	"main/signin"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	_, err := auth.LookupUser(r.Header.Get("Cookie"))
	if err != nil {
		w.Header().Set("Location", "/signin?redirect=" + url.QueryEscape(r.URL.String()))
		w.WriteHeader(302)
	} else if err == nil && r.URL.Path == "/" {
		w.Header().Set("Location", "/launch")
		w.WriteHeader(302)
	} else {
		w.WriteHeader(404)
		cmd := strings.TrimPrefix(r.URL.Path, "/")
		w.Write([]byte(cmd + ": app not found"))
	}
}

func main() {
	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username to sign in with [no default]: ")
	auth.User.Username, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	auth.User.Username = strings.TrimSuffix(auth.User.Username, "\n")
	auth.User.Username = strings.TrimSuffix(auth.User.Username, "\r")

	fmt.Print("Password to sign in with (will not echo) [no default]: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	auth.User.Password = string(password)

	mux := http.NewServeMux()
	mux.HandleFunc("/res/", res.Asset)
	mux.HandleFunc("/auth", auth.Handle)
	mux.HandleFunc("/fs/", filesys.Asset)

	mux.HandleFunc("/signin", signin.Main)
	mux.HandleFunc("/launch", launch.Main)
	mux.HandleFunc("/", redirect)

	fmt.Println("\nListening on localhost:2038...")
	err = http.ListenAndServe("localhost:2038", mux)
	if err != nil {
		panic(err)
	}
}
