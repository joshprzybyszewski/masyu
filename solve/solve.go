package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAttempts = 10_000_000
)

var (
	depth = 0
)

func solve(
	s state,
) (model.Solution, error) {
	depth = 0

	sol, ok := solveDFS(s)
	if !ok {
		return model.Solution{}, fmt.Errorf("did not find solution")
	}
	return sol, nil
}

func solveDFS(
	s state,
) (model.Solution, bool) {
	if depth > maxAttempts {
		return model.Solution{}, false
	}
	depth++

	sol, solved, valid := s.toSolution()
	if solved {
		return sol, true
	}
	if !valid {
		return model.Solution{}, false
	}

	c, isHor, valid := s.getMostInterestingPath()
	if !valid {
		return model.Solution{}, false
	}

	if isHor {
		s2 := s
		s2.lineHor(c.Row, c.Col)
		sol, ok := solveDFS(s2)
		if ok {
			return sol, true
		}

		s.avoidHor(c.Row, c.Col)
		return solveDFS(s)
	}

	s2 := s
	s2.lineVer(c.Row, c.Col)
	sol, ok := solveDFS(s2)
	if ok {
		return sol, true
	}

	s.avoidVer(c.Row, c.Col)
	return solveDFS(s)
}
