package solve

import (
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

var (
	emptyApply = func(*state) {}
)

type path struct {
	model.Coord
	IsHorizontal bool
}

type affectsApply struct {
	affects int

	fn applyFn
}

type rules struct {
	// if "this" row/col changes, then run these other checks
	// [row][col]
	horizontals [maxPinsPerLine][maxPinsPerLine]affectsApply
	verticals   [maxPinsPerLine][maxPinsPerLine]affectsApply

	// unknowns describes the paths that aren't initialized known.
	// They should exist in a sorted manner, where the first one has the most
	// rules associated with it, and so on. This can be used to find "the most
	// interesting space to investigate next."
	unknowns []path
}

func newRules(
	size model.Size,
) *rules {
	r := rules{
		unknowns: make([]path, 0, 2*int(size)*int(size-1)),
	}

	for i := range r.horizontals {
		for j := range r.horizontals[i] {
			r.horizontals[i][j].fn = emptyApply
			r.verticals[i][j].fn = emptyApply
		}
	}

	return &r
}

func (r *rules) populateRules(
	s *state,
) {

	for i := range s.nodes {
		if s.nodes[i].IsBlack {
			r.addBlackNode(s.nodes[i].Row, s.nodes[i].Col)
		} else {
			r.addWhiteNode(s.nodes[i].Row, s.nodes[i].Col)
		}
	}

	var pins [maxPinsPerLine][maxPinsPerLine]rule
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
			pins[row][col] = newDefaultRule(row, col)
		}
	}

	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col < model.Dimension(s.size); col++ {
			r.addHorizontalRule(row, col, &pins[row][col])
			r.addHorizontalRule(row, col, &pins[row][col+1])
		}
	}

	for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
		for row := model.Dimension(1); row < model.Dimension(s.size); row++ {
			r.addVerticalRule(row, col, &pins[row][col])
			r.addVerticalRule(row, col, &pins[row+1][col])
		}
	}
}

func (r *rules) populateUnknowns(
	s *state,
) {

	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {

			if !s.hasHorDefined(row, col) {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: true,
				})
			}

			if !s.hasVerDefined(row, col) {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: false,
				})
			}
		}
	}

	var ni, nj int
	var dri, dci, drj, dcj, tmp int
	is := int(s.size)

	sort.Slice(r.unknowns, func(i, j int) bool {
		if r.unknowns[i].IsHorizontal {
			ni = r.horizontals[r.unknowns[i].Row][r.unknowns[i].Col].affects
		} else {
			ni = r.verticals[r.unknowns[i].Row][r.unknowns[i].Col].affects
		}

		if r.unknowns[j].IsHorizontal {
			nj = r.horizontals[r.unknowns[j].Row][r.unknowns[j].Col].affects
		} else {
			nj = r.verticals[r.unknowns[j].Row][r.unknowns[j].Col].affects
		}

		if ni != nj {
			return ni < nj
		}

		// There are the same number of rules for this segment.
		// Prioritize a segment that is closer to the outer wall
		dri = int(r.unknowns[i].Row) - 1
		if tmp = is - int(r.unknowns[i].Row); tmp < dri {
			dri = tmp
		}
		dci = int(r.unknowns[i].Col) - 1
		if tmp = is - int(r.unknowns[i].Col); tmp < dci {
			dci = tmp
		}

		drj = int(r.unknowns[j].Row) - 1
		if tmp = is - int(r.unknowns[j].Row); tmp < drj {
			drj = tmp
		}
		dcj = int(r.unknowns[j].Col) - 1
		if tmp = is - int(r.unknowns[j].Col); tmp < dcj {
			dcj = tmp
		}
		ni = dri + dci
		nj = drj + dcj
		if ni != nj {
			return ni < nj
		}

		// They are equally close to the outer wall.
		// Prioritize the one in the top left.
		if r.unknowns[i].Row != r.unknowns[j].Row {
			return r.unknowns[i].Row < r.unknowns[j].Row
		}
		if r.unknowns[i].Col != r.unknowns[j].Col {
			return r.unknowns[i].Col < r.unknowns[j].Col
		}

		// Check horizontal first.
		return r.unknowns[i].IsHorizontal
	})
}

func (r *rules) addHorizontalRule(
	row, col model.Dimension,
	rule *rule,
) {

	r.horizontals[row][col].affects += rule.affects
	prev := r.horizontals[row][col].fn
	r.horizontals[row][col].fn = func(s *state) {
		rule.check(s)
		if !s.hasInvalid {
			prev(s)
		}
	}
}

func (r *rules) addVerticalRule(
	row, col model.Dimension,
	rule *rule,
) {
	r.verticals[row][col].affects += rule.affects
	prev := r.verticals[row][col].fn
	r.verticals[row][col].fn = func(s *state) {
		rule.check(s)
		if !s.hasInvalid {
			prev(s)
		}
	}
}

// TODO replace
func (r *rules) addBlackNode(
	row, col model.Dimension,
) {

	// ensure the black node is valid
	bv := newBlackValidator(row, col)
	if col > 1 {
		r.addHorizontalRule(row, col-2, &bv)
	}
	r.addHorizontalRule(row, col-1, &bv)
	r.addHorizontalRule(row, col, &bv)
	r.addHorizontalRule(row, col+1, &bv)
	if row > 1 {
		r.addVerticalRule(row-2, col, &bv)
	}
	r.addVerticalRule(row-1, col, &bv)
	r.addVerticalRule(row, col, &bv)
	r.addVerticalRule(row+1, col, &bv)

	// Look at extended "avoids"
	if col > 1 {
		left2 := newBlackL2Rule(row, col)
		r.addHorizontalRule(row, col-2, &left2)
		right2 := newBlackR2Rule(row, col)
		r.addHorizontalRule(row, col+1, &right2)
	}

	if row > 1 {
		up2 := newBlackU2Rule(row, col)
		r.addVerticalRule(row-2, col, &up2)
		down2 := newBlackD2Rule(row, col)
		r.addVerticalRule(row+1, col, &down2)
	}

	// Look at branches off the adjacencies.
	leftBranch := newBlackLBranchRule(row, col)
	r.addVerticalRule(row-1, col-1, &leftBranch)
	r.addVerticalRule(row, col-1, &leftBranch)

	rightBranch := newBlackRBranchRule(row, col)
	r.addVerticalRule(row-1, col+1, &rightBranch)
	r.addVerticalRule(row, col+1, &rightBranch)

	upBranch := newBlackUBranchRule(row, col)
	r.addHorizontalRule(row-1, col, &upBranch)
	r.addHorizontalRule(row-1, col-1, &upBranch)

	downBranch := newBlackDBranchRule(row, col)
	r.addHorizontalRule(row+1, col, &downBranch)
	r.addHorizontalRule(row+1, col-1, &downBranch)

	// look at inversions for black nodes
	if row > 1 && col > 1 {
		ih := newInvertHorizontalBlack(row, col)
		r.addHorizontalRule(row, col-2, &ih)
		r.addHorizontalRule(row, col+1, &ih)
		r.addVerticalRule(row-2, col, &ih)
		r.addVerticalRule(row+1, col, &ih)

		iv := newInvertVerticalBlack(row, col)
		r.addHorizontalRule(row, col-2, &iv)
		r.addHorizontalRule(row, col+1, &iv)
		r.addVerticalRule(row-2, col, &iv)
		r.addVerticalRule(row+1, col, &iv)
	}
}

// TODO replace
func (r *rules) addWhiteNode(
	row, col model.Dimension,
) {

	wv := newWhiteValidator(row, col)
	if col > 1 {
		r.addHorizontalRule(row, col-2, &wv)
	}
	r.addHorizontalRule(row, col-1, &wv)
	r.addHorizontalRule(row, col, &wv)
	r.addHorizontalRule(row, col+1, &wv)
	if row > 1 {
		r.addVerticalRule(row-2, col, &wv)
	}
	r.addVerticalRule(row-1, col, &wv)
	r.addVerticalRule(row, col, &wv)
	r.addVerticalRule(row+1, col, &wv)

	hb := newWhiteHorizontalBranchRule(row, col)
	if col > 1 {
		r.addHorizontalRule(row, col-2, &hb)
	}
	r.addHorizontalRule(row, col+1, &hb)
	r.addVerticalRule(row, col-1, &hb)
	r.addVerticalRule(row-1, col-1, &hb)
	r.addVerticalRule(row, col+1, &hb)
	r.addVerticalRule(row-1, col+1, &hb)

	vb := newWhiteVerticalBranchRule(row, col)
	if row > 1 {
		r.addVerticalRule(row-2, col, &vb)
	}
	r.addVerticalRule(row+1, col, &vb)
	r.addVerticalRule(row-1, col, &vb)
	r.addVerticalRule(row-1, col-1, &vb)
	r.addVerticalRule(row+1, col, &vb)
	r.addVerticalRule(row+1, col-1, &vb)
}
