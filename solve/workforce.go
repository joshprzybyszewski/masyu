package solve

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"

	"github.com/joshprzybyszewski/masyu/model"
)

type workforce struct {
	solution chan model.Solution

	work chan state

	freeWorkers int32
}

func newWorkforce() workforce {
	return workforce{
		solution: make(chan model.Solution, 1),
		work:     make(chan state, runtime.NumCPU()),
	}
}

func (w *workforce) start(
	ctx context.Context,
) {
	atomic.StoreInt32(&w.freeWorkers, int32(runtime.NumCPU()))
	for i := 0; i < runtime.NumCPU(); i++ {
		go w.startWorker(ctx)
	}
}

func (w *workforce) startWorker(
	ctx context.Context,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case s, ok := <-w.work:
			if !ok {
				return
			}
			atomic.AddInt32(&w.freeWorkers, -1)
			w.process(&s)
			atomic.AddInt32(&w.freeWorkers, 1)
		}
	}
}

func (w *workforce) stop() {
	// signal that "all the workers are free:#	"
	atomic.StoreInt32(&w.freeWorkers, 9001)
	close(w.work)
	close(w.solution)
}

func (w *workforce) solve(
	ctx context.Context,
	s *state,
) (model.Solution, error) {
	w.process(s)

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

func (w *workforce) process(
	s *state,
) {
	if atomic.LoadInt32(&w.freeWorkers) > 9000 {
		return
	}
	sol, solved, ok := s.toSolution()
	if solved {
		defer func() {
			// if the solution channel has been closed, then don't do anything.
			_ = recover()
		}()
		w.solution <- sol
		return
	}
	if !ok {
		return
	}

	c, isHor, ok := s.getMostInterestingPath()
	if !ok {
		return
	}

	if isHor {
		s2 := *s
		s2.lineHor(c.Row, c.Col)
		w.send(s2)

		s.avoidHor(c.Row, c.Col)
		w.send(*s)
		return
	}

	s2 := *s
	s2.lineVer(c.Row, c.Col)
	w.send(s2)

	s.avoidVer(c.Row, c.Col)
	w.send(*s)
}

func (w *workforce) send(
	s state,
) {
	fw := atomic.LoadInt32(&w.freeWorkers)
	if fw > 9000 {
		return
	}
	if fw == 0 {
		w.process(&s)
	} else {
		defer func() {
			// if the work channel has been closed, then don't do anything.
			_ = recover()
		}()
		w.work <- s
	}
}
