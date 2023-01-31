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
	numPuzzlesPerIter = 500

	readmeFilename     = `README.md`
	resultsStartMarker = `<resultsMarker>`
	resultsStopMarker  = `</resultsMarker>`
)

const (
	resultsTimeout = 90 * time.Second
)

func Update() {
	var allResults [model.MaxIterator + 1][]time.Duration
	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		if iter >= 13 && iter <= 15 {
			// These are the massive ones
			continue
		}

		durs, err := getResults(iter, resultsTimeout)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			continue
		}
		allResults[iter] = append(allResults[iter], durs...)
	}

	var sb strings.Builder

	sb.WriteByte('\n')
	sb.WriteByte('\n')
	sb.WriteString(fmt.Sprintf("_GOOS: %s_\n\n", runtime.GOOS))
	sb.WriteString(fmt.Sprintf("_GOARCH: %s_\n\n", runtime.GOARCH))

	if stats, err := cpu.Info(); err == nil {
		sb.WriteString(fmt.Sprintf("_cpu: %s_\n\n", stats[0].ModelName))
	}

	sb.WriteString(fmt.Sprintf("_Solve timeout: %s_\n\n", resultsTimeout))

	sb.WriteString("|Puzzle|Min|p25|Median|p75|p95|max|sample size|\n")
	sb.WriteString("|-|-|-|-|-|-|-|-:|\n")

	for iter := model.MinIterator; iter < model.Iterator(len(allResults)); iter++ {
		sb.WriteString(getTableRow(iter, allResults[iter]))
	}

	sb.WriteString(fmt.Sprintf("\n_Last Updated: %s_\n", time.Now().Format(time.RFC822)))

	writeResultsToReadme(sb.String())

	fmt.Printf("Results updated.\n")
}

func getResults(
	iter model.Iterator,
	timeout time.Duration,
) ([]time.Duration, error) {
	fmt.Printf("Starting %s\n", iter)

	inputs, err := fetch.ReadN(iter, numPuzzlesPerIter)
	if err != nil {
		return nil, err
	}

	durs := make([]time.Duration, 0, len(inputs))

	for i, sr := range inputs {
		ns := sr.Input.ToNodes()

		t0 := time.Now()
		sol, err := solve.FromNodesWithTimeout(
			iter.GetSize(),
			ns,
			timeout,
		)
		dur := time.Since(t0)
		durs = append(durs, dur)
		if i%10 == 0 {
			fmt.Printf("\n%2d: ", i/10)
		}
		if err != nil {
			fmt.Print("X")
			continue
		}
		if dur > 5*time.Second {
			fmt.Print("!")
		} else if dur > time.Second {
			fmt.Print("*")
		} else if dur > 500*time.Millisecond {
			fmt.Print("-")
		} else {
			fmt.Print(".")
		}

		if sr.Answer != `` {
			if sr.Answer != sol.ToAnswer() {
				fmt.Printf("\nIs this correct?\n")
				fmt.Printf("%s\n", sr.Input)
				fmt.Printf("%s\n", sol.Pretty(ns))
			}
		}

		for numGCs := 0; numGCs < 3; numGCs++ {
			time.Sleep(10 * time.Millisecond)
			runtime.GC()
		}
	}

	fmt.Printf("\nFinished %s\n\n", iter)
	fmt.Printf("%s\n\n", getTableRow(iter, durs))
	return durs, nil
}

func getTableRow(
	iter model.Iterator,
	durs []time.Duration,
) string {
	if len(durs) == 0 {
		return ``
	}

	sort.Slice(durs, func(i, j int) bool {
		return durs[i] < durs[j]
	})

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("|%s|", iter))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[0])))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[len(durs)/4])))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[len(durs)/2])))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[len(durs)*3/4])))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[len(durs)*19/20])))
	sb.WriteString(fmt.Sprintf("%s|", roundDuration(durs[len(durs)-1])))
	sb.WriteString(fmt.Sprintf("%d|\n", len(durs)))
	return sb.String()
}

func roundDuration(
	d time.Duration,
) time.Duration {

	if d > time.Second {
		return d.Truncate(10 * time.Millisecond)
	}

	if d > time.Millisecond {
		return d.Truncate(10 * time.Microsecond)
	}

	if d > time.Microsecond {
		return d.Truncate(10 * time.Nanosecond)
	}

	return d
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
