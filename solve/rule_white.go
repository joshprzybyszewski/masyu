package solve

import "github.com/joshprzybyszewski/masyu/model"

func (r *rule) setHorizontalWhite(
	s *state,
) {
	s.lineHor(r.row, r.col-1)
	s.lineHor(r.row, r.col)
	s.avoidVer(r.row-1, r.col)
	s.avoidVer(r.row, r.col)

	r.checkAdvancedHorizontalWhite(s)
}

func (r *rule) setVerticalWhite(
	s *state,
) {
	s.lineVer(r.row-1, r.col)
	s.lineVer(r.row, r.col)
	s.avoidHor(r.row, r.col-1)
	s.avoidHor(r.row, r.col)

	r.checkAdvancedVerticalWhite(s)
}

func newWhiteL1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteL1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteL1(
	s *state,
) {
	ll, la := s.horAt(r.row, r.col-1)
	if ll {
		r.setHorizontalWhite(s)
	} else if la {
		r.setVerticalWhite(s)
	}
}

func newWhiteR1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteR1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteR1(
	s *state,
) {
	rl, ra := s.horAt(r.row, r.col)
	if rl {
		r.setHorizontalWhite(s)
	} else if ra {
		r.setVerticalWhite(s)
	}
}

func newWhiteU1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteU1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteU1(
	s *state,
) {
	ul, ua := s.verAt(r.row-1, r.col)
	if ul {
		r.setVerticalWhite(s)
	} else if ua {
		r.setHorizontalWhite(s)
	}
}

func newWhiteD1Rule(
	nodeRow, nodeCol model.Dimension,
) rule {
	return rule{
		kind: whiteD1RuleKind,
		row:  nodeRow,
		col:  nodeCol,
	}
}

func (r *rule) checkWhiteD1(
	s *state,
) {
	dl, da := s.verAt(r.row, r.col)
	if dl {
		r.setVerticalWhite(s)
	} else if da {
		r.setHorizontalWhite(s)
	}
}
