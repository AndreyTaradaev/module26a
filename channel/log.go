package channel

import (
	"flag"
	"fmt"
	"time"
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
	usedebug := flag.Bool("debug", false, "Enable log view")
	flag.Parse()
	debug = *usedebug
	if debug {
		Log("Enable log")

	}
}

func Log(message string) {
	if !debug {
		return
	}
	start := time.Now()
	fmt.Println(string(ColorGreen), start .Format("2006/01/01  15:04:05"), string(ColorBlue), message,string(ColorReset) )
}
