package solve

import (
	"context"

	"github.com/joshprzybyszewski/masyu/model"
)

type worker struct {
	state state

	sendAnswer func(model.Solution)
}

func newWorker(
	sendAnswer func(model.Solution),
) worker {
	return worker{
		sendAnswer: sendAnswer,
	}
}

func (w *worker) process(
	ctx context.Context,
) {
	if w.sendAnswer == nil || ctx.Err() != nil {
		return
	}

	sol, solved, ok := w.state.toSolution()
	if solved {
		w.sendAnswer(sol)
		w.sendAnswer = nil
		return
	}
	if !ok {
		return
	}

	c, isHor, ok := w.state.getMostInterestingPath()
	if !ok {
		return
	}

	if isHor {
		s := w.state
		w.state.lineHor(c.Row, c.Col)
		w.process(ctx)

		w.state = s

		w.state.avoidHor(c.Row, c.Col)
		w.process(ctx)
		return
	}

	s := w.state
	w.state.lineVer(c.Row, c.Col)
	w.process(ctx)

	w.state = s

	w.state.avoidVer(c.Row, c.Col)
	w.process(ctx)
}
