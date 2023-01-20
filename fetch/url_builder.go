package fetch

import (
	"net/http"
	"strconv"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	baseURL       = `https://www.puzzle-masyu.com/`
	hallOfFameURL = baseURL + `hallsubmit.php`
)

var (
	myCookies = []string{}
)

func buildURL(
	i model.Iterator,
) string {
	return baseURL + `?size=` + strconv.Itoa(int(i))
}

func buildHeader() http.Header {
	header := http.Header{}

	for _, c := range myCookies {
		header.Add("Cookie", c)
	}

	return header
}
