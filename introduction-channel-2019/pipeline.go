// +build OMIT

package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("%d into first queue \n", n)
			out <- n
			fmt.Printf("%d completed into first queue \n", n)
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("%d into second queue \n", n)
			out <- n * n
			fmt.Printf("%d completed into second queue \n", n)
			time.Sleep(1 * time.Second)
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	out := sq(c)

	// Consume the output.
	for v := range out {
		fmt.Println("Result:", v)
	}
}
