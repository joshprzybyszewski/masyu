package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func newBlackValidator(
	row, col model.Dimension,
) rule {
	r := rule{
		row: row,
		col: col,
	}
	r.check = r.checkBlackValid
	return r
}

func (r *rule) checkBlackValid(
	s *state,
) {

	l, a := s.horAt(r.row, r.col)
	if l {
		s.lineHor(r.row, r.col+1)
		s.avoidHor(r.row, r.col-1)
	} else if a || s.horLineAt(r.row, r.col-1) {
		s.avoidHor(r.row, r.col)
		s.lineHor(r.row, r.col-1)
		if r.col > 1 {
			s.lineHor(r.row, r.col-2)
		} else {
			r.setInvalid(s)
		}
	}

	l, a = s.verAt(r.row, r.col)
	if l {
		s.lineVer(r.row+1, r.col)
		s.avoidVer(r.row-1, r.col)
	} else if a || s.verLineAt(r.row-1, r.col) {
		s.avoidVer(r.row, r.col)
		s.lineVer(r.row-1, r.col)
		if r.row > 1 {
			s.lineVer(r.row-2, r.col)
		} else {
			r.setInvalid(s)
		}
	}
}
