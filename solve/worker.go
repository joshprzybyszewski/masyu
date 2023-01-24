package solve

import (
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

func (w *worker) process() {
	w.dfs()
}

func (w *worker) eliminateInitialAlmostCycles() {
	if w.sendAnswer == nil {
		return
	}

	ok := eliminateInitialAlmostCycles(&w.state)
	if ok {
		sol, ok := w.state.toSolution()
		if ok {
			w.sendAnswer(sol)
			w.sendAnswer = nil
			return
		}
	}
}

func (w *worker) dfs() {
	w.eliminateInitialAlmostCycles()

	if w.sendAnswer == nil {
		return
	}

	valid, solved := w.state.isValidAndSolved()
	if solved {
		sol, ok := w.state.toSolution()
		if ok {
			w.sendAnswer(sol)
			w.sendAnswer = nil
		}
		return
	}
	if !valid {
		return
	}

	c, isHor, ok := w.state.getMostInterestingPath()
	if !ok {
		return
	}

	if isHor {
		s := w.state
		w.state.lineHor(c.Row, c.Col)
		w.dfs()

		w.state = s

		w.state.avoidHor(c.Row, c.Col)
		w.dfs()
		return
	}

	s := w.state
	w.state.lineVer(c.Row, c.Col)
	w.dfs()

	w.state = s

	w.state.avoidVer(c.Row, c.Col)
	w.dfs()
}
