package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackL2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 2,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackL2
	return r
}

func (r *rule) checkBlackL2(
	s *state,
) {
	if s.horAvoidAt(r.row, r.col-2) {
		s.avoidHor(r.row, r.col-1)
		s.lineHor(r.row, r.col)
	}
}

func newBlackR2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 2,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackR2
	return r
}

func (r *rule) checkBlackR2(
	s *state,
) {
	if s.horAvoidAt(r.row, r.col+1) {
		s.avoidHor(r.row, r.col)
		s.lineHor(r.row, r.col-1)
	}
}

func newBlackU2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 2,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackU2
	return r
}

func (r *rule) checkBlackU2(
	s *state,
) {
	if s.verAvoidAt(r.row-2, r.col) {
		s.avoidVer(r.row-1, r.col)
		s.lineVer(r.row, r.col)
	}
}

func newBlackD2Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 2,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackD2
	return r
}

func (r *rule) checkBlackD2(
	s *state,
) {
	if s.verAvoidAt(r.row+1, r.col) {
		s.avoidVer(r.row, r.col)
		s.lineVer(r.row-1, r.col)
	}
}
