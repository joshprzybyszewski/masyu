package solve

import "github.com/joshprzybyszewski/masyu/model"

func newInvertHorizontalBlack(
	nodeRow, nodeCol model.Dimension,
) *rule {
	if nodeRow < 2 || nodeCol < 2 {
		return nil
	}

	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkInvertHorizontalBlack
	return &r
}

func (r *rule) checkInvertHorizontalBlack(
	s *state,
) {

	// In this case, we are checking the paths at the %:
	// * % * ? B ? * % *

	ll, la := s.horAt(r.row, r.col-2)
	rl, ra := s.horAt(r.row, r.col+1)
	if la && ra {
		// if they are both avoided, then we are invalid.
		r.setInvalid(s)
		return
	}
	if !ll || !rl {
		// we aren't about to infer anything
		return
	}

	// In this case, both paths exist:
	// * - * ? B ? * - *

	if s.verAvoidAt(r.row-2, r.col) {
		// the turn can't go up. Send it down instead.
		s.avoidVer(r.row-1, r.col)
		s.lineVer(r.row, r.col)
		s.lineVer(r.row+1, r.col)
		return
	}

	if s.verAvoidAt(r.row+1, r.col) {
		// the turn can't go down. Send it up instead.
		s.avoidVer(r.row, r.col)
		s.lineVer(r.row-1, r.col)
		s.lineVer(r.row-2, r.col)
		return
	}
}

func newInvertVerticalBlack(
	nodeRow, nodeCol model.Dimension,
) *rule {
	if nodeRow < 2 || nodeCol < 2 {
		return nil
	}

	r := rule{
		row: nodeRow,
		col: nodeCol,
	}
	r.check = r.checkInvertVerticalBlack
	return &r
}

func (r *rule) checkInvertVerticalBlack(
	s *state,
) {
	// In this case, we are checking the paths at the %:
	// *
	// %
	// *
	// ?
	// W
	// ?
	// *
	// %
	// *

	ul, ua := s.verAt(r.row-2, r.col)
	dl, da := s.verAt(r.row+1, r.col)
	if ua && da {
		// if they are both avoided, then we are invalid.
		r.setInvalid(s)
		return
	}
	if !ul || !dl {
		// we aren't about to infer anything
		return
	}

	// In this case, both paths exist:

	if s.horAvoidAt(r.row, r.col-2) {
		// the turn can't go left. Send it right instead.
		s.avoidHor(r.row, r.col-1)
		s.lineHor(r.row, r.col)
		s.lineHor(r.row, r.col+1)
		return
	}

	if s.horAvoidAt(r.row, r.col+1) {
		// the turn can't go right. Send it left instead.
		s.avoidHor(r.row, r.col)
		s.lineHor(r.row, r.col-1)
		s.lineHor(r.row, r.col-2)
		return
	}
}
