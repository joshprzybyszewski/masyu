// +go:build !prod

package solve

import "time"

func SetTestTimeout() {
	maxAttemptDuration = 10 * time.Second
}
