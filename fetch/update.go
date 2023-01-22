package fetch

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joshprzybyszewski/masyu/model"
)

func GetNewPuzzle(
	iter model.Iterator,
) (input, error) {
	if !iter.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	puzz := input{
		iter: iter,
	}

	header := buildHeader()
	data := buildNewPuzzleData(iter, header)

	resp, err := post(baseURL, header, data)
	if err != nil {
		return input{}, err
	}

	err = populateInput(
		resp,
		&puzz,
	)
	if err != nil {
		return input{}, err
	}

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

func populateInput(
	resp []byte,
	input *input,
) error {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return err
	}

	input.ID = doc.Find(`#puzzleID`).First().Text()

	taskString, err := getTaskFromScriptText(
		doc.Find(`#rel`).Find(`script`).Text(),
	)
	if err != nil {
		log.Printf("getPuzzleInfo ajaxResponse: %q\n", doc.Find(`#ajaxResponse`).First().Text())
		return err
	}
	input.task = taskString

	input.param = doc.Find(`#puzzleForm`).First().
		Find(`.puzzleButtons input[name='param']`).
		AttrOr(`value`, `unset`)

	if input.param == `unset` {
		log.Printf("getPuzzleInfo ajaxResponse: %q\n", doc.Find(`#ajaxResponse`).First().Text())
		return fmt.Errorf("didn't have a secret param")
	}

	return nil
}

func getTaskFromScriptText(
	gameScript string,
) (string, error) {
	if len(gameScript) < len(expGameScriptPrefix) {
		return ``, fmt.Errorf(`gameScript did not start with expected prefix: %s`, gameScript)
	}
	if gameScript[:len(expGameScriptPrefix)] != expGameScriptPrefix {
		return ``, fmt.Errorf(`unexpected prefix! %q`, gameScript)
	}
	end := strings.Index(gameScript[len(expGameScriptPrefix):], `'`)
	return gameScript[len(expGameScriptPrefix) : len(expGameScriptPrefix)+end], nil
}
