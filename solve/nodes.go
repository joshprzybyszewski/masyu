package solve

import (
	"fmt"
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
		fmt.Printf("%s\n", &s)
		panic(`bad initialization`)
	}

	return solve(
		&s,
		dur,
	)
}
