package routes

import "github.com/gin-gonic/gin"

type indexRoute struct {
	r *gin.Engine
}

func newIndex(r *gin.Engine) Route {
	return &indexRoute{r}
}

func (r *indexRoute) Register() {
	r.r.StaticFile("/", "./webserver/assets/index.html")
}
