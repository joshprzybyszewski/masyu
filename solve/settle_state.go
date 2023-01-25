package solve

import "github.com/joshprzybyszewski/masyu/model"

type settledState uint8

const (
	invalid settledState = iota
	solved
	validUnsolved
)

// returns true if the state is still valid
func settle(
	s *state,
) settledState {
	if s.hasInvalid {
		return invalid
	}

	if !s.rules.runAllChecks(s) {
		return invalid
	}

	if !s.hasValidCrossings() {
		return invalid
	}

	if s.paths.hasCycle {
		if s.paths.cycleSeenNodes != len(s.nodes) {
			// there's a cycle, but it doesn't include all of the nodes.
			return invalid
		}

		// we found a state that includes a cycle with all of the nodes.
		// avoid all of the remaining spots in the state, and see if it's
		// still valid: this eliminates the bad state of having tertiary paths set.

		avoidAllUnknowns(s)

		if !checkEntireRuleset(s) {
			return invalid
		}

		return solved
	}

	// TODO before we eliminate "nearly cycles", we could check if there are any
	// rows/cols that have only one un-written line.

	return eliminateCycles(s)
	// TODO more checks? (like along each of the rows/cols)
}

func avoidAllUnknowns(
	s *state,
) {
	var col model.Dimension
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col = model.Dimension(1); col <= model.Dimension(s.size); col++ {
			if !s.horLineAt(row, col) {
				s.avoidHor(row, col)
			}
			if !s.verLineAt(row, col) {
				s.avoidVer(row, col)
			}
		}
	}
}

func checkEntireRuleset(s *state) bool {
	for row := model.Dimension(0); row <= model.Dimension(s.size+1); row++ {
		for col := model.Dimension(0); col <= model.Dimension(s.size+1); col++ {
			s.rules.checkHorizontal(row, col, s)
			s.rules.checkVertical(row, col, s)
		}
	}

	return s.rules.runAllChecks(s)
}

func eliminateCycles(
	s *state,
) settledState {

	c, isHor, seenNodes, hasNearlyCycle := s.paths.getNearlyCycle(s)
	if !hasNearlyCycle {
		return validUnsolved
	}

	if seenNodes == len(s.nodes) {
		if isHor {
			s.lineHor(c.Row, c.Col)
		} else {
			s.lineVer(c.Row, c.Col)
		}
		return settle(s)
	}

	for hasNearlyCycle {
		if isHor {
			s.avoidHor(c.Row, c.Col)
		} else {
			s.avoidVer(c.Row, c.Col)
		}

		if s.hasInvalid || s.paths.hasCycle {
			break
		}

		c, isHor, seenNodes, hasNearlyCycle = s.paths.getNearlyCycle(s)
		if seenNodes == len(s.nodes) {
			// error state: this should have been caught immediately
			panic(`unexpected state`)
		}
	}

	return settle(s)
}
