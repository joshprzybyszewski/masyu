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

func (r *rule) getExpensiveNodeCheckerRule(
	v model.Value,
) func(*state) {
	return func(s *state) {
		var pd, nd model.Dimension
		var total model.Value

		for pd = 0; s.horLineAt(r.row, r.col+pd); pd++ {
			total++
			if total > v {
				r.setInvalid(s)
				return
			}
		}
		for nd = 1; nd < r.col && s.horLineAt(r.row, r.col-nd); nd++ {
			total++
			if total > v {
				r.setInvalid(s)
				return
			}
		}
		for pd = 0; s.verLineAt(r.row+pd, r.col); pd++ {
			total++
			if total > v {
				r.setInvalid(s)
				return
			}
		}
		for nd = 1; nd < r.row && s.verLineAt(r.row-nd, r.col); nd++ {
			total++
			if total > v {
				r.setInvalid(s)
				return
			}
		}
	}
}
