package web

import (
	"log"
	"net/http"
)

func WriteISE(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}

type badReqData struct {
	LayoutData

	Error string
}

func WriteBadReq(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	data := badReqData{
		LayoutData: LayoutData{
			Title: "Wallup - Error",
		},

		Error: err.Error(),
	}

	if err := WriteTempl("error.html", data, w); err != nil {
		WriteISE(err, w)
		return
	}
}
