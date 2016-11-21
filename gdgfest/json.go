// +build OMIT

package main

import "encoding/json"
import "fmt"

type Response1 struct {
	Page   int
	Fruits []string
}

func main() {

	res1D := &Response1{
		Page: 1,
		PPAP: []string{"Pen", "Pineapple", "Apple Pen"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
}
