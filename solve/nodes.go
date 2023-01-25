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

	valid, solved := s.isValidAndSolved()
	if solved {
		sol, ok := s.toSolution()
		if ok {
			return sol, nil
		}
	}
	if !valid {
		panic(`how?`)
	}

	return solve(
		&s,
		dur,
	)
}
