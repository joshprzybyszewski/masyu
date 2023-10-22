package solve

import (
	"context"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

var (
	// TODO 5 minutes is a long time.
	maxAttemptDuration = 5 * time.Minute
)

func solveWithWorkforce(
	s *state,
	dur time.Duration,
) (model.Solution, error) {
	if dur > maxAttemptDuration {
		dur = maxAttemptDuration
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), dur)
	defer cancelFn()

	w := newWorkforce()
	w.start(ctx)
	defer w.stop()
	defer cancelFn()

	return w.solve(ctx, s)
}
