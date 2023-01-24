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
	return !p.isEmpty() &&
		p.a.Row == p.b.Row &&
		(p.a.Col == p.b.Col+1 || p.a.Col == p.b.Col-1)
}

func (p *pair) isVerticallyClose() bool {
	return !p.isEmpty() &&
		p.a.Col == p.b.Col &&
		(p.a.Row == p.b.Row+1 || p.a.Row == p.b.Row-1)
}

type pathCollector struct {
	pairs [model.MaxPointsPerLine][model.MaxPointsPerLine]pair

	nodes [model.MaxPointsPerLine]uint64

	hasCycle       bool
	cycleSeenNodes int

	hasInteresting bool
}

func newPathCollector(
	nodes []model.Node,
) pathCollector {
	pc := pathCollector{}

	for _, n := range nodes {
		pc.nodes[n.Row] |= (n.Col.Bit())
	}

	return pc
}

func (pc *pathCollector) getInteresting(
	s *state,
) (model.Coord, bool, bool) {
	var c model.Coord

	if !pc.hasInteresting {
		return c, false, false
	}

	size := model.Dimension(s.size)
	var l, a bool

	for c.Row = model.Dimension(1); c.Row < size; c.Row++ {
		for c.Col = model.Dimension(1); c.Col <= size; c.Col++ {
			if pc.pairs[c.Row][c.Col].isVerticallyClose() {
				if l, a = s.verAt(c.Row, c.Col); !l && !a {
					return c, false, true
				}
			}
		}
	}

	for c.Row = model.Dimension(1); c.Row <= size; c.Row++ {
		for c.Col = model.Dimension(1); c.Col < size; c.Col++ {
			if pc.pairs[c.Row][c.Col].isHorizontallyClose() {
				if l, a = s.horAt(c.Row, c.Col); !l && !a {
					return c, true, true
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
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row,
		Col: col + 1,
	}

	pc.add(mya, myb)
}

func (pc *pathCollector) addVertical(
	row, col model.Dimension,
) {
	mya := model.Coord{
		Row: row,
		Col: col,
	}
	myb := model.Coord{
		Row: row + 1,
		Col: col,
	}
	pc.add(mya, myb)
}

func (pc *pathCollector) add(
	mya, myb model.Coord,
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
}
