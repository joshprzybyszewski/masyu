package fetch

import (
	"net/http"
)

const (
	baseURL       = `https://www.puzzle-masyu.com/`
	hallOfFameURL = baseURL + `hallsubmit.php`
)

var (
	myCookies = []string{}
)

func buildHeader() http.Header {
	header := http.Header{}

	for _, c := range myCookies {
		header.Add("Cookie", c)
	}

	return header
}
