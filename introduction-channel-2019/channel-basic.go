// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	done <- struct{}{}
	time.Sleep(1 * time.Second)
	fmt.Println("main close")
}
