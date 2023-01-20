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
	useKnown = true
)

func main() {
	for iter := model.MinIterator; iter < 8; /*model.MaxIterator*/ iter++ {
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

	t0 := time.Now()
	input, err := fetch.Puzzle(iter)
	if !useKnown {
		t0 = time.Now()
		input, err = fetch.Update(iter)
	}

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
		fmt.Printf("Input: %+v\n", input)
		fmt.Printf("Solution:\n%s\n", &sol)
		fmt.Printf("Duration: %s\n", t1.Sub(t0))
	}(time.Now())

	return fetch.Submit(
		&input,
		&sol,
	)
}
