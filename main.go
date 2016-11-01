package main

import (
	"github.com/bionicrm/wallup/upload"
	"github.com/bionicrm/wallup/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",        upload.ShowUploadHandler)
	http.HandleFunc("/static/", web.HandleStatic)
	http.HandleFunc("/upload/", upload.DoUploadHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
