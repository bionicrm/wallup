package web

import (
	"io"
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

	f, err := os.Open("web" + r.URL.Path)
	if err != nil {
		WriteISE(err, w)
		return
	}
	defer f.Close()

	if _, err := io.Copy(w, f); err != nil {
		WriteISE(err, w)
		return
	}
}
