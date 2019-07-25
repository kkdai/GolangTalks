// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
		fmt.Println("Channel close")
	}()
	return out
}

func main() {
	c := gen(2, 3)

	// Fan-out
	c1 := sq(c)
	c2 := sq(c)

	// Fan-in: Consume the merged output from c1 and c2.
	done := make(chan struct{})
	out := merge(done, c1, c2)
	// 	done <- struct{}{} //What is I uncomment it? 4 or 9?
	fmt.Println(<-out)
	fmt.Println(<-out)
	done <- struct{}{} //What is I uncomment it? 4 or 9? or other?
	fmt.Println("main close")
	time.Sleep(1 * time.Second)
}
