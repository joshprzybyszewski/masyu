package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func FromNodes(
	size model.Size,
	ns []model.Node,
) (model.Solution, error) {

	s := newState(size, ns)

	solvedState := eliminateInitialAlmostCycles(s)
	if solvedState != nil {
		sol, ok := solvedState.toSolution()
		if ok {
			// fmt.Printf("trivial\n")
			return sol, nil
		}
	}

	return solve(
		&s,
	)
}
