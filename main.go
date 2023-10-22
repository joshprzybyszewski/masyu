package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/profile"
	"github.com/joshprzybyszewski/masyu/results"
	"github.com/joshprzybyszewski/masyu/solve"
)

var (
	updateResults = flag.Bool("results", false, "if set, then it will print the custom benchmark results")

	puzzID = flag.String("puzzID", "", "if set, then this will run a specific puzzle")

	iterStart     = flag.Int("start", int(model.MinIterator), "if set, this will override the iterators starting value")
	iterFinish    = flag.Int("finish", int(model.MinIterator+16), "if set, this will override the iterators final value")
	numIterations = flag.Int("numIterations", 1, "set this value to run through the puzzles many times")

	fetchNewPuzzles = flag.Bool("refresh", true, "if set, then it will fetch new puzzles")

	shouldProfile = flag.Bool("profile", false, "if set, will produce a profile output")
)

func main() {
	flag.Parse()

	if *updateResults {
		results.Update()
		return
	}

	if *shouldProfile {
		defer profile.Start()()
	}

	if *puzzID != `` {
		_ = runPuzzleID(
			model.Iterator(*iterStart),
			*puzzID,
		)
		return
	}

	for i := 0; i < *numIterations; i++ {
		for iter := model.Iterator(*iterStart); iter <= model.Iterator(*iterFinish); iter++ {
			for numGCs := 0; numGCs < 5; numGCs++ {
				time.Sleep(100 * time.Millisecond)
				runtime.GC()
			}

			err := compete(iter)
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				// panic(err)
			}
			time.Sleep(time.Second)
		}
	}
}

func compete(iter model.Iterator) error {

	fmt.Printf("Starting %s\n\t%s\n\n\n", iter, time.Now())
	input, err := fetch.Puzzle(iter)
	if *fetchNewPuzzles {
		input, err = fetch.GetNewPuzzle(iter)
	}
	fmt.Printf("Iter: %q, PuzzleID: %q, Task: %q\n", iter, input.ID, input.Task())

	if err != nil {
		return err
	}

	ns := input.ToNodes()

	t0 := time.Now()
	sol, err := solve.FromNodes(
		iter.GetSize(),
		ns,
	)
	defer func(dur time.Duration) {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Solution:\n%s\n", sol.Pretty(ns))
		fmt.Printf("Duration: %s\n\n\n", dur)
	}(time.Since(t0))

	if err != nil {
		_ = fetch.StorePuzzle(&input)
		return err
	}

	return fetch.Submit(
		&input,
		&sol,
	)
}

func runPuzzleID(
	iter model.Iterator,
	id string,
) error {
	fmt.Printf("Starting %s\n\t%s\n\n\n", iter, time.Now())
	input, err := fetch.GetPuzzleID(iter, id)
	if err != nil {
		fmt.Printf("runPuzzleID error: %+v\n", err)
		return err
	}
	fmt.Printf("Input %s\n", input)

	ns := input.ToNodes()

	t0 := time.Now()
	sol, err := solve.FromNodes(
		iter.GetSize(),
		ns,
	)
	defer func(dur time.Duration) {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Solution:\n%s\n", sol.Pretty(ns))
		fmt.Printf("Duration: %s\n\n\n", dur)
	}(time.Since(t0))

	if err != nil {
		return err
	}

	return fetch.Submit(
		&input,
		&sol,
	)
}
