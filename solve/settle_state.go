package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

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

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	if !s.rules.runAllChecks(s) {
		return invalid
	}

	if s.hasInvalid {
		return invalid
	}

	if s.paths.hasCycle {
		return settleCycle(s)
	}

	if ss := completeCrossings(s); ss != validUnsolved {
		return ss
	}

	return eliminateCycles(s)
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

	if !checkEntireRuleset(s) {
		return invalid
	}

	// re-validate our assumptions after checking all the rules
	if s.hasInvalid {
		return invalid
	}

	if s.paths.cycleSeenNodes != len(s.nodes) {
		// there's a cycle, but it doesn't include all of the nodes.
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

func completeCrossings(
	s *state,
) settledState {

	var changed bool
	var bit uint64
	var dim2, tmp model.Dimension
	var numLines uint8

	for dim1 := model.Dimension(1); dim1 < model.Dimension(s.size); dim1++ {
		bit = 1 << dim1
		numLines = 0
		dim2 = 0
		for tmp = 1; tmp <= model.Dimension(s.size); tmp++ {
			if s.horizontalLines[tmp]&bit == bit {
				numLines++
			} else if s.horizontalAvoids[tmp]&bit != bit {
				if dim2 == 0 {
					dim2 = tmp
				} else {
					dim2 = model.Dimension(s.size + 1)
					break
				}
			}
		}
		if dim2 == 0 {
			// we got through the whole vertical line, and all of them had
			// either a line or an avoid.
			if numLines%2 == 1 {
				return invalid
			}
		} else if dim2 <= model.Dimension(s.size) {
			changed = true
			if numLines%2 == 0 {
				s.avoidHor(dim2, dim1)
			} else {
				s.lineHor(dim2, dim1)
			}
		}

		numLines = 0
		dim2 = 0
		for tmp = 1; tmp <= model.Dimension(s.size); tmp++ {
			if s.verticalLines[tmp]&bit == bit {
				numLines++
			} else if s.verticalAvoids[tmp]&bit != bit {
				if dim2 == 0 {
					dim2 = tmp
				} else {
					dim2 = model.Dimension(s.size + 1)
					break
				}
			}
		}
		if dim2 == 0 {
			// we got through the whole horizontal line, and all of them had
			// either a line or an avoid.
			if numLines%2 == 1 {
				return invalid
			}
		} else if dim2 <= model.Dimension(s.size) {
			changed = true
			if numLines%2 == 0 {
				s.avoidVer(dim1, dim2)
			} else {
				s.lineVer(dim1, dim2)
			}
		}
	}

	if changed {
		return settle(s)
	}
	return validUnsolved
}
