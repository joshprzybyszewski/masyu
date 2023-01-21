package solve

import "github.com/joshprzybyszewski/masyu/model"

func FromNodes(
	size model.Size,
	ns []model.Node,
) (model.Solution, error) {

	s := newState(size, ns)

	return solve(
		&s,
	)
}
