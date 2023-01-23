package solve

import (
	"context"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAttemptDuration = 5 * time.Second
	// maxAttemptDuration = 15*time.Second
	// maxAttemptDuration = time.Minute
)

func solve(
	s *state,
) (model.Solution, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), maxAttemptDuration)
	defer cancelFn()

	w := newWorkforce()
	w.start(ctx)
	defer w.stop()

	return w.solve(ctx, s)
}
