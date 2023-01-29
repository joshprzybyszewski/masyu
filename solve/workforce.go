package solve

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	// numWorkers = 1
	numWorkers = 4
)

type workforce struct {
	solution chan model.Solution

	work chan state

	workers [numWorkers]worker
}

func newWorkforce() workforce {
	wf := workforce{
		solution: make(chan model.Solution, 1),
		work:     make(chan state, runtime.NumCPU()),
	}

	if len(wf.workers) > runtime.NumCPU() {
		panic(`dev error`)
	}

	for i := range wf.workers {
		i := i
		wf.workers[i] = newWorker(
			func(sol model.Solution) {
				defer func() {
					// if the solution channel has been closed, then don't do anything.
					_ = recover()
				}()
				wf.solution <- sol
			},
		)
	}

	return wf
}

func (w *workforce) start(
	ctx context.Context,
) {

	for i := range w.workers {
		i := i
		go w.startWorker(
			ctx,
			&w.workers[i],
			i,
		)
	}
}

func (w *workforce) startWorker(
	ctx context.Context,
	worker *worker,
	id int,
) {
	var ok bool

	idleLogDur := 500 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(idleLogDur):
			fmt.Printf("Worker %d is idle...\n", id)
			idleLogDur += idleLogDur
		case worker.state, ok = <-w.work:
			if !ok {
				return
			}
			worker.process(ctx)
			idleLogDur = 500 * time.Millisecond
		}
	}
}

func (w *workforce) stop() {
	close(w.work)
	close(w.solution)
}

func (w *workforce) solve(
	ctx context.Context,
	s *state,
) (model.Solution, error) {
	_ = checkEntireRuleset(s)
	ss := settle(s)
	if ss == solved {
		return s.toSolution(), nil
	} else if ss != validUnsolved {
		panic(`dev error!`)
	}

	go w.sendWork(ctx, *s)

	select {
	case <-ctx.Done():
		return model.Solution{}, fmt.Errorf("Ran out of time.")
	case sol, ok := <-w.solution:
		if !ok {
			return model.Solution{}, fmt.Errorf("did not find the solution")
		}
		return sol, nil
	}
}

func (w *workforce) sendWork(
	ctx context.Context,
	initial state,
) {
	defer func() {
		// if the work channel has been closed, then don't do anything.
		r := recover()
		if r != nil {
			fmt.Printf("caught: %+v\n", r)
			fmt.Printf("%s\n", debug.Stack())
		}
	}()

	if ctx.Err() != nil {
		return
	}

	w.work <- initial

	pf := newInitialPermutations()
	pf.populate(&initial)
	if ctx.Err() != nil {
		return
	}

	cpy := initial

	for i := 0; i < int(pf.numVals); i++ {
		pf.vals[i](&cpy)

		if ctx.Err() != nil {
			return
		}

		w.work <- cpy

		cpy = initial
	}

	// TODO even with _huge_ initial permutations, workers still become free
	// in relatively short order. we should start having the "running threads"
	// send their in-progress information over, if possible.
}
