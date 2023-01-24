package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

func findGimmes(
	s *state,
) {

	var blacks [model.MaxPointsPerLine][model.MaxPointsPerLine]bool
	var whites [model.MaxPointsPerLine][model.MaxPointsPerLine]bool

	for _, n := range s.nodes {
		if n.IsBlack {
			blacks[n.Row][n.Col] = true
		} else {
			whites[n.Row][n.Col] = true
		}
	}

	for _, n := range s.nodes {
		if n.IsBlack {
			// blacks next to each other require an X in between
			if blacks[n.Row][n.Col+1] {
				s.avoidHor(n.Row, n.Col)
			}
			if blacks[n.Row+1][n.Col] {
				s.avoidVer(n.Row, n.Col)
			}
		} else {
			// three or more whites in a row/col require an X in between
			if n.Col > 2 && whites[n.Row][n.Col-1] && whites[n.Row][n.Col-2] {
				s.avoidHor(n.Row, n.Col-1)
				s.avoidHor(n.Row, n.Col-2)
			}
			if n.Row > 2 && whites[n.Row-1][n.Col] && whites[n.Row-2][n.Col] {
				s.avoidVer(n.Row-1, n.Col)
				s.avoidVer(n.Row-2, n.Col)
			}
		}
	}
}

func eliminateInitialAlmostCycles(
	s *state,
) bool {
	prev := *s
	var valid, solved bool
	c, isHor, ok := s.paths.getNearlyCycle(s)

	for ok {
		prev = *s
		if isHor {
			s.lineHor(c.Row, c.Col)
		} else {
			s.lineVer(c.Row, c.Col)
		}

		valid, solved = s.isValidAndSolved()
		if solved {
			_, ok := s.toSolution()
			if ok {
				return true
			}
		}
		if valid {
			fmt.Printf("Connecting %v (isHorizontal: %v)\n", c, isHor)
			fmt.Printf("before:\n%s\n", &prev)
			fmt.Printf("after:\n%s\n", &s)
			panic(`how?`)
		}

		*s = prev
		if isHor {
			s.avoidHor(c.Row, c.Col)
		} else {
			s.avoidVer(c.Row, c.Col)
		}

		c, isHor, ok = s.paths.getNearlyCycle(s)
	}
	return false
}
