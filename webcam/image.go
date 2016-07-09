package webcam

import (
	"time"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/blackjack/webcam"
	"github.com/filiptc/gorbit/config"
	"gopkg.in/klaidliadon/console.v1"
)

type WebCam struct {
	o    *observable.Observable
	cs   *console.Console
	conf *config.Config
	Cam  *webcam.Webcam
}

func NewWebCam(o *observable.Observable, cs *console.Console, conf *config.Config) *WebCam {
	cam, err := webcam.Open(conf.Cam.Device)
	if err != nil {
		panic(err.Error())
	}
	return &WebCam{o, cs, conf, cam}
}

func (wc *WebCam) GetFormatSlice() []webcam.PixelFormat {

	format_desc := wc.Cam.GetSupportedFormats()
	var formats []webcam.PixelFormat
	for f := range format_desc {
		formats = append(formats, f)
	}
	return formats
}

func (wc *WebCam) InitStream() {
	defer wc.Cam.Close()
	wc.Cam.SetImageFormat(wc.GetFormatSlice()[0], wc.conf.Cam.Width, wc.conf.Cam.Height)

	err := wc.Cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}
	for {
		err = wc.Cam.WaitForFrame(uint32(time.Second))

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			wc.cs.Error("Video source timed out: %s", err)
			continue
		default:
			wc.cs.Error("Unhandled error: %s", err)
		}

		frame, err := wc.Cam.ReadFrame()
		if len(frame) != 0 {
			wc.o.Trigger("newFrame", frame)
		} else if err != nil {
			wc.cs.Error("Could not read frame: %s", err)
		}
	}
}
