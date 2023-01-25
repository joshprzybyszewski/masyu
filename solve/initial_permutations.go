package solve

import (
	"context"
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAllowedEmptyRow = 7
)

func getInitialPermutations(
	ctx context.Context,
	initial state,
) []*state {
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
		initial,
		row,
	)

	goalCrossings := int(initial.size) / 2

	sort.Slice(perms, func(i, j int) bool {
		il, _ := getNumLinesInRow(perms[i], row)
		jl, _ := getNumLinesInRow(perms[j], row)

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

	return perms
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
	prev state,
	row model.Dimension,
) []*state {
	if ctx.Err() != nil {
		return nil
	}

	numLines, col := getNumLinesInRow(&prev, row)
	if col == 0 {
		// there wasn't an empty column found.
		if numLines%2 == 0 {
			// if this row is valid, then return it.
			return []*state{
				&prev,
			}
		}
		return nil
	} else if numLines >= 0 {
		// an empty column was found, and we know about some existing lines
		// either avoid or draw the final empty column and return this state.
		if numLines%2 == 0 {
			prev.avoidVer(row, col)
		} else {
			prev.lineVer(row, col)
		}
		return []*state{
			&prev,
		}
	}
	// there's at least two empty columns left. Populate the empty column
	// and avoid it, and append the results together

	cpy := prev
	cpy.avoidVer(row, col)
	output := getPermutationsForRow(
		ctx,
		cpy,
		row,
	)

	prev.lineVer(row, col)
	output = append(output, getPermutationsForRow(
		ctx,
		prev,
		row,
	)...)

	return output
}

func getNumLinesInRow(
	s *state,
	row model.Dimension,
) (int, model.Dimension) {
	numLines := 0
	var emptyCol model.Dimension
	var l, a bool
	for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
		l, a = s.verAt(row, col)
		if l {
			numLines++
		} else if !a {
			if emptyCol != 0 {
				return -1, emptyCol
			}
			emptyCol = col
		}
	}
	return numLines, emptyCol
}
