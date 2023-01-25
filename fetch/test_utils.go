// +go:build !prod

package fetch

func DisableHTTPCalls() {
	requestsDisabled = true
}
