package launch

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"main/auth"
)

func Main(w http.ResponseWriter, r *http.Request) {
	_, err := auth.LookupUser(r.Header.Get("Cookie"))
	if err != nil {
		w.Header().Set("Location", "/signin?redirect=" + url.QueryEscape(r.URL.String()))
		w.WriteHeader(302)
		return
	}
	htmlFile, err := os.Open("../res/html/launch.html")
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(w, htmlFile)
	if err != nil {
		log.Println(err)
		return
	}
}
