package solve

import "github.com/joshprzybyszewski/masyu/model"

func newWhiteValidator(
	row, col model.Dimension,
) rule {
	r := rule{
		row: row,
		col: col,
	}
	r.check = r.checkWhiteValid
	return r
}

func (r *rule) checkWhiteValid(
	s *state,
) {

	var nh, ah, nv, av uint8

	l, a := s.horAt(r.row, r.col)
	if l {
		nh++
	} else if a {
		ah++
	}

	l, a = s.verAt(r.row, r.col)
	if l {
		nv++
	} else if a {
		av++
	}

	l, a = s.horAt(r.row, r.col-1)
	if l {
		nh++
	} else if a {
		ah++
	}

	l, a = s.verAt(r.row-1, r.col)
	if l {
		nv++
	} else if a {
		av++
	}

	if (nh == 1 && nv == 1) || (ah == 1 && av == 1) {
		// white nodes require 2 lines or 2 avoids in the same direction.
		// Write out state that is invalid.
		r.setInvalid(s)
		return
	}
	// TODO
	return

	if nh > 0 {
		if r.col > 1 &&
			s.horLineAt(r.row, r.col-2) &&
			s.horLineAt(r.row, r.col+1) {
			r.setInvalid(s)
		}
	} else if nv > 0 {
		if r.row > 1 &&
			s.verLineAt(r.row-2, r.col) &&
			s.verLineAt(r.row+1, r.col) {
			r.setInvalid(s)
		}
	}

	if nh+ah+nv+av == 4 {
		if !(((nh == 2 && nv == 0) ||
			(nh == 0 && nv == 2)) &&
			((ah == 2 && av == 0) ||
				(ah == 0 && av == 2))) {
			// the node is claiming that it's solved, but it ain't
			// white nodes require 2 lines or 2 avoids in the same direction.
			// Write out state that is invalid.
			r.setInvalid(s)
		}
	}
}
