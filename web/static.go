package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	var contentType string

	switch filepath.Ext(r.URL.Path)[1:] {
	case "css":
		contentType = "text/css"
	case "js":
		contentType = "application/javascript"
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}


	w.Header().Set("content-type", contentType + "; charset=utf-8")

	f, err := os.OpenFile("web" + r.URL.Path, os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := io.Copy(w, f); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
