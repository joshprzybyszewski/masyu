package solve

import (
	"context"
	"fmt"
	"runtime"

	"github.com/joshprzybyszewski/masyu/model"
)

type workforce struct {
	solution chan model.Solution

	work chan state

	workers [1]worker
}

func newWorkforce() workforce {
	wf := workforce{
		solution: make(chan model.Solution, 1),
		work:     make(chan state, runtime.NumCPU()),
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
			return
		case worker.state, ok = <-w.work: // s is on the heap
			if !ok {
				return
			}
			worker.process(ctx)
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
	w.work <- *s

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
