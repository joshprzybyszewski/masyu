package solve

import "github.com/joshprzybyszewski/masyu/model"

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

	if nh == 2 || nv == 2 || ah == 2 || av == 2 {
		// black nodes require at most 1 line or at most 1 avoid in each direction.
		// Write out state that is invalid.
		s.lineHor(r.row, r.col)
		s.avoidHor(r.row, r.col)
	}
}
