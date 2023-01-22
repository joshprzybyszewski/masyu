package fetch

import (
	"fmt"

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

	srs, err := Read(i)
	if err != nil {
		return input{}, err
	}
	for _, sr := range srs {
		return *sr.Input, nil
	}
	return GetNewPuzzle(i)
}

func GetPuzzleID(
	i model.Iterator,
	id string,
) (input, error) {
	if !i.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	sr, err := ReadID(i, id)
	if err != nil {
		return input{}, err
	}
	if sr.Input == nil {
		return input{}, fmt.Errorf("puzzle not known")
	}
	return *sr.Input, nil
}
