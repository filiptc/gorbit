package webcam

import (
	"time"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/blackjack/webcam"
	"gopkg.in/klaidliadon/console.v1"
)

type WebCam struct {
	o  *observable.Observable
	cs *console.Console
}

func NewWebCam(o *observable.Observable, cs *console.Console) *WebCam {
	return &WebCam{o, cs}
}

func (wc *WebCam) InitStream() {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		panic(err.Error())
	}
	defer cam.Close()

	format_desc := cam.GetSupportedFormats()
	var formats []webcam.PixelFormat
	for f := range format_desc {
		formats = append(formats, f)
	}

	cam.SetImageFormat(formats[0], uint32(800), uint32(600))

	err = cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}
	for {
		err = cam.WaitForFrame(uint32(time.Second))

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			wc.cs.Error("Video source timed out: %s", err)
			continue
		default:
			wc.cs.Error("Unhandled error: %s", err)
		}

		frame, err := cam.ReadFrame()
		if len(frame) != 0 {
			wc.o.Trigger("newFrame", frame)
		} else if err != nil {
			wc.cs.Error("Could not read frame: %s", err)
		}
	}
}
