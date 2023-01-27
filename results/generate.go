package results

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"

	"github.com/shirou/gopsutil/cpu"
)

const (
	numPuzzlesPerIter = 10

	readmeFilename     = `README.md`
	resultsStartMarker = `<resultsMarker>`
	resultsStopMarker  = `</resultsMarker>`
)

func Update() {
	var allResults [model.MaxIterator + 1][]time.Duration
	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		if iter >= 13 && iter <= 15 {
			// These are the massive ones
			continue
		}

		durs, err := getResults(iter)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			continue
		}
		allResults[iter] = append(allResults[iter], durs...)
	}

	// TODO remove and write to README.md
	var sb strings.Builder

	sb.WriteByte('\n')
	sb.WriteString(fmt.Sprintf("_GOOS: %s_\n\n", runtime.GOOS))
	sb.WriteString(fmt.Sprintf("_GOARCH: %s_\n\n", runtime.GOARCH))

	if stats, err := cpu.Info(); err == nil {
		sb.WriteString(fmt.Sprintf("_cpu: %s_\n\n", stats[0].ModelName))
	}

	sb.WriteString("|Puzzle|Min|Median|p75|p95|sample size|\n")
	sb.WriteString("|-|-|-|-|-|-:|\n")

	for iter := model.MinIterator; iter < model.Iterator(len(allResults)); iter++ {
		if len(allResults[iter]) == 0 {
			continue
		}

		sort.Slice(allResults[iter], func(i, j int) bool {
			return allResults[iter][i] < allResults[iter][j]
		})

		sb.WriteString(fmt.Sprintf("|%s|", iter))
		sb.WriteString(fmt.Sprintf("%s|", allResults[iter][0]))
		sb.WriteString(fmt.Sprintf("%s|", allResults[iter][len(allResults[iter])/2]))
		sb.WriteString(fmt.Sprintf("%s|", allResults[iter][len(allResults[iter])*3/4]))
		sb.WriteString(fmt.Sprintf("%s|", allResults[iter][len(allResults[iter])*19/20]))
		sb.WriteString(fmt.Sprintf("%d|\n", len(allResults[iter])))
	}

	sb.WriteString(fmt.Sprintf("\n_Last Updated: %s_\n", time.Now().Format(time.RFC822)))

	writeResultsToReadme(sb.String())

	fmt.Printf("Results updated.\n")
}

func getResults(iter model.Iterator) ([]time.Duration, error) {
	fmt.Printf("Starting %s\n", iter)
	inputs, err := fetch.ReadN(iter, numPuzzlesPerIter)
	if err != nil {
		return nil, err
	}

	durs := make([]time.Duration, 0, len(inputs))

	for _, sr := range inputs {
		ns := sr.Input.ToNodes()

		t0 := time.Now()
		sol, err := solve.FromNodes(
			iter.GetSize(),
			ns,
		)
		dur := time.Since(t0)
		if err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
			continue
		}

		if sr.Answer != `` {
			if sr.Answer != sol.ToAnswer() {
				fmt.Printf("Is this correct?\n")
				fmt.Printf("%s\n", sr.Input)
				fmt.Printf("%s\n", &sol)
			}
		}

		durs = append(durs, dur)

		for numGCs := 0; numGCs < 3; numGCs++ {
			time.Sleep(10 * time.Millisecond)
			runtime.GC()
		}
	}

	fmt.Printf("Finished %s\n\n\n", iter)
	return durs, nil
}

func writeResultsToReadme(
	res string,
) {

	fmt.Printf("Writing these results to the readme:\n%s\n", res)

	content, err := getREADMEContent()
	if err != nil {
		fmt.Printf("Error fetching readme content: %s\n", err.Error())
		return
	}

	i1 := strings.Index(content, resultsStartMarker)
	if i1 < 0 {
		fmt.Printf("Could not find %q in the README\n", resultsStartMarker)
		return
	}
	i2 := strings.Index(content, resultsStopMarker)
	if i2 < 0 {
		fmt.Printf("Could not find %q in the README\n", resultsStopMarker)
		return
	}

	newContent := content[:i1+len(resultsStartMarker)] + res + content[i2:]

	err = writeREADME(newContent)
	if err != nil {
		fmt.Printf("Error writing readme content: %s\n", err.Error())
		return
	}
}

func getREADMEContent() (string, error) {
	data, err := os.ReadFile(readmeFilename)
	if err != nil {
		return ``, err
	}
	return string(data), nil
}

func writeREADME(
	content string,
) error {
	f, err := os.Create(readmeFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
