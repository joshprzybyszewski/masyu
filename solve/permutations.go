package solve

import "github.com/joshprzybyszewski/masyu/model"

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

func getAdvancedNextPermutations(
	cur *state,
) []applyFn {
	row := getBestNextStartingRow(cur)
	if row == 0 {
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

	return getPermutationsFromStates(
		cur.size,
		states,
	)
}

func getBestNextStartingRow(
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

	for numEmpty := 1; numEmpty < len(knownBest); numEmpty++ {
		if knownBest[numEmpty].row > 0 {
			return knownBest[numEmpty].row
		}
	}
	// it's unlikely that all rows are filled...
	return 0
}
