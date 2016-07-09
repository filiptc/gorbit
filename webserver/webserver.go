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
	Width  uint32 `short:"w" long:"width" description:"Webcam width" default:"800"`
	Height uint32 `short:"h" long:"height" description:"Webcam height" default:"600"`
	Device string `short:"d" long:"device" description:"Device path" default:"/dev/video0"`
	Port   uint64 `short:"p" long:"port" description:"Server port" default:"8001"`

	conf *config.Config
	cs   *console.Console
	o    *observable.Observable
}

func (c *WebserverCommand) Execute(args []string) error {
	c.initConf()
	c.cs.Info("Server running...")
	wc := webcam.NewWebCam(c.o, c.cs, c.conf)
	go wc.InitStream()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routes.Register(r, c.conf, c.cs, c.o)
	return r.Run(":" + strconv.FormatUint(c.conf.Port, 10))
}

func NewWebserverCommand(config *config.Config) *WebserverCommand {
	console := console.Std()
	return &WebserverCommand{
		conf: config,
		cs:   console,
		o:    observable.New(),
	}
}
func (c *WebserverCommand) initConf() {
	c.conf.Cam.Device = c.Device
	c.conf.Cam.Width = c.Width
	c.conf.Cam.Height = c.Height
	c.conf.Port = c.Port
}
