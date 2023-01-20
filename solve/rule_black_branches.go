package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackLBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackLBranchRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackLBranch(
	s *state,
) {
	l, _ := s.verAt(r.row, r.col-1)
	if !l {
		return
	}
	l, _ = s.verAt(r.row-1, r.col-1)
	if !l {
		return
	}
	s.avoidHor(r.row, r.col-1)
}

func newBlackRBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackRBranchRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackRBranch(
	s *state,
) {
	l, _ := s.verAt(r.row, r.col+1)
	if !l {
		return
	}
	l, _ = s.verAt(r.row-1, r.col+1)
	if !l {
		return
	}
	s.avoidHor(r.row, r.col)
}

func newBlackUBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackUBranchRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackUBranch(
	s *state,
) {
	l, _ := s.horAt(r.row-1, r.col)
	if !l {
		return
	}
	l, _ = s.horAt(r.row-1, r.col-1)
	if !l {
		return
	}

	s.avoidVer(r.row-1, r.col)
}

func newBlackDBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: blackDBranchRuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkBlackDBranch(
	s *state,
) {
	l, _ := s.horAt(r.row+1, r.col)
	if !l {
		return
	}
	l, _ = s.horAt(r.row+1, r.col-1)
	if !l {
		return
	}

	s.avoidVer(r.row, r.col)
}
