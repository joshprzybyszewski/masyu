package solve_test

import (
	"os"
	"testing"
	"time"

	"github.com/joshprzybyszewski/masyu/fetch"
	"github.com/joshprzybyszewski/masyu/model"
	"github.com/joshprzybyszewski/masyu/solve"
)

func TestSpecifics(t *testing.T) {
	fetch.DisableHTTPCalls()
	solve.SetTestTimeout()
	// go decided that it should run tests in this directory.
	os.Chdir(`..`)

	testCases := []struct {
		iter model.Iterator
		id   string
	}{{
		iter: 0,
		id:   `5,734,527`,
		// }, {
		// 	iter: 2,
		// 	id:   `150,618`,
		// }, {
		// 	iter: 1,
		// 	id:   `1,527,476`,
		// }, {
		// 	iter: 5,
		// 	id:   `193,319`,
		// }, {
		// 	iter: 3,
		// 	id:   `1,817,845`,
		// }, {
		// 	iter: 3,
		// 	id:   `5,995,199`,
		// }, {
		// iter: 6,
		// id:   `7,191,910`,
		// }, {
		// 	iter: 8,
		// 	id:   `5,573,288`,
		// }, {
		// 	iter: 9,
		// 	id:   `3,118,955`,
	}}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.iter.String()+` `+tc.id, func(t *testing.T) {
			sr, err := fetch.ReadID(tc.iter, tc.id)
			if err != nil || sr.Input == nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}

			ns := sr.Input.ToNodes()
			sol, err := solve.FromNodesWithTimeout(
				tc.iter.GetSize(),
				ns,
				50*time.Second,
			)
			if err != nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}
			if sol.ToAnswer() != sr.Answer {
				t.Logf("Incorrect Answer\n")
				t.Logf("Exp: %q\n", sr.Answer)
				t.Logf("Act: %q\n", sol.ToAnswer())
				t.Logf("Board:\n%s\n\n", sol.Pretty(ns))
				t.Fail()
			}
		})
	}
}

func TestAccuracy(t *testing.T) {
	fetch.DisableHTTPCalls()
	solve.SetTestTimeout()
	// go decided that it should run tests in this directory.
	os.Chdir(`..`)
	max := model.MaxIterator

	for iter := model.MinIterator; iter <= max; iter++ {
		t.Run(iter.String(), func(t *testing.T) {
			srs, err := fetch.ReadN(iter, 100)
			if err != nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}
			if len(srs) == 0 {
				t.Skip()
			}

			for _, sr := range srs {
				sr := sr
				if sr.Answer == `` {
					t.Logf("Unknown answer: %+v", sr)
					t.Fail()
				}
				t.Run(sr.Input.ID, func(t *testing.T) {
					ns := sr.Input.ToNodes()
					sol, err := solve.FromNodesWithTimeout(
						iter.GetSize(),
						ns,
						time.Duration(iter+1)*100*time.Millisecond,
					)
					if err != nil {
						t.Logf("Error fetching input: %q", err)
						t.Fail()
					}
					if sol.ToAnswer() != sr.Answer {
						t.Logf("Incorrect Answer\n")
						t.Logf("Exp: %q\n", sr.Answer)
						t.Logf("Act: %q\n", sol.ToAnswer())
						t.Logf("Board:\n%s\n\n", sol.Pretty(ns))
						t.Fail()
					}
				})
			}
		})
	}
}

func BenchmarkAll(b *testing.B) {
	fetch.DisableHTTPCalls()
	solve.SetTestTimeout()
	// go decided that it should run benchmarks in this directory.
	os.Chdir(`..`)

	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		b.Run(iter.String(), func(b *testing.B) {
			srs, err := fetch.ReadN(iter, 10)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Fail()
			} else if len(srs) == 0 {
				b.Logf("No cached results")
				b.Fail()
			}

			for _, sr := range srs {
				sr := sr
				if sr.Answer == `` {
					b.Logf("Unknown answer: %+v", sr)
					b.Fail()
				}
				b.Run("PuzzleID "+sr.Input.ID, func(b *testing.B) {
					var sol model.Solution
					for n := 0; n < b.N; n++ {
						sol, err = solve.FromNodesWithTimeout(
							iter.GetSize(),
							sr.Input.ToNodes(),
							time.Duration(iter+1)*100*time.Millisecond,
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
