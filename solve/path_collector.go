package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type pair struct {
	a model.Coord
	b model.Coord

	numSeenNodes int
}

func newEmptyPair() pair {
	return pair{}
}

func (p *pair) isEmpty() bool {
	return p.a.Col == 0
}

func (p *pair) isHorizontallyClose() bool {
	return p.a.Row == p.b.Row &&
		(p.a.Col == p.b.Col+1 || p.a.Col == p.b.Col-1)
}

func (p *pair) isVerticallyClose() bool {
	return p.a.Col == p.b.Col &&
		(p.a.Row == p.b.Row+1 || p.a.Row == p.b.Row-1)
}

type pathCollector struct {
	pairs [model.MaxPointsPerLine][model.MaxPointsPerLine]pair

	nodes [model.MaxPointsPerLine]model.DimensionBit

	hasCycle       bool
	cycleSeenNodes int
}

func newPathCollector(
	nodes []model.Node,
) pathCollector {
	pc := pathCollector{}

	for _, n := range nodes {
		pc.nodes[n.Row] |= n.Col.Bit()
	}

	return pc
}

func (pc *pathCollector) getInteresting(
	s *state,
) (model.Coord, bool, bool) {
	var c model.Coord

	size := model.Dimension(s.size)
	var l, a bool

	for c.Row = model.Dimension(1); c.Row <= size; c.Row++ {
		for c.Col = model.Dimension(1); c.Col <= size; c.Col++ {
			if !pc.pairs[c.Row][c.Col].isEmpty() {
				if !pc.pairs[c.Row+1][c.Col].isEmpty() {
					if l, a = s.verAt(c.Row, c.Col); !l && !a {
						return c, false, true
					}
				}
				if !pc.pairs[c.Row][c.Col+1].isEmpty() {
					if l, a = s.horAt(c.Row, c.Col); !l && !a {
						return c, true, true
					}
				}
			}
		}
	}

	return c, false, false
}

func (pc *pathCollector) isNode(
	c model.Coord,
) bool {
	return pc.nodes[c.Row]&(c.Col.Bit()) != 0
}

func (pc *pathCollector) addHorizontal(
	row, col model.Dimension,
	s *state,
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row,
		Col: col + 1,
	}

	pc.add(
		mya, myb,
		s,
	)
}

func (pc *pathCollector) addVertical(
	row, col model.Dimension,
	s *state,
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row + 1,
		Col: col,
	}
	pc.add(
		mya, myb,
		s,
	)
}

func (pc *pathCollector) add(
	mya, myb model.Coord,
	s *state,
) {
	l := pc.pairs[mya.Row][mya.Col]
	r := pc.pairs[myb.Row][myb.Col]

	if l.isEmpty() && r.isEmpty() {
		p := pair{
			a: mya,
			b: myb,
		}

		if pc.isNode(mya) {
			p.numSeenNodes++
		}
		if pc.isNode(myb) {
			p.numSeenNodes++
		}

		pc.pairs[mya.Row][mya.Col] = p
		pc.pairs[myb.Row][myb.Col] = p
		return
	}

	if !l.isEmpty() && !r.isEmpty() {
		pc.pairs[mya.Row][mya.Col] = newEmptyPair()
		pc.pairs[myb.Row][myb.Col] = newEmptyPair()

		if l == r {
			if pc.hasCycle {
				// a second cycle? this is bad news.
				pc.cycleSeenNodes = -1
				return
			}
			pc.hasCycle = true
			pc.cycleSeenNodes = l.numSeenNodes
			return
		}

		p := l
		p.numSeenNodes += r.numSeenNodes
		if p.a == mya {
			if r.a == myb {
				p.a = r.b
			} else {
				p.a = r.a
			}
		} else {
			if r.a == myb {
				p.b = r.b
			} else {
				p.b = r.a
			}
		}
		pc.pairs[p.a.Row][p.a.Col] = p
		pc.pairs[p.b.Row][p.b.Col] = p
		pc.checkNewPair(
			p,
			s,
		)
		return
	}

	if !l.isEmpty() {
		p := l
		if pc.isNode(myb) {
			p.numSeenNodes++
		}
		pc.pairs[mya.Row][mya.Col] = newEmptyPair()
		if p.a == mya {
			p.a = myb
		} else {
			p.b = myb
		}

		pc.pairs[p.a.Row][p.a.Col] = p
		pc.pairs[p.b.Row][p.b.Col] = p
		pc.checkNewPair(
			p,
			s,
		)
		return
	}

	p := r
	if pc.isNode(mya) {
		p.numSeenNodes++
	}
	pc.pairs[myb.Row][myb.Col] = newEmptyPair()
	if p.a == myb {
		p.a = mya
	} else {
		p.b = mya
	}

	pc.pairs[p.a.Row][p.a.Col] = p
	pc.pairs[p.b.Row][p.b.Col] = p

	pc.checkNewPair(
		p,
		s,
	)
}

func (pc *pathCollector) checkNewPair(
	p pair,
	s *state,
) {
	if pc.hasCycle || s.hasInvalid {
		return
	}
	if p.isEmpty() {
		return
	}

	h := p.isHorizontallyClose()
	if !h && !p.isVerticallyClose() {
		return
	}

	r := p.a.Row
	if p.b.Row < r {
		r = p.b.Row
	}
	c := p.a.Col
	if p.b.Col < c {
		c = p.b.Col
	}

	// only need to check the state when we're about to write a line.
	// re-writing an avoid is no problem.
	if p.numSeenNodes == len(s.nodes) {
		cpy := *s
		if h {
			if !cpy.horAvoidAt(r, c) {
				cpy.lineHor(r, c)
			}
		} else {
			if !cpy.verAvoidAt(r, c) {
				cpy.lineVer(r, c)
			}
		}
		ss := settle(&cpy)
		if ss == solved || ss == validUnsolved {
			settle(s)
			return
		}
	}

	if h {
		s.avoidHor(r, c)
	} else {
		s.avoidVer(r, c)
	}
}
