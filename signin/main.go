package signin

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Main(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open("signin/main.html")
	if err != nil {
		log.Println(err)
	}
	io.Copy(w, htmlFile)
	if err != nil {
		log.Println(err)
	}
}
