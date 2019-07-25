// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	unbuffer := make(chan int)
	close(unbuffer)
	unbuffer <- 1
	time.Sleep(1 * time.Second)
	fmt.Println("main close")
}
