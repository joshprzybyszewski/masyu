// +go:build !prod

package solve

import "time"

func SetTestTimeout() {
	maxAttemptDuration = 100 * time.Millisecond
}
