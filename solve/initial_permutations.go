package solve

import (
	"context"
	"fmt"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAllowedEmptyRow = 9
)

func getInitialPermutations(
	ctx context.Context,
	initial state,
) []*state {
	t0 := time.Now()
	row := getMostEmptyRow(&initial)
	fmt.Printf("got most empty row (%d) in %s\n", row, time.Since(t0))
	if row == 0 {
		panic(`couldn't find one!`)
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

	return getPermutationsForRow(
		ctx,
		initial,
		row,
		col,
	)
}

func getMostEmptyRow(
	s *state,
) model.Dimension {
	var bestRow, col model.Dimension
	numEmpty := 0
	numRulesAffected := 0
	bestAffected := 0

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

		if numEmpty <= maxAllowedEmptyRow && numRulesAffected >= bestAffected {
			bestAffected = numRulesAffected
			bestRow = row
		}
	}

	return bestRow
}

func getPermutationsForRow(
	ctx context.Context,
	prev state,
	row, col model.Dimension,
) []*state {
	if ctx.Err() != nil {
		return nil
	}

	c2 := col + 1
	max := model.Dimension(prev.size) + 1
	var l, a bool
	numLines := 0
	for c2 < max {
		l, a = prev.verAt(row, c2)
		if l {
			numLines++
		} else if !a {
			break
		}
		c2++
	}
	if c2 == max {
		if numLines%2 == 0 {
			prev.avoidVer(row, col)
		} else {
			prev.lineVer(row, col)
		}
		return []*state{
			&prev,
		}
	}

	cpy := prev
	prev.lineVer(row, col)
	cpy.avoidVer(row, col)

	output := getPermutationsForRow(
		ctx,
		prev,
		row,
		c2,
	)
	output = append(output, getPermutationsForRow(
		ctx,
		cpy,
		row,
		c2,
	)...)

	return output
}
