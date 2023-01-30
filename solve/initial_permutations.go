package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	initialPermutationsFactoryNumVals = 1 << 12
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
	numEmptyInCol, col := pf.getBestNextStartingCol(cur)
	numEmptyInRow, row := pf.getBestNextStartingRow(cur)
	if col == 0 || !pf.hasRoomForNumEmpty(numEmptyInCol) {
		pf.populateBestRow(cur)
	} else if row == 0 || !pf.hasRoomForNumEmpty(numEmptyInRow) {
		pf.populateBestColumn(cur)
	} else if numEmptyInRow < numEmptyInCol {
		pf.populateBestColumn(cur)
	} else {
		pf.populateBestRow(cur)
	}

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
	var nodesInRow [model.MaxPointsPerLine]int
	for _, n := range s.nodes {
		nodesInRow[n.Row]++
	}

	var rowByNumEmpty [40]model.Dimension
	var numNodesInNumEmpty [40]int
	var ne, nn int

	for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
		ne = int(s.size) - int(s.crossings.rows[row]) - int(s.crossings.rowsAvoid[row])
		nn = nodesInRow[row] + nodesInRow[row+1]
		if rowByNumEmpty[ne] == 0 || nn > numNodesInNumEmpty[ne] {
			rowByNumEmpty[ne] = row
			numNodesInNumEmpty[ne] = nn
		}
	}

	return pf.chooseStart(rowByNumEmpty)
}

func (pf *initialPermutations) getBestNextStartingCol(
	s *state,
) (int, model.Dimension) {
	var nodesInCol [model.MaxPointsPerLine]int
	for _, n := range s.nodes {
		nodesInCol[n.Col]++
	}

	var colByNumEmpty [40]model.Dimension
	var numNodesInNumEmpty [40]int
	var ne, nn int

	for col := model.Dimension(1); col < model.Dimension(s.size); col++ {
		ne = int(s.size) - int(s.crossings.cols[col]) - int(s.crossings.colsAvoid[col])
		if colByNumEmpty[ne] == 0 || nn > numNodesInNumEmpty[ne] {
			colByNumEmpty[ne] = col
			numNodesInNumEmpty[ne] = nn
		}
	}

	return pf.chooseStart(colByNumEmpty)
}

func (pf *initialPermutations) chooseStart(
	byNumEmpty [40]model.Dimension,
) (int, model.Dimension) {

	for numEmpty := len(byNumEmpty) - 1; numEmpty > 2; numEmpty-- {
		if byNumEmpty[numEmpty] > 0 && pf.hasRoomForNumEmpty(numEmpty) {
			return numEmpty, byNumEmpty[numEmpty]
		}
	}

	// it's unlikely that all rows are filled...
	return 0, 0
}
