package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	initialPermutationsFactoryNumVals = 1 << 10
)

type initialPermutations struct {
	vals    [initialPermutationsFactoryNumVals]applyFn
	numVals uint16
}

func newInitialPermutations() initialPermutations {
	return initialPermutations{}
}

func (pf *initialPermutations) save(
	a applyFn,
) {
	pf.vals[pf.numVals] = a
	pf.numVals++
}

func (pf *initialPermutations) hasRoomForNumEmpty(
	numEmpty int,
) bool {
	if numEmpty == 0 {
		return false
	}
	numPerms := 1 << (numEmpty - 1)
	rem := len(pf.vals) - int(pf.numVals)
	return numPerms <= rem
}

func (pf *initialPermutations) populate(
	cur *state,
) {
	pf.populateBestRow(cur)
	pf.populateBestColumn(cur)
}

func (pf *initialPermutations) populateBestColumn(
	cur *state,
) {
	numEmpty, col := pf.getBestNextStartingCol(cur)
	if col == 0 || !pf.hasRoomForNumEmpty(numEmpty) {
		return
	}

	pf.buildColumn(
		cur,
		col,
		permutationsFactorySubstate{
			known:        0,
			numCrossings: getNumLinesInCol(cur, col),
			perms: func(s *state) {
				if s.hasInvalid {
					return
				}
				if getNextEmptyRow(s, col, 0) != 0 {
					fmt.Printf("Didn't fill the whole column\nColumn: %d\n%s\n", col, s)
					panic(`didn't fill the whole col?`)
				}
				if getNumLinesInCol(s, col)%2 != 0 {
					fmt.Printf("Wrong Number of Lines\nColumn: %d\n%s\n", col, s)
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)

}

func (pf *initialPermutations) buildColumn(
	s *state,
	col model.Dimension,
	cur permutationsFactorySubstate,
) {
	row := getNextEmptyRow(s, cur.known+1, col)
	if row == 0 {
		// there wasn't an empty column found.
		if cur.numCrossings%2 == 0 {
			pf.save(cur.perms)
		}
		return
	}

	a := permutationsFactorySubstate{
		known:        row,
		numCrossings: cur.numCrossings,
		perms: func(s *state) {
			s.avoidHor(row, col)
			cur.perms(s)
		},
	}
	l := permutationsFactorySubstate{
		known:        row,
		numCrossings: cur.numCrossings + 1,
		perms: func(s *state) {
			s.lineHor(row, col)
			cur.perms(s)
		},
	}

	if a.numCrossings >= int(s.size)/2 {
		pf.buildColumn(s, col, a)
		pf.buildColumn(s, col, l)
	} else {
		pf.buildColumn(s, col, l)
		pf.buildColumn(s, col, a)
	}
}

func (pf *initialPermutations) populateBestRow(
	cur *state,
) {
	numEmpty, row := pf.getBestNextStartingRow(cur)
	if row == 0 || !pf.hasRoomForNumEmpty(numEmpty) {
		// couldn't find a good starting row?
		return
	}

	pf.buildRow(
		cur,
		row,
		permutationsFactorySubstate{
			known:        0,
			numCrossings: getNumLinesInRow(cur, row),
			perms: func(s *state) {
				if s.hasInvalid {
					return
				}
				if getNextEmptyCol(s, row, 0) != 0 {
					panic(`didn't fill the whole row?`)
				}
				if getNumLinesInRow(s, row)%2 != 0 {
					fmt.Printf("Wrong Number of Lines\nRow: %d\n%s\n", row, s)
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)
}

func (pf *initialPermutations) buildRow(
	s *state,
	row model.Dimension,
	cur permutationsFactorySubstate,
) {

	col := getNextEmptyCol(s, row, cur.known+1)
	if col == 0 {
		if cur.numCrossings%2 == 0 {
			pf.save(cur.perms)
		}
		return
	}

	a := permutationsFactorySubstate{
		known:        col,
		numCrossings: cur.numCrossings,
		perms: func(s *state) {
			s.avoidVer(row, col)
			cur.perms(s)
		},
	}

	l := permutationsFactorySubstate{
		known:        col,
		numCrossings: cur.numCrossings + 1,
		perms: func(s *state) {
			s.lineVer(row, col)
			cur.perms(s)
		},
	}

	if a.numCrossings >= int(s.size)/2 {
		pf.buildRow(s, row, a)
		pf.buildRow(s, row, l)
	} else {
		pf.buildRow(s, row, l)
		pf.buildRow(s, row, a)
	}
}

func (pf *initialPermutations) getBestNextStartingRow(
	s *state,
) (int, model.Dimension) {
	var rowByNumEmpty [40]model.Dimension
	var ne int

	for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
		ne = int(s.size) - int(s.crossings.rows[row]) - int(s.crossings.rowsAvoid[row])
		rowByNumEmpty[ne] = row
	}

	return pf.chooseStart(rowByNumEmpty)
}

func (pf *initialPermutations) getBestNextStartingCol(
	s *state,
) (int, model.Dimension) {
	var colByNumEmpty [40]model.Dimension
	var ne int

	for col := model.Dimension(1); col < model.Dimension(s.size); col++ {
		ne = int(s.size) - int(s.crossings.cols[col]) - int(s.crossings.colsAvoid[col])
		colByNumEmpty[ne] = col
	}

	return pf.chooseStart(colByNumEmpty)
}

func (pf *initialPermutations) chooseStart(
	byNumEmpty [40]model.Dimension,
) (int, model.Dimension) {

	for numEmpty := 5; numEmpty < len(byNumEmpty); numEmpty++ {
		if byNumEmpty[numEmpty] > 0 {
			return numEmpty, byNumEmpty[numEmpty]
		}
		if !pf.hasRoomForNumEmpty(numEmpty) {
			break
		}
	}

	// it's unlikely that all rows are filled...
	return 0, 0
}
