package fetch

import (
	"net/http"
	"strconv"
)

const (
	baseURL = `https://www.puzzle-masyu.com/`
)

var (
	myCookies = []string{}
)

func buildURL(
	i Iterator,
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
