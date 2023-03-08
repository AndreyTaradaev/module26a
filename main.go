package main

import (
	//"fmt"
	"pipeline/channel"
)

func main() {
	var exit = make(chan bool)
	var source = make(chan int)
	ar := channel.CreateStage(channel.ReadConsole,channel.FilterNegative,channel.FilterSpecial,channel.ReadBuffer)
	c := ar.Run(exit, source)
	
	
	// write :=  func() {
	// 	for {
	// 		select {
	// 		case d := <-c:
	// 			fmt.Println(d)
	// 		case <-exit:				
	// 			return
	// 		}
	// 	}
	// }
	channel.WriteConsole(exit,c)

}
