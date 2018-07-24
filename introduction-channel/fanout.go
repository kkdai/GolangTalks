// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, 1000)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int, 1000)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i <= 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := range in {
					out <- n * n
					fmt.Printf("%d is out queue \n", n)
					time.Sleep(1 * time.Second)
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	// ... the rest is unchanged ...
}

func main() {
	// Set up the pipeline.
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	out := sq(sq(c))

	// Consume the output.
	for v := range out {
		fmt.Println(v)
	}
}
