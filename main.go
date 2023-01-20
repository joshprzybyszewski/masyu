package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/profile"
	"github.com/joshprzybyszewski/masyu/solve"
)

var (
	iterStart     = flag.Int("start", int(model.MinIterator), "if set, this will override the iterators starting value")
	iterFinish    = flag.Int("finish", int(model.MaxIterator), "if set, this will override the iterators final value")
	skipLarge     = flag.Bool("skipbig", true, "if set, this won't attempt the largest puzzles")
	numIterations = flag.Int("numIterations", 1, "set this value to run through the puzzles many times")

	fetchNewPuzzles = flag.Bool("refresh", true, "if set, then it will fetch new puzzles")

	shouldProfile = flag.Bool("profile", false, "if set, will produce a profile output")
)

func main() {
	flag.Parse()

	if *shouldProfile {
		defer profile.Start()()
	}

	for i := 0; i < *numIterations; i++ {
		for iter := model.Iterator(*iterStart); iter <= model.Iterator(*iterFinish); iter++ {
			if *skipLarge && iter >= 13 && iter <= 15 {
				// These are the massive ones
				continue
			}
			err := compete(iter)
			if err != nil {
				// panic(err)
			}

			for numGCs := 0; numGCs < 3; numGCs++ {
				time.Sleep(100 * time.Millisecond)
				runtime.GC()
			}
		}
	}
}

func compete(iter model.Iterator) error {

	input, err := fetch.Puzzle(iter)
	if *fetchNewPuzzles {
		input, err = fetch.Update(iter)
	}

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
		return err
	}

	return fetch.Submit(
		&input,
		&sol,
	)
}
