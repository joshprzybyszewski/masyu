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

	perms := getSimpleNextPermutations(&w.state)

	// fmt.Printf("iterating %d permuations\n", len(perms))

	beforeAll := w.state
	for _, perm := range perms {
		perm(&w.state)
		w.process()
		if w.sendAnswer == nil {
			return
		}
		w.state = beforeAll
	}
}
