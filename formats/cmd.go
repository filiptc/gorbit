package formats

import (
	"fmt"

	"github.com/GianlucaGuarini/go-observable"
	"github.com/filiptc/gorbit/config"
	wc "github.com/filiptc/gorbit/webcam"
	"gopkg.in/klaidliadon/console.v1"
)

type FormatsCommand struct {
	Device string `short:"d" long:"device" description:"Device path" default:"/dev/video0"`

	cnf *config.Config
}

func NewFormatsCommand(cnf *config.Config) *FormatsCommand {
	return &FormatsCommand{cnf: cnf}
}

func (c *FormatsCommand) Execute(args []string) error {
	c.cnf.Cam.Device = c.Device
	w := wc.NewWebCam(observable.New(), console.Std(), c.cnf)
	for k, v := range w.Cam.GetSupportedFormats() {
		fmt.Printf("%s:\n", v)
		for _, s := range w.Cam.GetSupportedFrameSizes(k) {
			fmt.Printf("\t- %vx%v\n", s.MaxWidth, s.MaxHeight)
		}
		fmt.Println("")
	}
	return nil
}
