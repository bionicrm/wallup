package upload

import (
	"github.com/bionicrm/wallup/gen"
	"github.com/bionicrm/wallup/web"
	"io/ioutil"
	"net/http"
	"strconv"
)

type showData struct {
	web.LayoutData
}

func ShowUploadHandler(w http.ResponseWriter, r *http.Request) {
	data := showData{
		web.LayoutData{
			Title: "Wallup",
		},
	}

	if err := web.WriteTempl("upload/show.html", data, w); err != nil {
		web.WriteISE(err, w)
		return
	}
}

func DoUploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		web.WriteBadReq(err, w)
		return
	}
	defer file.Close()

	var (
		xS      = r.FormValue("x")
		yS      = r.FormValue("y")
		widthLS = r.FormValue("width-l")
		widthRS = r.FormValue("width-r")
		heightS = r.FormValue("height")
		gapS    = r.FormValue("gap")
	)

	x, err := strconv.Atoi(xS)
	y, err := strconv.Atoi(yS)
	widthL, err := strconv.Atoi(widthLS)
	widthR, err := strconv.Atoi(widthRS)
	height, err := strconv.Atoi(heightS)
	gap, err := strconv.Atoi(gapS)

	if err != nil {
		web.WriteBadReq(err, w)
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		web.WriteISE(err, w)
		return
	}

	img, contentType, err := gen.Generate(b, x, y, widthL, widthR, height, gap)
	if err != nil {
		web.WriteBadReq(err, w)
		return
	}

	w.Header().Set("content-type", contentType)
	w.Write(img)
}
