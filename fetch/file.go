package fetch

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

func storeAnswer(
	input *input,
	sol *model.Solution,
) error {
	known, err := Read(input.iter)
	if err != nil {
		return err
	}
	if sr, ok := known[input.ID]; ok && sr.Answer != `` {
		return nil
	}
	return appendToFile(
		getInputFilname(input.iter),
		getFileLine(input, sol),
	)
}

func StorePuzzle(
	input *input,
) error {
	known, err := Read(input.iter)
	if err != nil {
		return err
	}
	if _, ok := known[input.ID]; ok {
		return nil
	}
	return appendToFile(
		getInputFilname(input.iter),
		getFileLine(input, nil),
	)
}

func appendToFile(
	filename string,
	line string,
) error {
	f, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_RDWR|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(line)
	if err != nil {
		return err
	}
	return nil
}

func getInputFilname(
	iter model.Iterator,
) string {
	return fmt.Sprintf("puzzles/%s.txt", strings.ReplaceAll(iter.String(), " ", "_"))
}

func getFileLine(
	input *input,
	sol *model.Solution,
) string {
	return fmt.Sprintf("%s:%s:%s\n", input.ID, input.task, sol.ToAnswer())
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
				ID:   parts[0],
				task: parts[1],
				iter: iter,
			},
			Answer: parts[2],
		}
	}

	return output, nil
}

func ReadN(
	iter model.Iterator,
	n int,
) ([]savedResult, error) {
	if n <= 0 {
		return nil, nil
	}

	puzzles, err := Read(iter)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(puzzles))
	for id := range puzzles {
		if puzzles[id].Input == nil {
			continue
		}
		ids = append(ids, id)
	}
	sort.Strings(ids)

	if n < len(ids) {
		ids = ids[:n]
	}

	output := make([]savedResult, len(ids))
	for i := range output {
		output[i] = puzzles[ids[i]]
	}
	return output, nil
}

func ReadID(
	iter model.Iterator,
	puzzID string,
) (savedResult, error) {

	puzzles, err := Read(iter)
	if err != nil {
		return savedResult{}, err
	}

	return puzzles[puzzID], nil
}
