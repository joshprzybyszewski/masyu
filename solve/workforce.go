package solve

import (
	"context"
	"fmt"
	"runtime"

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
		go w.startWorker(
			ctx,
			&w.workers[i],
		)
	}
}

func (w *workforce) startWorker(
	ctx context.Context,
	worker *worker,
) {
	var ok bool

	for {
		select {
		case <-ctx.Done():
			worker.sendAnswer = nil
			return
		case worker.state, ok = <-w.work: // s is on the heap
			if !ok {
				return
			}
			worker.process()
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
	} else if ss == invalid {
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
		}
	}()

	w.work <- initial

	pf := newIntialPermutationsFactory()
	pf.populate(&initial)
	if ctx.Err() != nil {
		return
	}

	tmp := initial
	// var ss settledState
	for i := uint8(0); i < pf.numVals; i++ {
		if i >= uint8(len(pf.vals)) {
			pf.moreSpace[i](&tmp)
		} else {
			pf.vals[i](&tmp)
		}
		w.work <- tmp

		// ss = settle(&tmp)
		// if ss == solved {
		// 	w.solution <- tmp.toSolution()
		// 	return
		// } else if ss == validUnsolved {
		// 	w.work <- tmp
		// }

		tmp = initial
	}

}
