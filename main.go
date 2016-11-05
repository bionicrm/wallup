package main

import (
	"github.com/bionicrm/wallup/upload"
	"github.com/bionicrm/wallup/web"
	"log"
	"net/http"
)

const Port = "8080"

func main() {
	http.HandleFunc("/",        upload.ShowUploadHandler)
	http.HandleFunc("/static/", web.HandleStatic)
	http.HandleFunc("/upload",  upload.DoUploadHandler)

	log.Println("listening on port", Port)

	if err := http.ListenAndServe(":" + Port, nil); err != nil {
		log.Fatalln(err)
	}
}
