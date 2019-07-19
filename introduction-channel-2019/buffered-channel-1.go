// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 3) //Create buffered channel with a capacity of 3
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	fmt.Println(<-ch) //1
	fmt.Println(<-ch) //2
	fmt.Println(<-ch) //3
}
