package upload

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
)

var boundsError = errors.New("new dimensions out of bounds")
var imgTypeError = errors.New("unsupported image type")

func generate(b []byte, x, y, widthL, widthR, height, gap int) ([]byte, string, error) {
	// Decode.
	img, imgType, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, "", err
	}

	if err := check(img.Bounds(), x, y, widthL, widthR, height, gap); err != nil {
		return nil, "", err
	}

	img = edit(img, x, y, widthL, widthR, height, gap)

	return encode(img, imgType)
}

func check(bounds image.Rectangle, x, y, widthL, widthR, height, gap int) error {
	reqX := x + widthL + gap + widthR;
	reqY := y + height;

	if reqX > bounds.Max.X || reqY > bounds.Max.Y {
		return boundsError
	}

	return nil
}

func edit(src image.Image, x, y, widthL, widthR, height, gap int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, widthL + widthR, height))

	dstLRect := image.Rect(0, 0, widthL, height)
	dstRRect := image.Rect(widthL, 0, widthL + widthR, height)

	srcLPt := image.Pt(x, y)
	srcRPt := image.Pt(x + widthL + gap, y)

	draw.Draw(dst, dstLRect, src, srcLPt, draw.Src)
	draw.Draw(dst, dstRRect, src, srcRPt, draw.Src)

	return dst
}

func encode(img image.Image, imgType string) ([]byte, string, error) {
	var err error
	var buf bytes.Buffer

	switch imgType {
	case "gif":
		err = gif.Encode(&buf, img, &gif.Options{
			NumColors: 256,
		})
	case "jpeg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{
			Quality: 100,
		})
	case "png":
		err = png.Encode(&buf, img)
	default:
		return nil, "", imgTypeError
	}

	return buf.Bytes(), "image/" + imgType, err
}
