// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 3) //建立大小為 3 的 buffered channel
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
