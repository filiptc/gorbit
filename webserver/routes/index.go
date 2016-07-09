package routes

import (
	"github.com/filiptc/gorbit/config"
	"github.com/gin-gonic/gin"
)

type indexRoute struct {
	r    *gin.Engine
	conf *config.Config
}

func newIndex(r *gin.Engine, conf *config.Config) Route {
	return &indexRoute{r, conf}
}

func (r *indexRoute) Register() {
	r.r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", gin.MIMEHTML)
		c.Writer.Write(r.conf.Index)
	})
	r.r.GET("/index.html", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", gin.MIMEHTML)
		c.Writer.Write(r.conf.Index)
	})
}
