package solve

import "github.com/joshprzybyszewski/masyu/model"

func newDefaultRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row: row,
		col: col,
	}
	r.check = r.checkDefault
	return r
}

func (r *rule) checkDefault(
	s *state,
) {

	var nl, na uint8

	rl, ra := s.horAt(r.row, r.col)
	if rl {
		nl++
	} else if ra {
		na++
	}

	dl, da := s.verAt(r.row, r.col)
	if dl {
		nl++
	} else if da {
		na++
	}

	ll, la := s.horAt(r.row, r.col-1)
	if ll {
		nl++
		if nl >= 3 {
			// We should never have 3 or 4 lines coming into a pin.
			// Write out state that is invalid.
			s.lineHor(r.row, r.col)
			s.avoidHor(r.row, r.col)
			return
		}
	} else if la {
		na++
	}

	ul, ua := s.verAt(r.row-1, r.col)
	if ul {
		nl++
		if nl >= 3 {
			// We should never have 3 or 4 lines coming into a pin.
			// Write out state that is invalid.
			s.lineHor(r.row, r.col)
			s.avoidHor(r.row, r.col)
			return
		}
	} else if ua {
		na++
	}

	if nl+na == 4 && nl == 1 {
		// All four directions are set, but there's only one line.
		// Write out state that is invalid.
		s.lineHor(r.row, r.col)
		s.avoidHor(r.row, r.col)
		return
	}

	if na == 3 {
		// there is one line left to be drawn: it's an avoid
		if !ll && !la {
			s.avoidHor(r.row, r.col-1)
		} else if !ul && !ua {
			s.avoidVer(r.row-1, r.col)
		} else if !rl && !ra {
			s.avoidHor(r.row, r.col)
		} else if !dl && !da {
			s.avoidVer(r.row, r.col)
		}
		return
	}
	if nl == 1 && na == 2 {
		// there is one line left to be drawn: it's a line
		if !ll && !la {
			s.lineHor(r.row, r.col-1)
		} else if !ul && !ua {
			s.lineVer(r.row-1, r.col)
		} else if !rl && !ra {
			s.lineHor(r.row, r.col)
		} else if !dl && !da {
			s.lineVer(r.row, r.col)
		}
		return
	}

	if nl != 2 {
		// otherwise there's not enough information to make a change
		return
	}

	if !rl && !ra {
		s.avoidHor(r.row, r.col)
	}
	if !dl && !da {
		s.avoidVer(r.row, r.col)
	}
	if !ll && !la {
		s.avoidHor(r.row, r.col-1)
	}
	if !ul && !ua {
		s.avoidVer(r.row-1, r.col)
	}
}
