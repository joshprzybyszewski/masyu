package solve

import (
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAllowedEmptyRow = 7
)

type horizontalCrossingPermState struct {
	knownCol model.Dimension
	// numCrossings is the number of lines at and before the knownCol
	numCrossings int

	perms applyFn
}

func getPermutationsFromStates(
	s model.Size,
	states []horizontalCrossingPermState,
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

func getInitialPermutations(
	initial state,
) []applyFn {
	row := getBestInitialStartingRow(&initial)
	if row == 0 {
		// couldn't find a good starting row?
		return nil
	}

	states := getPermutationsForRow(
		&initial,
		row,
		horizontalCrossingPermState{
			knownCol:     0,
			numCrossings: getNumLinesInRow(&initial, row),
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

	return getPermutationsFromStates(
		initial.size,
		states,
	)
}

func getBestInitialStartingRow(
	s *state,
) model.Dimension {
	var knownBest [40]struct {
		row model.Dimension

		numRulesAffected int
	}
	var col model.Dimension
	numEmpty := 0
	numRulesAffected := 0

	var l, a bool

	for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
		numEmpty = 0
		numRulesAffected = 0

		for col = model.Dimension(1); col <= model.Dimension(s.size); col++ {
			l, a = s.verAt(row, col)
			if !l && !a {
				numEmpty++
				numRulesAffected += len(s.rules.rules.verticals[row][col])
			}
		}
		if numRulesAffected >= knownBest[numEmpty].numRulesAffected {
			knownBest[numEmpty].numRulesAffected = numRulesAffected
			knownBest[numEmpty].row = row
		}
	}

	for numEmpty := maxAllowedEmptyRow; numEmpty > 0; numEmpty-- {
		if knownBest[numEmpty].row > 0 {
			return knownBest[numEmpty].row
		}
	}
	for numEmpty := maxAllowedEmptyRow + 1; numEmpty < len(knownBest); numEmpty++ {
		if knownBest[numEmpty].row > 0 {
			return knownBest[numEmpty].row
		}
	}
	// it's unlikely that all rows are filled...
	return 0
}

func getPermutationsForRow(
	s *state,
	row model.Dimension,
	cur horizontalCrossingPermState,
) []horizontalCrossingPermState {
	col := getNextEmptyCol(s, row, cur.knownCol+1)
	if col == 0 {
		// there wasn't an empty column found.
		if cur.numCrossings%2 == 0 {
			// if this row is valid, then return it.
			return []horizontalCrossingPermState{
				cur,
			}
		}
		return nil
	}

	output := getPermutationsForRow(
		s,
		row,
		horizontalCrossingPermState{
			knownCol:     col,
			numCrossings: cur.numCrossings,
			perms: func(s *state) {
				s.avoidVer(row, col)
				cur.perms(s)
			},
		},
	)

	output = append(output, getPermutationsForRow(
		s,
		row,
		horizontalCrossingPermState{
			knownCol:     col,
			numCrossings: cur.numCrossings + 1,
			perms: func(s *state) {
				s.lineVer(row, col)
				cur.perms(s)
			},
		},
	)...)

	return output
}

func getNextEmptyCol(
	s *state,
	row, col model.Dimension,
) model.Dimension {
	var l, a bool
	for ; col <= model.Dimension(s.size); col++ {
		l, a = s.verAt(row, col)
		if !l && !a {
			return col
		}
	}
	return 0
}

func getNumLinesInRow(
	s *state,
	row model.Dimension,
) int {
	numLines := 0
	for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
		if s.verLineAt(row, col) {
			numLines++
		}
	}
	return numLines
}
