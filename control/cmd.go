package control

import (
	"errors"

	"github.com/filiptc/gorbit/webcam"
)

type ControlCommand struct {
	Tilt  int  `short:"t" long:"tilt" description:"Tilt degrees"`
	Pan   int  `short:"p" long:"pan" description:"Pan degrees"`
	Reset bool `short:"r" long:"reset" description:"Reset pan/tilt"`
}

func NewControlCommand() *ControlCommand {
	return &ControlCommand{}
}

func (c *ControlCommand) Execute(args []string) error {
	ctrl := 0
	if c.Reset {
		ctrl++
	}

	if c.Pan != 0 {
		ctrl++
	}

	if c.Tilt != 0 {
		ctrl++
	}

	if ctrl > 1 {
		return errors.New("Multiple commands given. Please provide only one command.")
	}

	if ctrl == 0 {
		return errors.New("No command given. Please provide a command.")
	}

	if c.Pan != 0 {
		return webcam.PanTilt(c.Pan, 0)
	}

	if c.Tilt != 0 {
		return webcam.PanTilt(0, c.Tilt)
	}

	return webcam.Reset()
}
