package webcam

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

const (
	ProcessingCommandError = "A previous command is being processed"
	pan_command            = "Pan (relative)"
	tilt_command           = "Tilt (relative)"
	reset_command          = "Pan/Tilt Reset"
	antiJamDelay           = 500 * time.Millisecond
)

type command struct {
	x int
	y int
}

var commandQueue chan command

func PanTilt(x, y int) error {
	return enqueueCommand(command{x, y})
}

func Reset() error {
	return execCommand(reset_command, 0)
}

func ProcessCommands() {
	commandQueue = make(chan command, 1)
	throttle := time.Tick(antiJamDelay)
	for {
		<-throttle
		cmd := <-commandQueue

		if cmd.y != 0 {
			execCommand(tilt_command, cmd.y)
			time.Sleep(antiJamDelay)
		}

		if cmd.x != 0 {
			execCommand(pan_command, cmd.x)
		}

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

func execCommand(name string, value int) error {
	fmt.Printf("Issuing \"%s\": %d\n", name, value)
	cmd := exec.Command("uvcdynctrl", "-s", name, "--", strconv.Itoa(value))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(output))
	}
	return nil
}
