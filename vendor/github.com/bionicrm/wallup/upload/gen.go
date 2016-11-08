package upload

import (
	"bytes"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
)

var boundsError = errors.New("new dimensions out of bounds")
var imgTypeError = errors.New("unsupported image type")

func generate(b []byte, x, y, widthL, widthR, height, gap int, scale float64) ([]byte, string, error) {
	// Decode.
	img, imgType, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, "", err
	}

	if err := check(img.Bounds(), x, y, widthL, widthR, height, gap, scale); err != nil {
		return nil, "", err
	}

	img = edit(img, x, y, widthL, widthR, height, gap, scale)

	return encode(img, imgType)
}

func check(bounds image.Rectangle, x, y, widthL, widthR, height, gap int, scale float64) error {
	widthL = int(float64(widthL) * scale)
	widthR = int(float64(widthR) * scale)
	height = int(float64(height) * scale)

	reqX := x + widthL + gap + widthR;
	reqY := y + height;

	if reqX > bounds.Max.X || reqY > bounds.Max.Y {
		return boundsError
	}

	return nil
}

func edit(src image.Image, x, y, widthL, widthR, height, gap int, scale float64) image.Image {
	widthLScaled := int(float64(widthL) * scale)
	widthRScaled := int(float64(widthR) * scale)
	heightScaled := int(float64(height) * scale)

	dst := image.NewRGBA(image.Rect(0, 0, widthLScaled + widthRScaled, heightScaled))

	dstLRect := image.Rect(0, 0, widthLScaled, heightScaled)
	dstRRect := image.Rect(widthLScaled, 0, widthLScaled + widthRScaled, heightScaled)

	srcLPt := image.Pt(x, y)
	srcRPt := image.Pt(x + widthLScaled + gap, y)

	draw.Draw(dst, dstLRect, src, srcLPt, draw.Src)
	draw.Draw(dst, dstRRect, src, srcRPt, draw.Src)

	if (scale != 1) {
		// Scale back down.
		return resize.Resize(uint(widthL + widthR), uint(height), dst, resize.Bicubic)
	}

	return dst;
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
