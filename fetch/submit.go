package fetch

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func Submit(
	input *input,
	sol *model.Solution,
) error {

	header := buildHeader()
	data := buildSubmissionData(input, sol, header)

	resp, err := post(baseURL, header, data)
	if err != nil {
		return err
	}

	header = buildHeader()
	data, err = buildHallOfFameSubmission(resp, header)
	if err != nil {
		return err
	}

	_, err = post(hallOfFameURL, header, data)
	if err != nil {
		return err
	}

	return nil
}
