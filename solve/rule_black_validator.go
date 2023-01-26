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

	var nh, ah, nv, av uint8

	l, a := s.horAt(r.row, r.col)
	if l {
		nh++
		if s.horAvoidAt(r.row, r.col+1) {
			r.setInvalid(s)
			return
		}
	} else if a {
		ah++
	}

	l, a = s.verAt(r.row, r.col)
	if l {
		nv++
		if s.verAvoidAt(r.row+1, r.col) {
			r.setInvalid(s)
			return
		}
	} else if a {
		av++
	}

	l, a = s.horAt(r.row, r.col-1)
	if l {
		if nh == 1 {
			r.setInvalid(s)
			return
		}

		if r.col < 2 || s.horAvoidAt(r.row, r.col-2) {
			r.setInvalid(s)
			return
		}
	} else if a && ah == 1 {
		r.setInvalid(s)
		return
	}

	l, a = s.verAt(r.row-1, r.col)
	if l {
		if nv == 1 {
			r.setInvalid(s)
			return
		}
		if r.row < 2 || s.verAvoidAt(r.row-2, r.col) {
			r.setInvalid(s)
			return
		}
	} else if a && av == 1 {
		r.setInvalid(s)
		return
	}

	// TODO
	// return
	l, a = s.horAt(r.row, r.col)
	if l {
		if s.horAvoidAt(r.row, r.col+1) {
			r.setInvalid(s)
		}
		if s.horLineAt(r.row, r.col-1) {
			r.setInvalid(s)
		}
	} else if a {
		if s.horAvoidAt(r.row, r.col-1) {
			r.setInvalid(s)
		}
	} else if s.horLineAt(r.row, r.col-1) {
		if r.col < 2 || s.horAvoidAt(r.row, r.col-2) {
			r.setInvalid(s)
		}
	}

	l, a = s.verAt(r.row, r.col)
	if l {
		if s.verAvoidAt(r.row+1, r.col) {
			r.setInvalid(s)
		}
		if s.verLineAt(r.row-1, r.col) {
			r.setInvalid(s)
		}
	} else if a {
		if s.verAvoidAt(r.row-1, r.col) {
			r.setInvalid(s)
		}
	} else if s.verLineAt(r.row-1, r.col) {
		if r.row < 2 || s.verAvoidAt(r.row-2, r.col) {
			r.setInvalid(s)
		}
	}
}
