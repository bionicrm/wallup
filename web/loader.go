package web

import (
	"html/template"
	"net/http"
)

type LayoutData struct {
	Title string
}

func WriteTempl(filename string, data interface{}, w http.ResponseWriter) error {
	t, err := template.ParseFiles(
		"web/" + filename,
		"web/footer.html",
		"web/header.html")
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}
