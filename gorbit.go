package main

import (
	"C"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/filiptc/gorbit/config"
	"github.com/filiptc/gorbit/control"
	"github.com/filiptc/gorbit/formats"
	"github.com/filiptc/gorbit/webcam"
	"github.com/filiptc/gorbit/webserver"
	"github.com/jessevdk/go-flags"
)

func main() {
	var err error
	config := config.NewConfig()
	config.Font, err = Asset("assets/luximr.ttf")
	config.Index, err = Asset("assets/index.html")
	if err != nil {
		panic(err)
	}

	go webcam.ProcessCommands()
	parser := createCommandParser(config)

	if _, err = parser.Parse(); err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(1)
	}
}

func createCommandParser(config *config.Config) *flags.Parser {
	parser := flags.NewParser(nil, flags.Default)
	_, err := parser.AddCommand(
		"control",
		"control camera",
		"Control camera via commands like pan, tilt, reset, etc.",
		control.NewControlCommand(),
	)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand(
		"serve",
		"lauch webserver",
		"View, pan and tilt webcam via webserver",
		webserver.NewWebserverCommand(config),
	)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand(
		"formats",
		"view formats",
		"View webcam formats",
		formats.NewFormatsCommand(config),
	)
	if err != nil {
		panic(err)
	}

	return parser
}
