package solve_test

import (
	"os"
	"testing"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"
)

func BenchmarkAll(b *testing.B) {
	// go decided that it should run benchmarks in this directory.
	os.Chdir(`..`)

	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		b.Run(iter.String(), func(b *testing.B) {
			srs, err := fetch.Read(iter)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Fail()
			} else if len(srs) == 0 {
				b.Logf("No cached results")
				b.Fail()
			}

			max := 10
			if iter.GetSize() > 15 {
				max = 2
			}
			for k := range srs {
				if len(srs) <= max {
					break
				}
				delete(srs, k)
			}

			for id, sr := range srs {
				b.Run("PuzzleID "+id, func(b *testing.B) {
					var sol model.Solution
					for n := 0; n < b.N; n++ {
						sol, err = solve.FromNodes(
							iter.GetSize(),
							sr.Input.ToNodes(),
						)
						if err != nil {
							b.Logf("got unexpected error: %q", err)
							b.Fail()
						}
						if sol.ToAnswer() != sr.Answer {
							b.Logf("expected answer %q, but got %q", sr.Answer, sol.ToAnswer())
							b.Fail()
						}
					}
				})
			}
		})
	}
}
