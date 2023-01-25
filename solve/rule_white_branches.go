package solve

import "github.com/joshprzybyszewski/masyu/model"

func newWhiteHorizontalBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkWhiteHorizontalBranch
	return r
}

func (r *rule) checkWhiteHorizontalBranch(
	s *state,
) {
	// TODO [MICRO] this could be minutely faster from manual bitmasking
	if !s.verAvoidAt(r.row, r.col-1) ||
		!s.verAvoidAt(r.row-1, r.col-1) ||
		!s.verAvoidAt(r.row, r.col+1) ||
		!s.verAvoidAt(r.row-1, r.col+1) {
		return
	}
	// all four possible branches for a horizontal white node are avoided.
	// Therefore, this must be a vertical node
	r.setVerticalWhite(s)
}

func newWhiteVerticalBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkWhiteVerticalBranch
	return r
}

func (r *rule) checkWhiteVerticalBranch(
	s *state,
) {
	// TODO [MICRO] this could be minutely faster from manual bitmasking
	if !s.horAvoidAt(r.row-1, r.col) ||
		!s.horAvoidAt(r.row-1, r.col-1) ||
		!s.horAvoidAt(r.row+1, r.col) ||
		!s.horAvoidAt(r.row+1, r.col-1) {
		return
	}
	// all four possible branches for a vertical white node are avoided.
	// Therefore, this must be a horizontal node
	r.setHorizontalWhite(s)
}
