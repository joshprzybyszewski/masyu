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
