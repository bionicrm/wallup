package web

import (
	"html/template"
	"net/http"
)

func WriteTempl(filename string, data interface{}, w http.ResponseWriter) error {
	t, err := template.ParseFiles(
		"web/" + filename,
		"web/footer.html",
		"web/header.html")
	if err != nil {
		return err
	}

	if t.Execute(w, data) != nil {
		return err
	}

	return nil
}
