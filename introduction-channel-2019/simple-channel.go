// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan bool)

	go func() {
		fmt.Println("In the goroutine")
		ch <- true
	}()

	fmt.Println("Keep waiting after goroutine run")
	<-ch
}
