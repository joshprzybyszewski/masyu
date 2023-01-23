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
		s.lineHor(r.row, r.col)
		s.avoidHor(r.row, r.col)
	}
}
