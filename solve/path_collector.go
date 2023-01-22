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
			if l.numSeenNodes != r.numSeenNodes {
				panic(`wtf`)
			}
			pc.cycleSeenNodes = l.numSeenNodes
			return
		}

		// fmt.Printf("connecting %v %v\n", mya, myb)
		p := *l
		p.numSeenNodes += r.numSeenNodes
		if p.a == mya {
			if r.a == myb {
				p.a = r.b
			} else {
				if r.b != myb {
					panic(`dev error`)
				}
				p.a = r.a
			}
		} else {
			if p.b != mya {
				panic(`dev error`)
			}
			if r.a == myb {
				p.b = r.b
			} else {
				if r.b != myb {
					panic(`dev error`)
				}
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
		// fmt.Printf("extending left %v %v\n", mya, myb)
		p := *l
		if pc.isNode(myb) {
			p.numSeenNodes++
		}
		pc.pairs[mya.Row][mya.Col] = nil
		if p.a == mya {
			if p.b == myb {
				panic(`ahh`)
				// fmt.Printf("cycle detected %v %v\n", mya, myb)
				// pc.hasCycle = true
				// pc.cycleSeenNodes = l.numSeenNodes
				// pc.pairs[myb.Row][myb.Col] = nil
				// return
			}
			p.a = myb
			pc.pairs[p.a.Row][p.a.Col] = &p
			pc.pairs[p.b.Row][p.b.Col] = &p
			return
		}
		if p.b != mya {
			panic(`dev error`)
		}

		if p.a == myb {
			panic(`ahh`)
			// fmt.Printf("cycle detected %v %v\n", mya, myb)
			// pc.hasCycle = true
			// pc.cycleSeenNodes = l.numSeenNodes
			// pc.pairs[myb.Row][myb.Col] = nil
			// return
		}
		p.b = myb
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	// fmt.Printf("extending right %v %v\n", mya, myb)
	p := *r
	if pc.isNode(mya) {
		p.numSeenNodes++
	}
	pc.pairs[myb.Row][myb.Col] = nil
	if p.a == myb {
		if p.b == mya {
			panic(`ahh`)
			// fmt.Printf("cycle detected %v %v\n", mya, myb)
			// pc.hasCycle = true
			// pc.cycleSeenNodes = l.numSeenNodes
			// pc.pairs[mya.Row][mya.Col] = nil
			// return
		}
		p.a = mya
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	if p.b != myb {
		panic(`dev error`)
	}
	if p.a == mya {
		panic(`ahh`)
		// fmt.Printf("cycle detected %v %v\n", mya, myb)
		// pc.hasCycle = true
		// pc.cycleSeenNodes = l.numSeenNodes
		// pc.pairs[mya.Row][mya.Col] = nil
		// return
	}
	p.b = mya
	pc.pairs[p.a.Row][p.a.Col] = &p
	pc.pairs[p.b.Row][p.b.Col] = &p

}
