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

func (w *worker) eliminateNearlyCycles() bool {
	if w.sendAnswer == nil {
		return false
	}

CHECK_CYCLES:
	c, isHor, seenNodes, hasNearlyCycle := w.state.paths.getNearlyCycle(&w.state)
	if !hasNearlyCycle {
		return true
	}

	if seenNodes == len(w.state.nodes) {
		if isHor {
			w.state.lineHor(c.Row, c.Col)
		} else {
			w.state.lineVer(c.Row, c.Col)
		}
		hasNearlyCycle = false
	}

	for hasNearlyCycle {
		if isHor {
			w.state.avoidHor(c.Row, c.Col)
		} else {
			w.state.avoidVer(c.Row, c.Col)
		}

		if w.state.hasInvalid || w.state.paths.hasCycle {
			break
		}

		c, isHor, seenNodes, hasNearlyCycle = w.state.paths.getNearlyCycle(&w.state)
		if seenNodes == len(w.state.nodes) {
			// error state: this should have been caught immediately
			return false
		}
	}

	valid, solved := w.state.isValidAndSolved()
	if solved {
		sol, ok := w.state.toSolution()
		if ok {
			w.sendAnswer(sol)
			w.sendAnswer = nil
			return false
		}
	}
	if !valid {
		return false
	}
	goto CHECK_CYCLES
}

func (w *worker) dfs() {
	// TODO before we eliminate "nearly cycles", we could check if there are any
	// rows/cols that have only one un-written line.
	if !w.eliminateNearlyCycles() {
		return
	}

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
