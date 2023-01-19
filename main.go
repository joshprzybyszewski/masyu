package main

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/fetch"
)

func main() {
	input, err := fetch.Puzzle(0)

	if err != nil {
		panic(err)
	}

	fmt.Printf("input: %+v\n", input)

	ns := input.ToNodes()
	fmt.Printf("input.Convert(): %+v\n", ns)
}
