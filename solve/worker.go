package solve

import (
	"context"

	"github.com/joshprzybyszewski/masyu/model"
)

type worker struct {
	sendAnswer func(model.Solution)

	state state
}

func newWorker(
	sendAnswer func(model.Solution),
) worker {
	w := worker{}

	w.sendAnswer = sendAnswer

	return w
}

func (w *worker) process(
	ctx context.Context,
) {
	if ctx.Err() != nil {
		return
	}

	ss := settle(&w.state)
	if ss == solved {
		w.sendAnswer(w.state.toSolution())
		return
	} else if ss == invalid {
		return
	} else if ss == unexpected {
		panic(`ahh`)
	}

	pf := newPermutationsFactory()

	cpy := w.state
	pf.populate(&w.state)
	for i := uint16(0); i < pf.numVals; i++ {
		pf.vals[i](&w.state)
		w.process(ctx)
		if ctx.Err() != nil {
			return
		}
		w.state = cpy
	}
}
