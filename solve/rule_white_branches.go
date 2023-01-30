package solve

import "github.com/joshprzybyszewski/masyu/model"

func newWhiteHorizontalBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 4,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkWhiteHorizontalBranch
	return r
}

func (r *rule) checkWhiteHorizontalBranch(
	s *state,
) {
	cannotBranchLeft := (r.col > 1 && s.horLineAt(r.row, r.col-2)) ||
		(s.verAvoidAt(r.row, r.col-1) &&
			s.verAvoidAt(r.row-1, r.col-1))
	cannotBranchRight := s.horLineAt(r.row, r.col+1) ||
		(s.verAvoidAt(r.row, r.col+1) &&
			s.verAvoidAt(r.row-1, r.col+1))

	if cannotBranchLeft && cannotBranchRight {
		// all four possible branches for a horizontal white node are impossible.
		// Therefore, this must be a vertical node
		s.lineVer(r.row-1, r.col)
		s.lineVer(r.row, r.col)
		s.avoidHor(r.row, r.col-1)
		s.avoidHor(r.row, r.col)
	}
}

func newWhiteVerticalBranchRule(
	nodeRow, nodeCol model.Dimension,
) rule {
	r := rule{
		affects: 4,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.checkWhiteVerticalBranch
	return r
}

func (r *rule) checkWhiteVerticalBranch(
	s *state,
) {
	cannotBranchUp := (r.row > 1 && s.verLineAt(r.row-2, r.col)) ||
		(s.horAvoidAt(r.row-1, r.col) &&
			s.horAvoidAt(r.row-1, r.col-1))
	cannotBranchDown := s.verLineAt(r.row+1, r.col) ||
		(s.horAvoidAt(r.row+1, r.col) &&
			s.horAvoidAt(r.row+1, r.col-1))

	if cannotBranchUp && cannotBranchDown {
		// all four possible branches for a vertical white node are impossible.
		// Therefore, this must be a horizontal node
		s.lineHor(r.row, r.col-1)
		s.lineHor(r.row, r.col)
		s.avoidVer(r.row-1, r.col)
		s.avoidVer(r.row, r.col)
	}
}
