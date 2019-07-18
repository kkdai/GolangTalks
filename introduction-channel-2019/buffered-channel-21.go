// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 3) //Create size 3 buffered channel
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
