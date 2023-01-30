package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackLBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackLBranch
	return r
}

func (r *rule) checkBlackLBranch(
	s *state,
) {
	if !s.verLineAt(r.row, r.col-1) ||
		!s.verLineAt(r.row-1, r.col-1) {
		return
	}
	s.avoidHor(r.row, r.col-1)
}

func newBlackRBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackRBranch
	return r
}

func (r *rule) checkBlackRBranch(
	s *state,
) {
	if !s.verLineAt(r.row, r.col+1) ||
		!s.verLineAt(r.row-1, r.col+1) {
		return
	}
	s.avoidHor(r.row, r.col)
}

func newBlackUBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackUBranch
	return r

}

func (r *rule) checkBlackUBranch(
	s *state,
) {
	if !s.horLineAt(r.row-1, r.col) ||
		!s.horLineAt(r.row-1, r.col-1) {
		return
	}

	s.avoidVer(r.row-1, r.col)
}

func newBlackDBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 1,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkBlackDBranch
	return r
}

func (r *rule) checkBlackDBranch(
	s *state,
) {
	if !s.horLineAt(r.row+1, r.col) ||
		!s.horLineAt(r.row+1, r.col-1) {
		return
	}

	s.avoidVer(r.row, r.col)
}
