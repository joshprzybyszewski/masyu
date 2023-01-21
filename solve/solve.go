package solve

import (
	"context"
	"time"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxAttemptDuration = 10 * time.Second
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
