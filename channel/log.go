package channel

import (
	"flag"
	"fmt"
	"time"
	"runtime"
	"reflect"
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

func currentFunction(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func init() {
	usedebug := flag.Bool("debug", false, "Enable log view")
	flag.Parse()
	debug = *usedebug
	if debug {
		Log("init","Enable log")

	}
}

func Log(i interface{},message string) {
	if !debug {
		return
	}
	var function string 
	switch i2 := i.(type) {
	case string: function = i2 
	default : function = currentFunction(i)		
	}
	start := time.Now()
	fmt.Println(string(ColorGreen), start .Format("2006/01/01  15:04:05"),"\t", string(ColorBlue), function , string(ColorYellow), message,string(ColorReset) )
}
