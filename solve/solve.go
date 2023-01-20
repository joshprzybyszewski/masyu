package solve

import "fmt"

func solve(
	s state,
) Solution {
	fmt.Printf("state:\n%s\n", &s)
	return Solution{
		size: s.size,

		// hor: s.horizontalLines << 1,
		// ver s.verticalLines << 1,
	}
}
