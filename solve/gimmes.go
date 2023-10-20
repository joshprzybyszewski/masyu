package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func findGimmes(
	s *state,
) {

	var whites [maxPinsPerLine][maxPinsPerLine]model.Value

	for _, n := range s.nodes {
		if n.IsBlack {
			continue
		}
		whites[n.Row][n.Col] = n.Value
	}

	for _, n := range s.nodes {
		if n.IsBlack {
			// don't know of any gimmes for blacks at this point in time.
			continue
		}

		// two whites next two each other (with different values) in a row/col require an X in between
		if n.Col > 2 &&
			whites[n.Row][n.Col-1] != 0 &&
			whites[n.Row][n.Col] != whites[n.Row][n.Col-1] {
			s.avoidHor(n.Row, n.Col)
			s.avoidHor(n.Row, n.Col-1)
			s.avoidHor(n.Row, n.Col-2)

			s.lineVer(n.Row-1, n.Col)
			s.lineVer(n.Row, n.Col)

			s.lineVer(n.Row-1, n.Col-1)
			s.lineVer(n.Row, n.Col-1)
		}
		if n.Row > 2 &&
			whites[n.Row-1][n.Col] != 0 &&
			whites[n.Row-1][n.Col] != whites[n.Row][n.Col] {
			s.avoidVer(n.Row, n.Col)
			s.avoidVer(n.Row-1, n.Col)
			s.avoidVer(n.Row-2, n.Col)

			s.lineHor(n.Row, n.Col-1)
			s.lineHor(n.Row, n.Col)

			s.lineHor(n.Row-1, n.Col-1)
			s.lineHor(n.Row-1, n.Col)
		}
	}
}
