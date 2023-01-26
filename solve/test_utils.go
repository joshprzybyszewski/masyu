// +go:build !prod

package solve

import "time"

func SetTestTimeout() {
	maxAttemptDuration = time.Second
}
