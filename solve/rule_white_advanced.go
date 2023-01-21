package solve

import "github.com/joshprzybyszewski/masyu/model"

func newAdvancedHorizontalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkAdvancedHorizontalWhite
	return r
}

func (r *rule) checkAdvancedHorizontalWhite(
	s *state,
) {
	if !s.horLineAt(r.row, r.col) {
		return
	}

	if r.col < 2 {
		// we can't have a horizontal line for white
		// along the left side.
		// error state
		s.lineVer(r.row, r.col)
		return
	}

	l, a := s.horAt(r.row, r.col-2)
	if l {
		s.avoidHor(r.row, r.col+1)
	} else if !a && s.horLineAt(r.row, r.col+1) {
		s.avoidHor(r.row, r.col-2)
	}
}

func newAdvancedVerticalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkAdvancedVerticalWhite
	return r
}

func (r *rule) checkAdvancedVerticalWhite(
	s *state,
) {
	if !s.verLineAt(r.row, r.col) {
		return
	}

	if r.row < 2 {
		// we can't have a vertical line for white
		// along the top side.
		// error state
		s.lineHor(r.row, r.col)
		return
	}

	l, a := s.verAt(r.row-2, r.col)
	if l {
		s.avoidVer(r.row+1, r.col)
	} else if !a && s.verLineAt(r.row+1, r.col) {
		s.avoidVer(r.row-2, r.col)
	}
}
