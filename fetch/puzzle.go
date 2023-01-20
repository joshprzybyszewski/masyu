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
	return Update(i)
}
