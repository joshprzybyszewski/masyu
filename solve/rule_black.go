package solve

import "github.com/joshprzybyszewski/masyu/model"

// TODO add advanced black node rules:
// if the non-adjacent path that is avoided,
//   then the adjacent path is avoided.
// if there is a line branching off the adjacent path,
//   then the adjacent path is avoided.

func newBlackL1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackL1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackL1(
	s *state,
) {
	ll, la := s.horAt(r.row, r.col-1)
	if ll {
		s.lineHor(r.row, r.col-2)
		s.avoidHor(r.row, r.col)
		s.avoidVer(r.row-1, r.col-1)
		s.avoidVer(r.row, r.col-1)
	} else if la {
		s.lineHor(r.row, r.col)
		s.lineHor(r.row, r.col+1)
		s.avoidVer(r.row-1, r.col+1)
		s.avoidVer(r.row, r.col+1)
	}
}

func newBlackR1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackR1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackR1(
	s *state,
) {
	rl, ra := s.horAt(r.row, r.col)
	if rl {
		s.lineHor(r.row, r.col+1)
		s.avoidHor(r.row, r.col-1)
		s.avoidVer(r.row-1, r.col+1)
		s.avoidVer(r.row, r.col+1)
	} else if ra {
		s.lineHor(r.row, r.col-1)
		s.lineHor(r.row, r.col-2)
		s.avoidVer(r.row-1, r.col-1)
		s.avoidVer(r.row, r.col-1)
	}
}

func newBlackU1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackU1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackU1(
	s *state,
) {
	ul, ua := s.verAt(r.row-1, r.col)
	if ul {
		s.lineVer(r.row-2, r.col)
		s.avoidVer(r.row, r.col)
		s.avoidHor(r.row-1, r.col-1)
		s.avoidHor(r.row-1, r.col)
	} else if ua {
		s.lineVer(r.row, r.col)
		s.lineVer(r.row+1, r.col)
		s.avoidHor(r.row+1, r.col-1)
		s.avoidHor(r.row+1, r.col)
	}
}

func newBlackD1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackD1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackD1(
	s *state,
) {
	dl, da := s.verAt(r.row, r.col)
	if dl {
		s.lineVer(r.row+1, r.col)
		s.avoidVer(r.row-1, r.col)
		s.avoidHor(r.row+1, r.col-1)
		s.avoidHor(r.row+1, r.col)
	} else if da {
		s.lineVer(r.row-1, r.col)
		s.lineVer(r.row-2, r.col)
		s.avoidHor(r.row-1, r.col)
		s.avoidHor(r.row-1, r.col-1)
	}

}
