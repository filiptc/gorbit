package routes

import (
	"github.com/GianlucaGuarini/go-observable"
	"github.com/filiptc/gorbit/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/klaidliadon/console.v1"
)

type Route interface {
	Register()
}

func Register(r *gin.Engine, conf *config.Config, cs *console.Console, o *observable.Observable) {
	newIndex(r, conf).Register()
	newMove(r, conf, cs).Register()
	newStream(r, conf, cs, o).Register()
}
