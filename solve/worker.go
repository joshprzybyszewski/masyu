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
	if w.sendAnswer == nil {
		return
	}

	ss := settle(&w.state)
	if ss == solved {
		w.sendAnswer(w.state.toSolution())
		w.sendAnswer = nil
		return
	} else if ss == invalid {
		return
	}

	c, isHor, ok := w.state.getMostInterestingPath()
	if !ok {
		return
	}

	if isHor {
		s := w.state
		w.state.lineHor(c.Row, c.Col)
		w.process()

		w.state = s

		w.state.avoidHor(c.Row, c.Col)
		w.process()
		return
	}

	s := w.state
	w.state.lineVer(c.Row, c.Col)
	w.process()

	w.state = s
	w.state.avoidVer(c.Row, c.Col)
	w.process()
}
