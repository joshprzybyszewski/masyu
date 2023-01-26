package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type settledState uint8

const (
	invalid       settledState = 1
	solved        settledState = 2
	validUnsolved settledState = 3
)

// returns true if the state is still valid
func settle(
	s *state,
) settledState {
	if s.hasInvalid {
		return invalid
	}

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	if s.rules.runAllChecks(s) == invalid {
		return invalid
	}

	if s.hasInvalid {
		return invalid
	}

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	return validUnsolved
}

func settleCycle(
	s *state,
) settledState {

	// Please only call this if the state has a cycle
	// if s.paths.hasCycle {
	// 	return unexpected
	// }

	if s.paths.cycleSeenNodes != len(s.nodes) {
		// there's a cycle, but it doesn't include all of the nodes.
		return invalid
	}

	// we found a state that includes a cycle with all of the nodes.
	// avoid all of the remaining spots in the state, and see if it's
	// still valid: this eliminates the bad state of having tertiary paths set.

	avoidAllUnknowns(s)

	if !hasValidCrossings(s) {
		return invalid
	}

	if checkEntireRuleset(s) == invalid {
		return invalid
	}

	// re-validate our assumptions after checking all the rules
	if s.hasInvalid || s.paths.cycleSeenNodes != len(s.nodes) {
		return invalid
	}

	return solved
}

func hasValidCrossings(s *state) bool {
	bit := uint64(1 << 1)
	for i := 1; i <= int(s.size); i++ {
		if !hasValidCrossingsForVerticalBit(s, bit) ||
			!hasValidCrossingsForHorizontalBit(s, bit) {
			return false
		}
		bit <<= 1
	}

	return true
}

func hasValidCrossingsForVerticalBit(
	s *state,
	bit uint64,
) bool {
	var l uint8
	for i := 1; i <= int(s.size); i++ {
		if s.horizontalLines[i]&bit == bit {
			l++
		} else if s.horizontalAvoids[i]&bit != bit {
			return true
		}
	}
	return l%2 == 0
}

func hasValidCrossingsForHorizontalBit(
	s *state,
	bit uint64,
) bool {
	var l uint8
	for i := 1; i <= int(s.size); i++ {
		if s.verticalLines[i]&bit == bit {
			l++
		} else if s.verticalAvoids[i]&bit != bit {
			return true
		}
	}
	return l%2 == 0
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

func checkEntireRuleset(s *state) settledState {
	max := model.Dimension(s.size + 1)
	var col model.Dimension
	for row := model.Dimension(0); row <= max; row++ {
		for col = model.Dimension(0); col <= max; col++ {
			s.rules.checkHorizontal(row, col)
			s.rules.checkVertical(row, col)
		}
	}

	return s.rules.runAllChecks(s)
}
