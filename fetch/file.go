package fetch

import (
	"fmt"
	"os"
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

func store(
	input *input,
	sol *model.Solution,
) error {
	known, err := Read(input.iter)
	if err != nil {
		return err
	}
	if _, ok := known[input.id]; ok {
		return nil
	}

	f, err := os.Create(getInputFilname(input.iter))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(getFileLine(input, sol))
	if err != nil {
		return err
	}
	return nil
}

func getInputFilname(
	iter model.Iterator,
) string {
	return fmt.Sprintf("puzzles/iter-%d.txt", iter)
}

func getFileLine(
	input *input,
	sol *model.Solution,
) string {
	return fmt.Sprintf("%s:%s:%s\n", input.id, input.task, sol.ToAnswer())
}

type savedResult struct {
	Input  *input
	Answer string
}

func Read(
	iter model.Iterator,
) (map[string]savedResult, error) {

	data, err := os.ReadFile(getInputFilname(iter))
	if err != nil {
		if strings.Contains(err.Error(), `no such file`) {
			return nil, nil
		}
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	output := make(map[string]savedResult, len(lines))

	var parts []string
	for _, line := range lines {
		parts = strings.Split(line, ":")
		if len(parts) != 3 {
			continue
		}

		output[parts[0]] = savedResult{
			Input: &input{
				id:   parts[0],
				task: parts[1],
				iter: iter,
			},
			Answer: parts[2],
		}
	}

	return output, nil
}
