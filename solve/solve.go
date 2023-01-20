package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAttempts = 5000
)

func solve(
	s state,
) (model.Solution, error) {

	sol, ok := solveDFS(s)
	if !ok {
		return model.Solution{}, fmt.Errorf("did not find solution")
	}
	return sol, nil
}

func solveDFS(
	s state,
) (model.Solution, bool) {

	s.settleNodes()
	sol, eop, solved, valid := s.toSolution()
	if solved {
		return sol, true
	}
	if !valid {
		return model.Solution{}, false
	}

	l, a := s.horAt(eop.Row, eop.Col)
	if !l && !a {
		s2 := s
		s2.lineHor(eop.Row, eop.Col)
		sol, ok := solveDFS(s2)
		if ok {
			return sol, true
		}

		s.avoidHor(eop.Row, eop.Col)
		return solveDFS(s)
	}

	l, a = s.verAt(eop.Row, eop.Col)
	if !l && !a {
		s2 := s
		s2.lineVer(eop.Row, eop.Col)
		sol, ok := solveDFS(s2)
		if ok {
			return sol, true
		}

		s.avoidVer(eop.Row, eop.Col)
		return solveDFS(s)
	}

	l, a = s.horAt(eop.Row, eop.Col-1)
	if !l && !a {
		s2 := s
		s2.lineHor(eop.Row, eop.Col-1)
		sol, ok := solveDFS(s2)
		if ok {
			return sol, true
		}
		s.avoidHor(eop.Row, eop.Col-1)
		return solveDFS(s)
	}

	l, a = s.verAt(eop.Row-1, eop.Col)
	if !l && !a {
		s2 := s
		s2.lineVer(eop.Row-1, eop.Col)
		sol, ok := solveDFS(s2)
		if ok {
			return sol, true
		}

		s.avoidVer(eop.Row-1, eop.Col)
		return solveDFS(s)
	}

	return model.Solution{}, false
}
