package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func doRequest(
	method string,
	url string,
	header http.Header,
	data io.Reader,
) ([]byte, http.Header, error) {

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, nil, err
	}
	if header != nil {
		req.Header = header
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var resBytes []byte
		resBytes, err = ioutil.ReadAll(response.Body)

		log.Printf("Got a not-ok response: %+v\n%s\n%s\n",
			response,
			response.Body,
			string(resBytes),
		)
		if err != nil {
			return nil, nil, err
		}

		contentType := response.Header.Get(`Content-Type`)
		if strings.Contains(contentType, `text/plain`) {
			return nil, nil, fmt.Errorf("bad response: \"%s\"", string(resBytes))
		}

		return nil, nil, fmt.Errorf(`bad response from server`)
	}

	resBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	return resBytes, response.Header, nil
}
