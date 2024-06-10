package signin

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"main/auth"
)

func Main(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open("../res/html/signin.html")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = auth.LookupUser(r.Header.Get("Cookie"))
	if err == nil {
		path := r.URL.Query().Get("redirect")
		if !strings.HasPrefix(path, "/") {
			path = "/launch"
		}
		w.Header().Set("Location", path)
		w.WriteHeader(302)
		return
	}
	io.Copy(w, htmlFile)
	if err != nil {
		log.Println(err)
		return
	}
}
