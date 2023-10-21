package solve

import "github.com/joshprzybyszewski/masyu/model"

// TODO maybe this could be different for different node types?
func newNodeCheckerExpensiveRule(
	nodeRow, nodeCol model.Dimension,
	value model.Value,
) rule {
	r := rule{
		affects: int(value),
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.getExpensiveNodeCheckerRule(value)
	return r
}

// TODO figure this one out!!!
func (r *rule) getExpensiveNodeCheckerRule(
	v model.Value,
) func(*state) {
	// TODO is it better to scope these vars once up here?
	// var right, down, left, up uint32
	// var cr, cd, cl, cu bool
	return func(s *state) {
		cr, cd, cl, cu := true, true, true, true
		pd, nd := model.Dimension(0), model.Dimension(1)
		var total model.Value

		for {
			// check right
			if cr {
				if s.horLineAt(r.row, r.col+pd) {
					total++
					if total > v {
						r.setInvalid(s)
						return
					}
				} else {
					cr = false
				}
			}
			// check left
			if cl {
				if nd < r.col && s.horLineAt(r.row, r.col-nd) {
					total++
					if total > v {
						r.setInvalid(s)
						return
					}
				} else {
					cl = false
				}
			}
			// check down
			if cd {
				if s.verLineAt(r.row+pd, r.col) {
					total++
					if total > v {
						r.setInvalid(s)
						return
					}
				} else {
					cd = false
				}
			}
			// check up
			if cu {
				if nd < r.row && s.verLineAt(r.row-nd, r.col) {
					total++
					if total > v {
						r.setInvalid(s)
						return
					}
				} else {
					cu = false
				}
			}

			if !cr && !cd && !cl && !cu {
				break
			}
			pd++
			nd++
		}
	}
}
