package fetch

import (
	"fmt"
)

func Puzzle(
	i Iterator,
) (input, error) {
	if !i.Valid() {
		return input{}, fmt.Errorf("invalid iterator!")
	}

	puzz := input{
		size:       i.size(),
		difficulty: i.difficulty(),
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
