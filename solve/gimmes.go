package solve

import (
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
			if n.Col <= 2 {
				s.avoidHor(n.Row, n.Col-1)
			}
			if n.Row <= 2 {
				s.avoidVer(n.Row-1, n.Col)
			}

			// blacks next to each other require an X in between
			if blacks[n.Row][n.Col+1] {
				s.avoidHor(n.Row, n.Col)
			}
			if blacks[n.Row+1][n.Col] {
				s.avoidVer(n.Row, n.Col)
			}

			// Thanks [wikipedia](https://en.wikipedia.org/wiki/Masyu#Solution_methods)
			// If there are two whites diagonal to a black, then the black cannot go
			// between the whites.
			if whites[n.Row-1][n.Col-1] && whites[n.Row+1][n.Col-1] {
				s.lineHor(n.Row, n.Col)
			}
			if whites[n.Row-1][n.Col+1] && whites[n.Row+1][n.Col+1] {
				s.lineHor(n.Row, n.Col-1)
			}
			if whites[n.Row-1][n.Col-1] && whites[n.Row-1][n.Col+1] {
				s.lineVer(n.Row, n.Col)
			}
			if whites[n.Row+1][n.Col-1] && whites[n.Row+1][n.Col+1] {
				s.lineVer(n.Row-1, n.Col)
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

			if !whites[n.Row][n.Col-1] && whites[n.Row][n.Col+1] && !whites[n.Row][n.Col+2] {
				r := newPairWhiteHorizontalRule(n.Row, n.Col)
				if n.Col > 1 {
					s.rules.rules.addHorizontalRule(n.Row, n.Col-2, &r)
				}
				s.rules.rules.addHorizontalRule(n.Row, n.Col+2, &r)
			}

			if !whites[n.Row-1][n.Col] && whites[n.Row+1][n.Col] && !whites[n.Row+2][n.Col] {
				r := newPairWhiteVerticalRule(n.Row, n.Col)
				if n.Row > 1 {
					s.rules.rules.addVerticalRule(n.Row-2, n.Col, &r)
				}
				s.rules.rules.addVerticalRule(n.Row+2, n.Col, &r)
			}
		}
	}
}
