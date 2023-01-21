package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackL2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkBlackL2
	return r
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
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkBlackR2
	return r
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
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkBlackU2
	return r
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
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkBlackD2
	return r
}

func (r *rule) checkBlackD2(
	s *state,
) {
	_, a := s.verAt(r.row+1, r.col)
	if a {
		s.avoidVer(r.row, r.col)
	}
}
