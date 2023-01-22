package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type pair struct {
	a model.Coord
	b model.Coord

	numSeenNodes int
}

type pathCollector struct {
	pairs [model.MaxPointsPerLine][model.MaxPointsPerLine]*pair

	nodes [model.MaxPointsPerLine]uint64

	hasCycle       bool
	cycleSeenNodes int
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

	if l == nil && r == nil {
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

		pc.pairs[mya.Row][mya.Col] = &p
		pc.pairs[myb.Row][myb.Col] = &p
		return
	}

	if l != nil && r != nil {
		if l == r {
			pc.pairs[mya.Row][mya.Col] = nil
			pc.pairs[myb.Row][myb.Col] = nil
			pc.hasCycle = true
			pc.cycleSeenNodes = l.numSeenNodes
			return
		}

		p := *l
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
		pc.pairs[mya.Row][mya.Col] = nil
		pc.pairs[myb.Row][myb.Col] = nil
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	if l != nil {
		p := *l
		if pc.isNode(myb) {
			p.numSeenNodes++
		}
		pc.pairs[mya.Row][mya.Col] = nil
		if p.a == mya {
			p.a = myb
			pc.pairs[p.a.Row][p.a.Col] = &p
			pc.pairs[p.b.Row][p.b.Col] = &p
			return
		}

		p.b = myb
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	p := *r
	if pc.isNode(mya) {
		p.numSeenNodes++
	}
	pc.pairs[myb.Row][myb.Col] = nil
	if p.a == myb {
		p.a = mya
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	p.b = mya
	pc.pairs[p.a.Row][p.a.Col] = &p
	pc.pairs[p.b.Row][p.b.Col] = &p
}
