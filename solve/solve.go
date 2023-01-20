package solve

import "fmt"

func solve(
	s state,
) Solution {

	var l, a bool
	var s2 state

	var pending []state
	pending = append(pending, s)

	for len(pending) > 0 {
		s := pending[0]
		fmt.Printf("processing:\n%s\n", &s)
		s.settleNodes()
		solved, valid := s.toSolution()
		if solved != nil {
			fmt.Printf("found solution:\n%s\n", &s)
			return *solved
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

	fmt.Printf("did not find solution:\n%s\n", &s)
	return Solution{
		size: s.size,

		// hor: s.horizontalLines << 1,
		// ver s.verticalLines << 1,
	}
}
