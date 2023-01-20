package main

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"
)

func main() {
	iter := model.Iterator(0)
	input, err := fetch.Puzzle(iter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("input: %+v\n", input)

	ns := input.ToNodes()
	fmt.Printf("input.Convert(): %+v\n", ns)

	sol := solve.FromNodes(
		iter.GetSize(),
		ns,
	)
	fmt.Printf("Solution: %+v\n", sol)
}
