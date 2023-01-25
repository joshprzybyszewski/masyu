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
	_ = s.checkEntireRuleset()

	valid, solved := s.isValidAndSolved()
	if solved {
		sol, ok := s.toSolution()
		if ok {
			return sol, nil
		}
		panic(`dev error!`)
	} else if !valid {
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
	w.work <- initial

	// TODO instead of sending out "completed" white nodes, send out a completed
	// "row" or "col" of the puzzle (because we know there must be an even number of
	// lines through every row/col)

	var cpy state
	sendCpy := func() {
		if !cpy.rules.runAllChecks(&cpy) {
			return
		}
		defer func() {
			// if the work channel has been closed, then don't do anything.
			_ = recover()
		}()
		w.work <- cpy
	}
	var l, a bool
	var i int

	whites := make([]model.Coord, 0, len(initial.nodes))
	for i = len(initial.nodes) - 1; i >= 0; i-- {
		if initial.nodes[i].IsBlack {
			continue
		}
		if l, a = initial.horAt(initial.nodes[i].Row, initial.nodes[i].Col); l || a {
			continue
		}
		if l, a = initial.verAt(initial.nodes[i].Row, initial.nodes[i].Col); l || a {
			continue
		}
		if l, a = initial.horAt(initial.nodes[i].Row, initial.nodes[i].Col+1); l || a {
			continue
		}
		if l, a = initial.verAt(initial.nodes[i].Row+1, initial.nodes[i].Col); l || a {
			continue
		}
		whites = append(whites, initial.nodes[i].Coord)
	}
	start := 0
	var decisions, bit uint64

	for start < len(whites) {
		decisions = 0
		for {
			cpy = initial
			for bit, i = 0x01, 0; i < 32; i++ {
				if start+i >= len(whites) {
					break
				}

				if decisions&bit == bit {
					cpy.lineHor(whites[start+i].Row, whites[start+i].Col)
					bit <<= 1
					if decisions&bit == bit {
						cpy.lineHor(whites[start+i].Row, whites[start+i].Col+1)
					} else {
						cpy.avoidHor(whites[start+i].Row, whites[start+i].Col+1)
					}
				} else {
					cpy.lineVer(whites[start+i].Row, whites[start+i].Col)
					bit <<= 1
					if decisions&bit == bit {
						cpy.lineVer(whites[start+i].Row+1, whites[start+i].Col)
					} else {
						cpy.avoidVer(whites[start+i].Row+1, whites[start+i].Col)
					}
				}
				bit <<= 1
			}

			select {
			case <-ctx.Done():
				return
			default:
			}

			sendCpy()
			if decisions == 0xFFFFFFFFFFFFFFFF {
				break
			}
			decisions++
		}
		start += 32
	}
}
