package fetch

import (
	"io"
	"net/http"
)

func post(
	url string,
	header http.Header,
	data io.Reader,
) ([]byte, http.Header, error) {

	return doRequest(`POST`, url, header, data)
}
