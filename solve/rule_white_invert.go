package solve

import "github.com/joshprzybyszewski/masyu/model"

func newInvertHorizontalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteInvertHorizontalRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkInvertHorizontalWhite(
	s *state,
) {
	l, _ := s.horAt(r.row, r.col+1)
	if !l {
		return
	}
	l, _ = s.horAt(r.row, r.col-2)
	if !l {
		return
	}

	// In this case:
	// * - * ? W ? * - *
	// We know that the white node cannot be horizontal.

	r.setVerticalWhite(s)
}

func newInvertVerticalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteInvertVerticalRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkInvertVerticalWhite(
	s *state,
) {
	l, _ := s.verAt(r.row+1, r.col)
	if !l {
		return
	}
	l, _ = s.verAt(r.row-2, r.col)
	if !l {
		return
	}

	// In this case:
	// *
	// |
	// *
	// ?
	// W
	// ?
	// *
	// |
	// *
	// We know that the white node cannot be vertical.

	r.setHorizontalWhite(s)
}
