package solve

import (
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

func FromNodes(
	size model.Size,
	ns []model.Node,
) (model.Solution, error) {
	return FromNodesWithTimeout(
		size,
		ns,
		maxAttemptDuration,
	)
}

func FromNodesWithTimeout(
	size model.Size,
	ns []model.Node,
	dur time.Duration,
) (model.Solution, error) {

	s := newState(size, ns)

	ss := settle(&s)
	if ss == solved {
		return s.toSolution(), nil
	} else if ss == invalid {
		panic(`bad initialization`)
	}

	return solve(
		&s,
		dur,
	)
}
