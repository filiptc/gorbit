package routes

import (
	"fmt"

	"github.com/filiptc/gorbit/config"
	"github.com/filiptc/gorbit/webcam"
	"github.com/gin-gonic/gin"
	"gopkg.in/klaidliadon/console.v1"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type moveRoute struct {
	r    *gin.Engine
	conf *config.Config
	cs   *console.Console
}

func newMove(r *gin.Engine, conf *config.Config, cs *console.Console) Route {
	return &moveRoute{r, conf, cs}
}

func (r *moveRoute) Register() {
	r.r.POST("/move", func(c *gin.Context) {
		var json Position
		var err error

		if err = c.BindJSON(&json); err != nil {
			c.JSON(400, gin.H{"status": fmt.Errorf("Bad request: %s", err)})
			return
		}

		degreesX := json.X * float64(r.conf.FieldOfView.Width)
		degreesY := json.Y * float64(r.conf.FieldOfView.Height)
		stepsX := int(degreesX / r.conf.AngleFactor)
		stepsY := int(degreesY / r.conf.AngleFactor)
		if webcam.PanTilt(stepsX, stepsY); err != nil {
			c.JSON(400, gin.H{"status": fmt.Sprintf("Error: %s", err)})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})
}
