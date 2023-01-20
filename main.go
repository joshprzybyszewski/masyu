package main

import (
	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"
)

func main() {
	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		compete(iter)
		break
	}
}

func compete(iter model.Iterator) error {
	input, err := fetch.Puzzle(iter)

	if err != nil {
		return err
	}

	ns := input.ToNodes()

	sol, err := solve.FromNodes(
		iter.GetSize(),
		ns,
	)
	if err != nil {
		return err
	}

	return fetch.Submit(
		&input,
		&sol,
	)
}
