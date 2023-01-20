package fetch

import (
	"io"
	"net/http"
)

func post(
	url string,
	header http.Header,
	data io.Reader,
) ([]byte, error) {

	b, _, err := doRequest(`POST`, url, header, data)
	return b, err
}
