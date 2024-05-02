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
	file, err := os.Open("../res/" + path)
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(w, file)
	if err != nil {
		log.Println(err)
		return
	}
}
