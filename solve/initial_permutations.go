package solve

import (
	"context"
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

func getInitialPermutations(
	ctx context.Context,
	initial state,
) []applyFn {
	row := getBestStartingRow(&initial)
	if row == 0 {
		// couldn't find a good starting row?
		return nil
	}

	col := model.Dimension(1)
	max := model.Dimension(initial.size) + 1
	var l, a bool

	for col < max {
		l, a = initial.verAt(row, col)
		if !l && !a {
			break
		}
		col++
	}

	perms := getPermutationsForRow(
		ctx,
		&initial,
		row,
		horizontalCrossingPermState{
			knownCol:     0,
			numCrossings: 0,
			perms: func(s *state) {
			},
		},
	)

	goalCrossings := int(initial.size) / 2

	sort.Slice(perms, func(i, j int) bool {
		il := perms[i].numCrossings
		jl := perms[j].numCrossings

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

	output := make([]applyFn, 0, len(perms))
	for _, c := range perms {
		output = append(output, c.perms)
	}
	return output
}

func getBestStartingRow(
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
	ctx context.Context,
	s *state,
	row model.Dimension,
	cur horizontalCrossingPermState,
) []horizontalCrossingPermState {
	if ctx.Err() != nil {
		return nil
	}

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
		ctx,
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
		ctx,
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
