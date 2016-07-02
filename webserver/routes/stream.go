package routes

import (
	"io"

	"mime/multipart"

	"fmt"

	"net/textproto"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/filiptc/gorbit/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/klaidliadon/console.v1"
)

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
		c.Stream(r.streamWebcam(c))
	})
}

func (r *streamRoute) streamWebcam(c *gin.Context) func(w io.Writer) bool {
	return func(w io.Writer) bool {
		mimeWriter := multipart.NewWriter(w)
		contentType := fmt.Sprintf("multipart/x-mixed-replace;boundary=%s", mimeWriter.Boundary())
		c.Writer.Header().Add("Content-Type", contentType)

		cb := handleNewFrame(c, mimeWriter)
		r.o.On("newFrame", cb)

		clientGone := w.(gin.ResponseWriter).CloseNotify()
		for {
			select {
			case <-clientGone:
				r.o.Off("newFrame", cb)
				return false
			}
		}
		return true
	}
}

func handleNewFrame(c *gin.Context, mw *multipart.Writer) func(frame []byte) {
	return func(frame []byte) {
		partHeader := make(textproto.MIMEHeader)
		partHeader.Add("Content-Type", "image/jpeg")
		partWriter, partErr := mw.CreatePart(partHeader)
		if nil != partErr {
			panic(partErr)
		}

		if _, writeErr := partWriter.Write(frame); nil != writeErr {
			panic(writeErr)
		}
	}
}
