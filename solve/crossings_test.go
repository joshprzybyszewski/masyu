package solve

import (
	"testing"

	"github.com/joshprzybyszewski/masyu/model"
	"github.com/stretchr/testify/assert"
)

func TestCrossings(t *testing.T) {
	// This is puzz 3,368,748
	s := newState(
		model.Size(8),
		puzz8x8_3368748_Nodes,
	)

	expString := `0 X 1 X 2 X 3 X 4 X 5 X 6 X 7 X 8 X 9 X 
X   X   X   X   X   X   X   X   X   X   
1 X * - *   *   *   B   *   * - * X     
X   |                           |   X   
2 X W X *   *   *   *   W   * X W X     
X   |                           |   X   
3 X *   *   *   *   *   *   *   * X     
X                                   X   
4 X *   W   W   *   W   *   *   * X     
X                           |   |   X   
5 X *   *   *   *   W   * X * X * X     
X                       X   |   |   X   
6 X *   *   W   *   * - * - B X * X     
X                       X   X   |   X   
7 X *   *   W   *   *   * X * X * X     
X   |   X                   X   |   X   
8 X * - W - *   *   *   * - * - B X     
X   X   X   X   X   X   X   X   X   X   
9 X   X   X   X   X   X   X   X   X   X 
X                                   X   
`

	assert.Equal(t, expString, s.String())

	assert.Equal(t, model.Dimension(2), s.crossings.cols[1])
	assert.Equal(t, model.Dimension(1), s.crossings.colsAvoid[1])

	assert.Equal(t, model.Dimension(1), s.crossings.cols[2])
	assert.Equal(t, model.Dimension(0), s.crossings.colsAvoid[2])

	assert.Equal(t, model.Dimension(1), s.crossings.cols[5])
	assert.Equal(t, model.Dimension(0), s.crossings.colsAvoid[5])

	assert.Equal(t, model.Dimension(2), s.crossings.cols[6])
	assert.Equal(t, model.Dimension(2), s.crossings.colsAvoid[6])

	assert.Equal(t, model.Dimension(2), s.crossings.cols[7])
	assert.Equal(t, model.Dimension(4), s.crossings.colsAvoid[7])

	_ = checkEntireRuleset(&s)
	ss := settle(&s)
	assert.Equal(t, validUnsolved, ss)

	assert.Equal(t, model.Dimension(2), s.crossings.cols[1])
	assert.Equal(t, model.Dimension(1), s.crossings.colsAvoid[1])

	assert.Equal(t, model.Dimension(1), s.crossings.cols[2])
	assert.Equal(t, model.Dimension(0), s.crossings.colsAvoid[2])

	assert.Equal(t, model.Dimension(1), s.crossings.cols[5])
	assert.Equal(t, model.Dimension(0), s.crossings.colsAvoid[5])

	assert.Equal(t, model.Dimension(2), s.crossings.cols[6])
	assert.Equal(t, model.Dimension(2), s.crossings.colsAvoid[6])

	assert.Equal(t, model.Dimension(2), s.crossings.cols[7])
	assert.Equal(t, model.Dimension(4), s.crossings.colsAvoid[7])

	pf := newPermutationsFactory()
	pf.populate(&s)

	assert.Equal(t, uint16(2), pf.numVals)

	beforeAll := s
	for i := uint16(0); i < pf.numVals; i++ {
		pf.vals[i](&s)
		t.Logf("Permutation %d:\n%s\n\n", i, &s)
		if i < pf.numVals-1 {
			assert.Equal(t, model.Dimension(4), s.crossings.cols[7], "for permutation %d", i)
			assert.Equal(t, model.Dimension(4), s.crossings.colsAvoid[7], "for permutation %d", i)
		} else {
			assert.Equal(t, model.Dimension(2), s.crossings.cols[7], "for permutation %d", i)
			assert.Equal(t, model.Dimension(6), s.crossings.colsAvoid[7], "for permutation %d", i)
		}
		_ = checkEntireRuleset(&s)
		ss := settle(&s)
		if i < pf.numVals-1 {
			assert.Equal(t, model.Dimension(4), s.crossings.cols[7], "for permutation %d", i)
			assert.Equal(t, model.Dimension(4), s.crossings.colsAvoid[7], "for permutation %d", i)
		} else {
			assert.Equal(t, model.Dimension(2), s.crossings.cols[7], "for permutation %d", i)
			assert.Equal(t, model.Dimension(6), s.crossings.colsAvoid[7], "for permutation %d", i)
		}
		t.Logf("Permutation %d Settled:\n%s\n\n", i, &s)
		assert.Equal(t, validUnsolved, ss, "for permutation %d", i)
		s = beforeAll
	}

	t.Logf("Original:\n%s\n\n", &s)
	assert.Equal(t, expString, s.String())
}

var (
	puzz8x8_3368748_Nodes = []model.Node{{
		Coord: model.Coord{
			Row: 0,
			Col: 4,
		},
		IsBlack: true,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 0,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 5,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 1,
			Col: 7,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 1,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 2,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 3,
			Col: 4,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 4,
			Col: 4,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 5,
			Col: 2,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 5,
			Col: 6,
		},
		IsBlack: true,
	}, {
		Coord: model.Coord{
			Row: 6,
			Col: 2,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 7,
			Col: 1,
		},
		IsBlack: false,
	}, {
		Coord: model.Coord{
			Row: 7,
			Col: 7,
		},
		IsBlack: true,
	}}
)
