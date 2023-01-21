package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type pair struct {
	a model.Coord
	b model.Coord
}

type pathCollector struct {
	pairs [model.MaxPointsPerLine][model.MaxPointsPerLine]*pair

	hasCycle bool
}

func newPathCollector() pathCollector {
	return pathCollector{
		//
	}
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
		// fmt.Printf("first add %v %v\n", mya, myb)
		p := pair{
			a: mya,
			b: myb,
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
			return
		}
		// fmt.Printf("connecting %v %v\n", mya, myb)
		p := *l
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
		pc.pairs[mya.Row][mya.Col] = nil
		if p.a == mya {
			if p.b == myb {
				// fmt.Printf("cycle detected %v %v\n", mya, myb)
				pc.hasCycle = true
				pc.pairs[myb.Row][myb.Col] = nil
				return
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
			// fmt.Printf("cycle detected %v %v\n", mya, myb)
			pc.hasCycle = true
			pc.pairs[myb.Row][myb.Col] = nil
			return
		}
		p.b = myb
		pc.pairs[p.a.Row][p.a.Col] = &p
		pc.pairs[p.b.Row][p.b.Col] = &p
		return
	}

	// fmt.Printf("extending right %v %v\n", mya, myb)
	p := *r
	pc.pairs[myb.Row][myb.Col] = nil
	if p.a == myb {
		if p.b == mya {
			// fmt.Printf("cycle detected %v %v\n", mya, myb)
			pc.hasCycle = true
			pc.pairs[mya.Row][mya.Col] = nil
			return
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
		// fmt.Printf("cycle detected %v %v\n", mya, myb)
		pc.hasCycle = true
		pc.pairs[mya.Row][mya.Col] = nil
		return
	}
	p.b = mya
	pc.pairs[p.a.Row][p.a.Col] = &p
	pc.pairs[p.b.Row][p.b.Col] = &p

}
