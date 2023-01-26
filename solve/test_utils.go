// +go:build !prod

package solve

import "time"

func SetTestTimeout() {
	maxAttemptDuration = 15 * time.Second
}
