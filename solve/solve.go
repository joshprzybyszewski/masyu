package solve

import (
	"context"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	// maxAttemptDuration = 5 * time.Second
	// maxAttemptDuration = 15 * time.Second
	maxAttemptDuration = time.Minute
	// maxAttemptDuration = time.Hour
)

func solve(
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
