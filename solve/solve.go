package solve

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

func solve(
	s state,
) (model.Solution, error) {

	var l, a bool
	var s2 state

	var pending []state
	pending = append(pending, s)

	for len(pending) > 0 {
		s := pending[0]
		s.settleNodes()
		sol, solved, valid := s.toSolution()
		if solved {
			return sol, nil
		} else if valid {
			for r := 1; r <= int(s.size); r++ {
				for c := 1; c <= int(s.size); c++ {
					l, a = s.horAt(r, c)
					if !l && !a {
						s2 = s
						s2.lineHor(r, c)
						pending = append(pending, s2)

						s2 = s
						s2.avoidHor(r, c)
						pending = append(pending, s2)
					}

					l, a = s.verAt(r, c)
					if !l && !a {
						s2 = s
						s2.lineVer(r, c)
						pending = append(pending, s2)
						s2 = s
						s2.avoidVer(r, c)
						pending = append(pending, s2)
					}
				}
			}
		}

		pending = pending[1:]
	}

	return model.Solution{}, fmt.Errorf("did not find solution")
}
