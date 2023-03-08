package channel

import (
	"flag"
	"fmt"
)

var debug bool

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Enable log view")
	if debug {
		fmt.Println("Enable log")

	}

}

func Log(message string) {

}
