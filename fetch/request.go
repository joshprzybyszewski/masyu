package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	requestsDisabled = false
)

func doRequest(
	method string,
	url string,
	header http.Header,
	data io.Reader,
) ([]byte, error) {

	if requestsDisabled {
		return nil, nil
	}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var resBytes []byte
		resBytes, err = io.ReadAll(response.Body)

		log.Printf("Got a not-ok response: %+v\n%s\n%s\n",
			response,
			response.Body,
			string(resBytes),
		)
		if err != nil {
			return nil, err
		}

		contentType := response.Header.Get(`Content-Type`)
		if strings.Contains(contentType, `text/plain`) {
			return nil, fmt.Errorf("bad response: \"%s\"", string(resBytes))
		}

		return nil, fmt.Errorf(`bad response from server`)
	}

	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return resBytes, nil
}
