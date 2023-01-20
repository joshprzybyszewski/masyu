package fetch

import (
	"fmt"

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
