package routes

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/filiptc/gorbit/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/klaidliadon/console.v1"
)

const mjpeg_boundary = "frame"

type streamRoute struct {
	r    *gin.Engine
	conf *config.Config
	cs   *console.Console
	o    *observable.Observable
}

func newStream(r *gin.Engine, conf *config.Config, cs *console.Console, o *observable.Observable) Route {
	return &streamRoute{r, conf, cs, o}
}

func (r *streamRoute) Register() {
	r.r.GET("/stream", func(c *gin.Context) {
		defer r.cs.Debug("Client %s closed connection", c.Request.RemoteAddr)
		r.cs.Debug("New client %s", c.Request.RemoteAddr)

		contentType := fmt.Sprintf("multipart/x-mixed-replace;boundary=%s", mjpeg_boundary)
		c.Writer.Header().Add("Content-Type", contentType)

		frameQueue := make(chan []byte, 1)

		cb := func(frame []byte) { frameQueue <- frame }
		r.o.On("newFrame", cb)

		c.Stream(r.streamFrame(frameQueue, cb))
	})
}

func (r *streamRoute) streamFrame(frameQueue chan []byte, cb func(frame []byte)) func(w io.Writer) bool {
	return func(w io.Writer) bool {
		mimeWriter := multipart.NewWriter(w)
		if err := mimeWriter.SetBoundary(mjpeg_boundary); err != nil {
			r.cs.Error("Could not set boundary: %s", err)
			return false
		}

		return r.writeFrameOrQuit(w, cb, frameQueue, mimeWriter)

	}
}

func (r *streamRoute) writeFrameOrQuit(
	w io.Writer,
	cb func(frame []byte),
	frameQueue chan []byte,
	mw *multipart.Writer,
) bool {

	select {
	case <-w.(gin.ResponseWriter).CloseNotify():
		r.o.Off("newFrame", cb)
		return false
	case frame := <-frameQueue:
		/* removed for performance (this call lags 600ms on Rpi3)
		frame, err := image.MergeOverlay(frame, r.conf)
		if err != nil {
			r.cs.Error("Error decorating: %s", err)
			return false
		}
		*/
		if err := r.writeFrame(frame, mw); err != nil {
			r.cs.Error("Error writing frame into stream: %s", err)
			return false
		}
	}
	return true
}

func (r *streamRoute) writeFrame(frame []byte, mw *multipart.Writer) error {
	partHeader := make(textproto.MIMEHeader)
	partHeader.Add("Content-Type", "image/jpeg")
	partWriter, partErr := mw.CreatePart(partHeader)
	if nil != partErr {
		return partErr
	}

	if _, writeErr := partWriter.Write(frame); nil != writeErr {
		return writeErr
	}

	return nil
}
