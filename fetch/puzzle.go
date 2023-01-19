package fetch

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

func Puzzle(
	i model.Iterator,
) (input, error) {
	if !i.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	puzz := input{
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
