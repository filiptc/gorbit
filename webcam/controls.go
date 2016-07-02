package webcam

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

const ProcessingCommandError = "A previous command is being processed"

type command struct {
	name  string
	value int
}

var commandQueue chan command

func Pan(amount int) error {
	return enqueueCommand(command{"Pan (relative)", amount})
}

func Tilt(amount int) error {
	return enqueueCommand(command{"Tilt (relative)", amount})
}

func Reset() error {
	return enqueueCommand(command{"Pan/Tilt Reset", 0})
}

func ProcessCommands() {
	commandQueue = make(chan command, 2)
	antiJamDelay := 500 * time.Millisecond
	throttle := time.Tick(antiJamDelay)
	for {
		<-throttle
		cmd := <-commandQueue
		execCommand(cmd)
		throttle = time.Tick(antiJamDelay)
	}
}

func enqueueCommand(cmd command) error {
	select {
	case commandQueue <- cmd:
	default:
		return errors.New(ProcessingCommandError)
	}
	return nil
}

func execCommand(c command) error {
	cmd := exec.Command("uvcdynctrl", "-s", c.name, "--", strconv.Itoa(c.value))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(output))
	}
	return nil
}
