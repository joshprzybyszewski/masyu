package solve

import "github.com/joshprzybyszewski/masyu/model"

func newDefaultRule(
	row, col model.Dimension,
) rule {
	return rule{
		kind: defaultRuleKind,
		row:  row,
		col:  col,
	}
}

func (r *rule) checkDefault(
	s *state,
) {
	rl, ra := s.horAt(r.row, r.col)
	dl, da := s.verAt(r.row, r.col)
	ll, la := s.horAt(r.row, r.col-1)
	ul, ua := s.verAt(r.row-1, r.col)

	var nl, na, dir uint8
	if rl {
		nl++
		dir |= 1
	}
	if ra {
		na++
		dir |= 1
	}
	if dl {
		nl++
		dir |= 1 << 1
	}
	if da {
		na++
		dir |= 1 << 1
	}
	if ll {
		nl++
		dir |= 1 << 2
	}
	if la {
		na++
		dir |= 1 << 2
	}
	if ul {
		nl++
		dir |= 1 << 3
	}
	if ua {
		na++
		dir |= 1 << 3
	}

	if nl != 2 && nl+na != 3 {
		return
	}

	if dir&1 == 0 {
		if nl == 1 {
			s.lineHor(r.row, r.col)
		} else {
			s.avoidHor(r.row, r.col)
		}
	}
	if dir&(1<<1) == 0 {
		if nl == 1 {
			s.lineVer(r.row, r.col)
		} else {
			s.avoidVer(r.row, r.col)
		}
	}
	if dir&(1<<2) == 0 {
		if nl == 1 {
			s.lineHor(r.row, r.col-1)
		} else {
			s.avoidHor(r.row, r.col-1)
		}
	}
	if dir&(1<<3) == 0 {
		if nl == 1 {
			s.lineVer(r.row-1, r.col)
		} else {
			s.avoidVer(r.row-1, r.col)
		}
	}
}
