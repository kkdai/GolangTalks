// +build OMIT

package main

import "fmt"

func main() {
	// if
	if true {
		fmt.Printf("Hello")
	}

	// for loop
	for a := 0; a < 1; a++ {
		fmt.Printf(" Google\n")
	}

	// switch
	x := 2
	switch {
	case x == 2:
		fmt.Printf(" Good!")
	case x < 5:
		fmt.Printf(" No way!")
	}
}
