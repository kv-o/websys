package res

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Asset(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/res/")
	path = filepath.Clean(path)
	if strings.Index(path, "../") != -1 {
		log.Println("attempt to get files in parent folders")
		return
	}
	if strings.HasSuffix(path, ".html") {
		w.Header().Set(`Content-Type`, `text/html, charset="utf-8"`)
	} else if strings.HasSuffix(path, ".css") {
		w.Header().Set(`Content-Type`, `text/css, charset="utf-8"`)
	} else if strings.HasSuffix(path, ".js") {
		w.Header().Set(`Content-Type`, `text/javascript, charset="utf-8"`)
	} else if strings.HasSuffix(path, ".png") {
		w.Header().Set(`Content-Type`, `image/png`)
	} else if strings.HasSuffix(path, ".ttf") {
		w.Header().Set(`Content-Type`, `font/ttf`)
	}
	file, err := os.Open("../res/" + path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	io.Copy(w, file)
	if err != nil {
		log.Println(err)
		return
	}
}
