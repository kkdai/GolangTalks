// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 3) //Create buffered channel with a capacity of 3
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // close channel
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
