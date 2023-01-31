package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func findGimmes(
	s *state,
) {

	var blacks [maxPinsPerLine][maxPinsPerLine]bool
	var whites [maxPinsPerLine][maxPinsPerLine]bool

	for _, n := range s.nodes {
		if n.IsBlack {
			blacks[n.Row][n.Col] = true
		} else {
			whites[n.Row][n.Col] = true
		}
	}

	for _, n := range s.nodes {
		if n.IsBlack {
			// blacks near the edge must go the other way
			if n.Col <= 2 {
				s.avoidHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col)
				s.lineHor(n.Row, n.Col+1)
			}
			if n.Row <= 2 {
				s.avoidVer(n.Row-1, n.Col)
				s.lineVer(n.Row, n.Col)
				s.lineVer(n.Row+1, n.Col)
			}
			if n.Col >= model.Dimension(s.size)-1 {
				s.avoidHor(n.Row, n.Col)
				s.lineHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col-2)
			}
			if n.Row >= model.Dimension(s.size)-1 {
				s.avoidVer(n.Row, n.Col)
				s.lineVer(n.Row-1, n.Col)
				s.lineVer(n.Row-2, n.Col)
			}

			// blacks next to each other require an X in between
			if blacks[n.Row][n.Col+1] {
				s.avoidHor(n.Row, n.Col)

				s.lineHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col-2)

				s.lineHor(n.Row, n.Col+1)
				s.lineHor(n.Row, n.Col+2)
			}
			if blacks[n.Row+1][n.Col] {
				s.avoidVer(n.Row, n.Col)

				s.lineVer(n.Row-1, n.Col)
				s.lineVer(n.Row-2, n.Col)

				s.lineVer(n.Row+1, n.Col)
				s.lineVer(n.Row+2, n.Col)
			}

			// Thanks [wikipedia](https://en.wikipedia.org/wiki/Masyu#Solution_methods)
			// If there are two whites diagonal to a black, then the black cannot go
			// between the whites.
			if whites[n.Row-1][n.Col-1] && whites[n.Row+1][n.Col-1] {
				s.avoidHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col)
				s.lineHor(n.Row, n.Col+1)
			}
			if whites[n.Row-1][n.Col+1] && whites[n.Row+1][n.Col+1] {
				s.avoidHor(n.Row, n.Col)
				s.lineHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col-2)
			}
			if whites[n.Row-1][n.Col-1] && whites[n.Row-1][n.Col+1] {
				s.avoidVer(n.Row-1, n.Col)
				s.lineVer(n.Row, n.Col)
				s.lineVer(n.Row+1, n.Col)
			}
			if whites[n.Row+1][n.Col-1] && whites[n.Row+1][n.Col+1] {
				s.avoidVer(n.Row, n.Col)
				s.lineVer(n.Row-1, n.Col)
				s.lineVer(n.Row-2, n.Col)
			}
		} else {
			// three or more whites in a row/col require an X in between
			if n.Col > 2 && whites[n.Row][n.Col-1] && whites[n.Row][n.Col-2] {
				s.avoidHor(n.Row, n.Col)
				s.avoidHor(n.Row, n.Col-1)
				s.avoidHor(n.Row, n.Col-2)
				s.avoidHor(n.Row, n.Col-3)

				s.lineVer(n.Row-1, n.Col)
				s.lineVer(n.Row, n.Col)

				s.lineVer(n.Row-1, n.Col-1)
				s.lineVer(n.Row, n.Col-1)

				s.lineVer(n.Row-1, n.Col-2)
				s.lineVer(n.Row, n.Col-2)
			}
			if n.Row > 2 && whites[n.Row-1][n.Col] && whites[n.Row-2][n.Col] {
				s.avoidVer(n.Row, n.Col)
				s.avoidVer(n.Row-1, n.Col)
				s.avoidVer(n.Row-2, n.Col)
				s.avoidVer(n.Row-3, n.Col)

				s.lineHor(n.Row, n.Col-1)
				s.lineHor(n.Row, n.Col)

				s.lineHor(n.Row-1, n.Col-1)
				s.lineHor(n.Row-1, n.Col)

				s.lineHor(n.Row-2, n.Col-1)
				s.lineHor(n.Row-2, n.Col)
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
