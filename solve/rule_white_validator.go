package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func newWhiteValidator(
	row, col model.Dimension,
) rule {
	r := rule{
		affects: 5,
		row:     row,
		col:     col,
	}
	r.check = r.checkWhiteValid
	return r
}

func (r *rule) checkWhiteValid(
	s *state,
) {

	h, v := s.horAt(r.row, r.col)

	if !h && !v {
		v, h = s.verAt(r.row, r.col)

		if !h && !v {
			h, v = s.horAt(r.row, r.col-1)

			if !h && !v {
				v, h = s.verAt(r.row-1, r.col)
			}
		}
	}

	if h {
		s.lineHor(r.row, r.col-1)
		s.lineHor(r.row, r.col)
		s.avoidVer(r.row-1, r.col)
		s.avoidVer(r.row, r.col)

		if r.col > 1 {
			if s.horLineAt(r.row, r.col+1) {
				s.avoidHor(r.row, r.col-2)
			} else if s.horLineAt(r.row, r.col-2) {
				s.avoidHor(r.row, r.col+1)
			}
		}
	}
	if v {
		s.lineVer(r.row-1, r.col)
		s.lineVer(r.row, r.col)
		s.avoidHor(r.row, r.col-1)
		s.avoidHor(r.row, r.col)

		if r.row > 1 {
			if s.verLineAt(r.row+1, r.col) {
				s.avoidVer(r.row-2, r.col)
			} else if s.verLineAt(r.row-2, r.col) {
				s.avoidVer(r.row+1, r.col)
			}
		}
	}
}
