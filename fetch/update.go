package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

func Update(
	iter model.Iterator,
) (input, error) {
	if !iter.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	puzz := input{
		iter:       iter,
		size:       iter.GetSize(),
		difficulty: iter.GetDifficulty(),
	}

	header := buildHeader()
	data := buildNewPuzzleData(iter, header)

	resp, err := post(baseURL, header, data)
	if err != nil {
		return puzz, err
	}

	populateInput(
		resp,
		&puzz,
	)

	return puzz, nil
}

func buildNewPuzzleData(
	iter model.Iterator,
	header http.Header,
) io.Reader {

	formData := url.Values{}
	formData.Add(`size`, strconv.Itoa(int(iter)))
	formData.Add(`robot`, `1`)
	formData.Add(`new`, `	New Puzzle   `)

	encodedVals := formData.Encode()

	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("Content-Length", strconv.Itoa(len(encodedVals)))

	return strings.NewReader(encodedVals)
}
