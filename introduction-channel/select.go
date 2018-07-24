// +build OMIT

package main

import (
	"fmt"
	"time"
)

func RunLoop(ch chan int) {
	for {
		select {
		//check
		case x, ok := <-ch:
			if ok {
				fmt.Printf("Value %d was read.\n", x)
			} else {
				fmt.Println("Channel closed!")
			}
		//time out pattern
		case <-time.After(time.Second * 1):
			fmt.Println("tick...")
		default:
			fmt.Println("No value ready, moving on.")
		}
	}
}
