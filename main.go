package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"
)

var (
	useKnown = false
)

func main() {
	for iter := model.MinIterator; iter < model.MaxIterator; iter++ {
		if iter >= 13 && iter <= 15 {
			// These are the massive ones
			continue
		}
		err := compete(iter)
		if err != nil {
			panic(err)
		}

		time.Sleep(100 * time.Millisecond)
		runtime.GC()
		time.Sleep(100 * time.Millisecond)
	}
}

func compete(iter model.Iterator) error {

	input, err := fetch.Puzzle(iter)
	if !useKnown {
		input, err = fetch.Update(iter)
	}
	t0 := time.Now()

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

	defer func(t1 time.Time) {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Solution:\n%s\n", sol.Pretty(ns))
		fmt.Printf("Duration: %s\n", t1.Sub(t0))
	}(time.Now())

	return fetch.Submit(
		&input,
		&sol,
	)
}
