package solve

import "github.com/joshprzybyszewski/masyu/model"

func newInvertHorizontalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkInvertHorizontalWhite
	return r
}

func (r *rule) checkInvertHorizontalWhite(
	s *state,
) {
	if !s.horLineAt(r.row, r.col+1) ||
		!s.horLineAt(r.row, r.col-2) {
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
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkInvertVerticalWhite
	return r
}

func (r *rule) checkInvertVerticalWhite(
	s *state,
) {
	if !s.verLineAt(r.row+1, r.col) ||
		!s.verLineAt(r.row-2, r.col) {
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
