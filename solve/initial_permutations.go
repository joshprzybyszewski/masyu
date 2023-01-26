package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAllowedEmptyRow = 7
)

type intialPermutationsFactory struct {
	permutationsFactory
	moreSpace [256]applyFn
}

func newIntialPermutationsFactory() intialPermutationsFactory {
	return intialPermutationsFactory{}
}

func (pf *intialPermutationsFactory) save(
	a applyFn,
) {
	if pf.numVals < uint8(len(pf.vals)) {
		pf.vals[pf.numVals] = a
	} else {
		pf.moreSpace[pf.numVals] = a
	}
	pf.numVals++
}

func (pf *intialPermutationsFactory) hasRoomForNumEmpty(
	numEmpty int,
) bool {
	numPerms := 1 << (numEmpty - 1)
	rem := len(pf.vals) + len(pf.moreSpace) - int(pf.numVals)
	return numPerms <= rem
}

func (pf *intialPermutationsFactory) chooseStart(
	byNumEmpty [40]model.Dimension,
) (int, model.Dimension) {
	for numEmpty := len(byNumEmpty); numEmpty > 0; numEmpty-- {
		if byNumEmpty[numEmpty] > 0 && pf.hasRoomForNumEmpty(numEmpty) {
			return numEmpty, byNumEmpty[numEmpty]
		}
	}

	return 0, 0
}
