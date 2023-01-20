package solve

import "fmt"

func solve(
	s state,
) Solution {

	var pending []state
	pending = append(pending, s)

	for len(pending) > 0 {
		s := pending[0]
		s.settleNodes()
		solved, valid := s.toSolution()
		if solved != nil {
			fmt.Printf("state:\n%s\n", &s)
			return *solved
		} else if valid {
			// TODO mutate s a few times and add it to pending

		}

		pending = pending[1:]
	}

	fmt.Printf("state:\n%s\n", &s)
	return Solution{
		size: s.size,

		// hor: s.horizontalLines << 1,
		// ver s.verticalLines << 1,
	}
}
