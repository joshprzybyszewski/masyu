package fetch

import (
	"net/http"
	"os"
	"strings"
)

func init() {
	secretStr := os.Getenv(`MASYU_SECRET`)
	if secretStr == `` {
		return
	}
	vals := strings.Split(secretStr, `,`)
	myCookies = append(myCookies, vals...)
}

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
