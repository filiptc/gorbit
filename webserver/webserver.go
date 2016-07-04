package webserver

import (
	"strconv"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/filiptc/gorbit/config"
	"github.com/filiptc/gorbit/webcam"
	"github.com/filiptc/gorbit/webserver/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/klaidliadon/console.v1"
)

type WebserverCommand struct {
	conf *config.Config
	cs   *console.Console
	o    *observable.Observable
}

func (c *WebserverCommand) Execute(args []string) error {
	gin.SetMode(gin.ReleaseMode)
	c.cs.Info("Server running...")
	wc := webcam.NewWebCam(c.o, c.cs)
	go wc.InitStream()
	r := gin.Default()
	routes.Register(r, c.conf, c.cs, c.o)
	return r.Run(":" + strconv.Itoa(c.conf.Port))
}

func NewWebserverCommand(config *config.Config) *WebserverCommand {
	console := console.Std()
	return &WebserverCommand{config, console, observable.New()}
}
