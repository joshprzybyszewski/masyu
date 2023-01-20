package solve

import "github.com/joshprzybyszewski/masyu/model"

func (r *rule) setHorizontalWhite(
	s *state,
) {
	s.lineHor(r.row, r.col-1)
	s.lineHor(r.row, r.col)
	s.avoidVer(r.row-1, r.col)
	s.avoidVer(r.row, r.col)

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

func (r *rule) setVerticalWhite(
	s *state,
) {
	s.lineVer(r.row-1, r.col)
	s.lineVer(r.row, r.col)
	s.avoidHor(r.row, r.col-1)
	s.avoidHor(r.row, r.col)

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
	} else if !a {
		l, _ = s.verAt(r.row+1, r.col)
		if l {
			s.avoidVer(r.row-2, r.col)
		}
	}
}

func newWhiteL1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteL1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteL1(
	s *state,
) {
	ll, la := s.horAt(r.row, r.col-1)
	if ll {
		r.setHorizontalWhite(s)
	} else if la {
		r.setVerticalWhite(s)
	}
}

func newWhiteR1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteR1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteR1(
	s *state,
) {
	rl, ra := s.horAt(r.row, r.col)
	if rl {
		r.setHorizontalWhite(s)
	} else if ra {
		r.setVerticalWhite(s)
	}
}

func newWhiteU1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteU1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteU1(
	s *state,
) {
	ul, ua := s.verAt(r.row-1, r.col)
	if ul {
		r.setVerticalWhite(s)
	} else if ua {
		r.setHorizontalWhite(s)
	}
}

func newWhiteD1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteD1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteD1(
	s *state,
) {
	dl, da := s.verAt(r.row, r.col)
	if dl {
		r.setVerticalWhite(s)
	} else if da {
		r.setHorizontalWhite(s)
	}
}

/* TODO remove

func (s *state) checkWhite(
	r, c model.Dimension,
) {
	l1, a1 := s.horAt(r, c-1)
	l2, a2 := s.horAt(r, c)
	hl := l1 || l2
	vl := a1 || a2

	l1, a1 = s.verAt(r-1, c)
	l2, a2 = s.verAt(r, c)
	hl = hl || a1 || a2
	vl = vl || l1 || l2

	if hl {
		s.lineHor(r, c-1)
		s.lineHor(r, c)
		s.avoidVer(r-1, c)
		s.avoidVer(r, c)
		if c < 2 {
			// error state
			s.lineVer(r, c)
			return
		}
		l1, a1 = s.horAt(r, c-2)
		if l1 {
			s.avoidHor(r, c+1)
		} else if !a1 {
			l2, _ = s.horAt(r, c+1)
			if l2 {
				s.avoidHor(r, c-2)
			}
		}
	} else if vl {
		s.avoidHor(r, c-1)
		s.avoidHor(r, c)
		s.lineVer(r-1, c)
		s.lineVer(r, c)
		if r < 2 {
			// error state
			s.lineHor(r, c)
			return
		}
		l1, a1 = s.verAt(r-2, c)
		if l1 {
			s.avoidVer(r+1, c)
		} else if !a1 {
			l2, _ = s.verAt(r+1, c)
			if l2 {
				s.avoidVer(r-2, c)
			}
		}
	}
}
*/
