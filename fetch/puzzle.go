package fetch

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joshprzybyszewski/masyu/model"
)

const (
	expGameScriptPrefix = ` var Game = {}; var Puzzle = {}; var task = '`
)

func Puzzle(
	i model.Iterator,
) (input, error) {
	if !i.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	puzz := input{
		iter:       i,
		size:       i.GetSize(),
		difficulty: i.GetDifficulty(),
	}

	url := buildURL(i)
	header := buildHeader()

	resp, err := get(url, header)
	if err != nil {
		return input{}, err
	}

	populateInput(
		resp,
		&puzz,
	)

	return puzz, nil
}

func populateInput(
	resp []byte,
	input *input,
) error {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return err
	}

	input.id = doc.Find(`#puzzleID`).First().Text()

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
