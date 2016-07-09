package image

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"time"

	"github.com/filiptc/gorbit/config"
	"github.com/golang/freetype"
)

func MergeOverlay(frame []byte, conf *config.Config) ([]byte, error) {
	r := bytes.NewReader(frame)
	frameImage, err := jpeg.Decode(r)
	if err != nil {
		return []byte{}, err
	}
	composite, err := getOverlay(frameImage, conf)
	if err != nil {
		return []byte{}, err
	}
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, composite, nil); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil

}

func getOverlay(img image.Image, conf *config.Config) (image.Image, error) {
	f, err := freetype.ParseFont(conf.Font)
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(12)
	c.SetClip(b)
	c.SetDst(m)
	c.SetSrc(image.White)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(12)>>6))
	_, err = c.DrawString(time.Now().Format("15:04:05"), pt)
	if err != nil {
		return nil, err
	}
	return m, nil
}
