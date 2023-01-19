package fetch

import (
	"net/http"
)

func get(
	url string,
	header http.Header,
) ([]byte, error) {

	b, _, err := doRequest(`GET`, url, header, nil)
	return b, err
}
