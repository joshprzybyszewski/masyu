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
	if s.horLineAt(r.row, r.col+1) &&
		s.horLineAt(r.row, r.col-2) {
		// In this case:
		// * - * ? W ? * - *
		// We know that the white node cannot be horizontal.
		s.avoidHor(r.row, r.col)
		s.avoidHor(r.row, r.col-1)
	}
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
	if s.verLineAt(r.row+1, r.col) &&
		s.verLineAt(r.row-2, r.col) {
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

		s.avoidVer(r.row-1, r.col)
		s.avoidVer(r.row, r.col)
	}
}
