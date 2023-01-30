package solve

import "github.com/joshprzybyszewski/masyu/model"

// Thanks [wikipedia](https://en.wikipedia.org/wiki/Masyu#Solution_methods)

func newPairWhiteHorizontalRule(
	row, leftCol model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     row,
		col:     leftCol,
	}
	r.check = r.checkPairWhiteHorizontal
	return r
}

func (r *rule) checkPairWhiteHorizontal(
	s *state,
) {
	// In either case:
	// * ? * ? W ? W ? * - *
	// * - * ? W ? W ? * ? *
	// we need to set the white nodes to go vertically
	if s.horLineAt(r.row, r.col+2) || (r.col > 1 && s.horLineAt(r.row, r.col-2)) {
		s.avoidHor(r.row, r.col)
	}
}

func newPairWhiteVerticalRule(
	topRow, col model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     topRow,
		col:     col,
	}
	r.check = r.checkPairWhiteVertical
	return r
}

func (r *rule) checkPairWhiteVertical(
	s *state,
) {
	// In either case:
	// *   *
	// ?   |
	// *   *
	// ?   ?
	// W   W
	// ?   ?
	// W   W
	// ?   ?
	// *   *
	// |   ?
	// *   *
	// we need to set the white nodes to go horizontally
	if s.verLineAt(r.row+2, r.col) || (r.row > 1 && s.verLineAt(r.row-2, r.col)) {
		s.avoidVer(r.row, r.col)
	}
}
