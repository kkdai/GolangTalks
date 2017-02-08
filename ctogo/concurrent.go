// +build OMIT

package main

import "fmt"

func someRoutine(c chan bool) {
	fmt.Println("hello goroutine")
	c <- true
}

func main() {
	forever := make(chan bool)
	fmt.Println("hello world")
	go someRoutine(forever)

	<-forever
}
