package upload

import (
	"github.com/bionicrm/wallup/web"
	"log"
	"net/http"
)

type showData struct {
	Title string
}

func ShowUploadHandler(w http.ResponseWriter, r *http.Request) {
	data := showData{
		Title: "Wallup",
	}

	if err := web.WriteTempl("upload/show.html", data, w); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DoUploadHandler(w http.ResponseWriter, r *http.Request) {

}
