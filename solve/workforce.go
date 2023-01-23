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
	w.work <- initial

	var cpy state
	sendCpy := func() {
		defer func() {
			// if the work channel has been closed, then don't do anything.
			_ = recover()
		}()
		w.work <- cpy
	}
	var l, a bool

	whites := make([]model.Coord, 0, len(initial.nodes))
	for _, n := range initial.nodes {
		if n.IsBlack {
			continue
		}
		if l, a = initial.horAt(n.Row, n.Col); !l && !a {
			whites = append(whites, n.Coord)
		}
	}
	start := 0
	var decisions, bit uint8
	var i int

	for start < len(whites) {
		for decisions = 0; decisions < 255; decisions++ {
			cpy = initial
			for bit, i = 0x01, 0; i < 8; i++ {
				if start+i >= len(whites) {
					break
				}

				if decisions&bit == bit {
					cpy.lineHor(whites[start+i].Row, whites[start+i].Col)
				} else {
					cpy.lineVer(whites[start+i].Row, whites[start+i].Col)
				}

				bit <<= 1
			}

			select {
			case <-ctx.Done():
				return
			default:
			}

			sendCpy()
		}
		start += 8
	}
}
