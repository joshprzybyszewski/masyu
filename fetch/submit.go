package fetch

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joshprzybyszewski/masyu/model"
)

func Submit(
	input *input,
	sol *model.Solution,
) error {
	defer StorePuzzle(input)
	if input == nil || input.param == `` {
		return nil
	}

	header := buildHeader()
	data := buildSubmissionData(input, sol, header)

	resp, err := post(baseURL, header, data)
	if err != nil {
		return err
	}

	header = buildHeader()
	data, err = buildHallOfFameSubmission(resp, header)
	if err != nil {
		fmt.Printf("Answer: %q\n", sol.ToAnswer())
		return err
	}

	_, err = post(hallOfFameURL, header, data)
	if err != nil {
		return err
	}

	return storeAnswer(input, sol)
}

func buildSubmissionData(
	input *input,
	sol *model.Solution,
	header http.Header,
) io.Reader {

	formData := url.Values{}
	formData.Add(`ansH`, sol.ToAnswer())
	formData.Add(`size`, strconv.Itoa(int(input.iter)))
	formData.Add(`param`, input.param)
	formData.Add(`robot`, `1`)
	formData.Add(`ready`, `   Done   `)

	encodedVals := formData.Encode()

	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("Content-Length", strconv.Itoa(len(encodedVals)))

	return strings.NewReader(encodedVals)
}

func buildHallOfFameSubmission(
	resp []byte,
	header http.Header,
) (io.Reader, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}

	solParams := doc.Find(`#pageContent`).First().
		Find(`input[name='solparams']`).
		AttrOr(`value`, `unset`)

	if solParams == `unset` {
		return nil, fmt.Errorf(`something was wrong`)
	}

	formData := url.Values{}
	formData.Add(`solparams`, solParams)
	formData.Add(`robot`, `1`)

	encodedVals := formData.Encode()
	data := strings.NewReader(encodedVals)

	header.Add("Content-Type", "application/x-www-form-urlencoded")

	return data, nil
}
