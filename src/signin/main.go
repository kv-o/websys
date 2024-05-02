package signin

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Main(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open("../res/html/ui.html")
	if err != nil {
		log.Println(err)
	}
	io.Copy(w, htmlFile)
	if err != nil {
		log.Println(err)
	}
}
