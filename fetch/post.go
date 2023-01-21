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

	return doRequest(`POST`, url, header, data)
}
