package solve

import (
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxEmpty = 4
)

type applyFn func(*state)

func getSimpleNextPermutations(
	s *state,
) []applyFn {

	c, isHor, ok := s.getMostInterestingPath()
	if !ok {
		return nil
	}

	if isHor {
		return []applyFn{
			func(s *state) {
				s.lineHor(c.Row, c.Col)
			},
			func(s *state) {
				s.avoidHor(c.Row, c.Col)
			},
		}
	}

	return []applyFn{
		func(s *state) {
			s.lineVer(c.Row, c.Col)
		},
		func(s *state) {
			s.avoidVer(c.Row, c.Col)
		},
	}
}

// TODO use this?
func getAdvancedNextPermutations(
	cur *state,
) []applyFn {

	// TODO consider making the row/col distinction based on how many are already filled?

	numEmpty, col := getBestNextStartingCol(cur)
	if col == 0 || numEmpty > maxEmpty {
		return getAdvancedNextPermutationsRow(cur)
	}

	states := getPermutationsForCol(
		cur,
		col,
		verticalCrossingPermState{
			knownRow:     0,
			numCrossings: getNumLinesInCol(cur, col),
			perms: func(s *state) {
				if getNextEmptyRow(s, col, 0) != 0 {
					panic(`didn't fill the whole col?`)
				}
				if getNumLinesInCol(s, col)%2 != 0 {
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)

	return getPermutationsFromVerticalStates(
		cur.size,
		states,
	)
}

func getAdvancedNextPermutationsRow(
	cur *state,
) []applyFn {
	numEmpty, row := getBestNextStartingRow(cur)
	if row == 0 || numEmpty > maxEmpty {
		// couldn't find a good starting row?
		return getSimpleNextPermutations(cur)
	}

	states := getPermutationsForRow(
		cur,
		row,
		horizontalCrossingPermState{
			knownCol:     0,
			numCrossings: getNumLinesInRow(cur, row),
			perms: func(s *state) {
				if getNextEmptyCol(s, row, 0) != 0 {
					panic(`didn't fill the whole row?`)
				}
				if getNumLinesInRow(s, row)%2 != 0 {
					panic(`didn't place the right amount of lines`)
				}
			},
		},
	)
	if len(states) > maxEmpty {
		return getSimpleNextPermutations(cur)
	}

	return getPermutationsFromStates(
		cur.size,
		states,
	)
}

func getBestNextStartingRow(
	s *state,
) (int, model.Dimension) {
	var rowByNumEmpty [40]model.Dimension

	for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
		rowByNumEmpty[s.size-model.Size(s.crossings.rows[row]+s.crossings.rowsAvoid[row])] = row
	}

	for numEmpty := 2; numEmpty < len(rowByNumEmpty); numEmpty++ {
		if rowByNumEmpty[numEmpty] > 0 {
			return numEmpty, rowByNumEmpty[numEmpty]
		}
	}
	// it's unlikely that all rows are filled...
	return 0, 0
}

func getBestNextStartingCol(
	s *state,
) (int, model.Dimension) {
	var colByNumEmpty [40]model.Dimension

	for col := model.Dimension(1); col < model.Dimension(s.size); col++ {
		colByNumEmpty[s.size-model.Size(s.crossings.cols[col]+s.crossings.colsAvoid[col])] = col
	}

	for numEmpty := 1; numEmpty <= int(s.size); numEmpty++ {
		if colByNumEmpty[numEmpty] > 0 {
			return numEmpty, colByNumEmpty[numEmpty]
		}
	}

	// it's unlikely that all rows are filled...
	return 0, 0
}

type verticalCrossingPermState struct {
	knownRow model.Dimension
	// numCrossings is the number of lines currently known to be in this row
	numCrossings int

	perms applyFn
}

func getPermutationsForCol(
	s *state,
	col model.Dimension,
	cur verticalCrossingPermState,
) []verticalCrossingPermState {
	row := getNextEmptyRow(s, cur.knownRow+1, col)
	if row == 0 {
		// there wasn't an empty column found.
		if cur.numCrossings%2 == 0 {
			// if this row is valid, then return it.
			return []verticalCrossingPermState{
				cur,
			}
		}
		return nil
	}

	output := getPermutationsForCol(
		s,
		col,
		verticalCrossingPermState{
			knownRow:     row,
			numCrossings: cur.numCrossings,
			perms: func(s *state) {
				cur.perms(s)
				s.avoidHor(row, col)
			},
		},
	)

	output = append(output, getPermutationsForCol(
		s,
		col,
		verticalCrossingPermState{
			knownRow:     row,
			numCrossings: cur.numCrossings + 1,
			perms: func(s *state) {
				cur.perms(s)
				s.lineHor(row, col)
			},
		},
	)...)

	return output
}

func getNextEmptyRow(
	s *state,
	row, col model.Dimension,
) model.Dimension {
	var l, a bool
	for ; row <= model.Dimension(s.size); row++ {
		l, a = s.horAt(row, col)
		if !l && !a {
			return row
		}
	}
	return 0
}

func getNumLinesInCol(
	s *state,
	col model.Dimension,
) int {
	return int(s.crossings.cols[col])
}

func getPermutationsFromVerticalStates(
	s model.Size,
	states []verticalCrossingPermState,
) []applyFn {
	goalCrossings := int(s) / 2

	sort.Slice(states, func(i, j int) bool {
		il := states[i].numCrossings
		jl := states[j].numCrossings

		di := goalCrossings - il
		if di < 0 {
			di = -di
		}

		dj := goalCrossings - jl
		if dj < 0 {
			dj = -dj
		}
		if di != dj {
			// whichever is closer to the goal number of crossings
			return di < dj
		}

		// whichever has more lines
		return il > jl
	})

	output := make([]applyFn, 0, len(states))
	for _, c := range states {
		output = append(output, c.perms)
	}
	return output
}
