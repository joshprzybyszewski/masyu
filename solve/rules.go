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

	var nodeMap [maxPinsPerLine][maxPinsPerLine]model.Node

	for i := range s.nodes {
		nodeMap[s.nodes[i].Row][s.nodes[i].Col] = s.nodes[i]
	}

	for i := range s.nodes {
		if s.nodes[i].IsBlack {
			r.addBlackNode(s.nodes[i], &nodeMap)
		} else {
			r.addWhiteNode(s.nodes[i], &nodeMap)
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

func (r *rules) addBlackNode(
	node model.Node,
	nodes *[maxPinsPerLine][maxPinsPerLine]model.Node,
) {
	row, col := node.Row, node.Col
	v := node.Value

	// TODO make this actually the size
	size := model.Dimension(model.MaxPinsPerLine)

	// ensure the black node is valid
	// TODO remove newBlackValidator?
	bv := newBlackValidator(row, col)
	r.addHorizontalRule(row, col-1, &bv)
	r.addHorizontalRule(row, col, &bv)
	r.addVerticalRule(row-1, col, &bv)
	r.addVerticalRule(row, col, &bv)

	be := newBlackExpensiveRule(node, nodes)
	vd := model.Dimension(v)
	r.addHorizontalRule(row, col, &be)
	r.addVerticalRule(row, col, &be)
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col+delta < size {
			r.addHorizontalRule(row, col+delta, &be)
			r.addVerticalRule(row-1, col+delta, &be)
			r.addVerticalRule(row, col+delta, &be)
		}
		if row+delta < size {
			r.addVerticalRule(row+delta, col, &be)
			r.addHorizontalRule(row+delta, col-1, &be)
			r.addHorizontalRule(row+delta, col, &be)
		}
	}
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col >= delta {
			r.addHorizontalRule(row, col-delta, &be)
			r.addVerticalRule(row-1, col-delta, &be)
			r.addVerticalRule(row, col-delta, &be)
		}
		if row >= delta {
			r.addVerticalRule(row-delta, col, &be)
			r.addHorizontalRule(row-delta, col-1, &be)
			r.addHorizontalRule(row-delta, col, &be)
		}
	}

	bce := newNodeCheckerExpensiveRule(row, col, v)
	for delta := model.Dimension(0); delta < vd; delta++ {
		if col+delta < size {
			r.addHorizontalRule(row, col+delta, &bce)
		}
		if row+delta < size {
			r.addVerticalRule(row+delta, col, &bce)
		}
	}
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col >= delta {
			r.addHorizontalRule(row, col-delta, &bce)
		}
		if row >= delta {
			r.addVerticalRule(row-delta, col, &bce)
		}
	}
}

func (r *rules) addWhiteNode(
	node model.Node,
	nodes *[maxPinsPerLine][maxPinsPerLine]model.Node,
) {
	row, col := node.Row, node.Col
	v := node.Value

	// TODO make this actually the size
	size := model.Dimension(model.MaxPinsPerLine)

	// TODO remove newWhiteValidator?
	wv := newWhiteValidator(row, col)
	r.addHorizontalRule(row, col-1, &wv)
	r.addHorizontalRule(row, col, &wv)
	r.addVerticalRule(row-1, col, &wv)
	r.addVerticalRule(row, col, &wv)

	we := newWhiteExpensiveRule(row, col, v)
	vd := model.Dimension(v)
	r.addHorizontalRule(row, col, &we)
	r.addVerticalRule(row, col, &we)
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col+delta < size {
			r.addHorizontalRule(row, col+delta, &we)
			r.addVerticalRule(row-1, col+delta, &we)
			r.addVerticalRule(row, col+delta, &we)
		}
		if row+delta < size {
			r.addVerticalRule(row+delta, col, &we)
			r.addHorizontalRule(row+delta, col-1, &we)
			r.addHorizontalRule(row+delta, col, &we)
		}
	}
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col >= delta {
			r.addHorizontalRule(row, col-delta, &we)
			r.addVerticalRule(row-1, col-delta, &we)
			r.addVerticalRule(row, col-delta, &we)
		}
		if row >= delta {
			r.addVerticalRule(row-delta, col, &we)
			r.addHorizontalRule(row-delta, col-1, &we)
			r.addHorizontalRule(row-delta, col, &we)
		}
	}

	wce := newNodeCheckerExpensiveRule(row, col, v)
	for delta := model.Dimension(0); delta < vd; delta++ {
		if col+delta < size {
			r.addHorizontalRule(row, col+delta, &wce)
		}
		if row+delta < size {
			r.addVerticalRule(row+delta, col, &wce)
		}
	}
	for delta := model.Dimension(1); delta <= vd; delta++ {
		if col >= delta {
			r.addHorizontalRule(row, col-delta, &wce)
		}
		if row >= delta {
			r.addVerticalRule(row-delta, col, &wce)
		}
	}
}
