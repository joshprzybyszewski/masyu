package solve

import "github.com/joshprzybyszewski/masyu/model"

func newAdvancedHorizontalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteAdvancedHorizontalRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkAdvancedHorizontalWhite(
	s *state,
) {
	l, _ := s.horAt(r.row, r.col)
	if !l {
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
	} else if !a {
		l, _ = s.horAt(r.row, r.col+1)
		if l {
			s.avoidHor(r.row, r.col-2)
		}
	}
}

func newAdvancedVerticalWhite(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteAdvancedVerticalRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkAdvancedVerticalWhite(
	s *state,
) {
	l, _ := s.verAt(r.row, r.col)
	if !l {
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
		return
	}

	if !a {
		l, _ = s.verAt(r.row+1, r.col)
		if l {
			s.avoidVer(r.row-2, r.col)
		}
	}
}
