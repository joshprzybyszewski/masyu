package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func newBlackValidator(
	row, col model.Dimension,
) rule {
	r := rule{
		affects: 4,
		row:     row,
		col:     col,
	}
	r.check = r.checkBlackValid
	return r
}

func (r *rule) checkBlackValid(
	s *state,
) {

	l, a := s.horAt(r.row, r.col)
	if l {
		s.avoidHor(r.row, r.col-1)
	} else if a {
		s.lineHor(r.row, r.col-1)
	} else {
		// only check the opposite side if we don't know what we are
		l, a = s.horAt(r.row, r.col-1)
		if l {
			s.avoidHor(r.row, r.col)
		} else if a {
			s.lineHor(r.row, r.col)
		}
	}

	l, a = s.verAt(r.row, r.col)
	if l {
		s.avoidVer(r.row-1, r.col)
	} else if a {
		s.lineVer(r.row-1, r.col)
	} else {
		// only check the opposite side if we don't know what we are.
		l, a = s.verAt(r.row-1, r.col)
		if l {
			s.avoidVer(r.row, r.col)
		} else if a {
			s.lineVer(r.row, r.col)
		}
	}
}
