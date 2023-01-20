package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackL2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackL2RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackL2(
	s *state,
) {
	_, a := s.horAt(r.row, r.col-2)
	if a {
		s.avoidHor(r.row, r.col-1)
	}
}

func newBlackR2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackR2RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackR2(
	s *state,
) {
	_, a := s.horAt(r.row, r.col+1)
	if a {
		s.avoidHor(r.row, r.col)
	}
}

func newBlackU2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackU2RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackU2(
	s *state,
) {
	_, a := s.verAt(r.row-2, r.col)
	if a {
		s.avoidVer(r.row-1, r.col)
	}
}

func newBlackD2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackD2RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackD2(
	s *state,
) {
	_, a := s.verAt(r.row+1, r.col)
	if a {
		s.avoidVer(r.row, r.col)
	}
}
