package solve

import "github.com/joshprzybyszewski/masyu/model"

func newWhiteExpensiveRule(
	nodeRow, nodeCol model.Dimension,
	value model.Value,
) rule {
	r := rule{
		affects: int(value) + 4,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.getExpensiveWhiteRule(value)
	return r
}

func (r *rule) getExpensiveWhiteRule(
	v model.Value,
) func(*state) {
	if v > 32 {
		// I use 32 bits to keep track of how far it can go. If the puzzle is one
		// of the monthly specials (41 pins) or weekly specials (36 pins), then I'm
		// just gonna say "nope".
		panic(`unsupported puzzle`)
	}
	// TODO is it better to scope these vars once up here?
	// var right, down, left, up uint32
	// var cr, cd, cl, cu bool
	return func(s *state) {
		var right, down, left, up uint32 // := 0, 0, 0, 0
		cr, cd, cl, cu := true, true, true, true
		posBit := uint32(1)
		negBit := uint32(1 << (v - 2))
		pd, nd := model.Dimension(0), model.Dimension(1)

		// TODO copy the logic from the black node eval.
		for {
			// check right
			if cr {
				if s.horAvoidAt(r.row, r.col+pd) {
					cr = false
				} else {
					right |= posBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.verLineAt(r.row, r.col+pd+1) ||
						s.verLineAt(r.row-1, r.col+pd+1) {
						// cannot continue
						cr = false
					}
				}
			}
			// check left
			if cl {
				if nd >= r.col || s.horAvoidAt(r.row, r.col-nd) {
					cl = false
				} else {
					left |= negBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.verLineAt(r.row, r.col-nd) ||
						s.verLineAt(r.row-1, r.col-nd) {
						// cannot continue
						cl = false
					}
				}
			}
			// check down
			if cd {
				if s.verAvoidAt(r.row+pd, r.col) {
					cd = false
				} else {
					down |= posBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.horLineAt(r.row+pd+1, r.col) ||
						s.horLineAt(r.row+pd+1, r.col-1) {
						// cannot continue
						cd = false
					}
				}
			}
			// check up
			if cu {
				if nd >= r.row || s.verAvoidAt(r.row-nd, r.col) {
					cu = false
				} else {
					up |= negBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.horLineAt(r.row-nd, r.col) ||
						s.horLineAt(r.row-nd, r.col-1) {
						// cannot continue
						cu = false
					}
				}
			}

			if !cr && !cd && !cl && !cu {
				break
			}
			posBit <<= 1
			negBit >>= 1
			if negBit == 0 {
				// hit the max size for lines.
				break
			}
			pd++
			nd++
		}

		if right&left == 0 {
			// cannot go right and left; must go up and down
			if down&up == 0 {
				// cannot go up and down? invalid!
				r.setInvalid(s)
				return
			}

			negBit = uint32(1 << (v - 2))
			for nd = 1; ; {
				s.lineVer(r.row-nd, r.col)
				nd++

				if negBit&up != negBit {
					// we can't stop here; must continue
					negBit <<= 1
					continue
				}

				if up&(^negBit) == 0 {
					s.avoidVer(r.row-nd, r.col)
				}
				break
			}

			posBit = uint32(1)
			for pd = 0; ; {
				s.lineVer(r.row+pd, r.col)
				pd++

				if posBit&down == 0 {
					// we can't stop here; must continue
					posBit <<= 1
					continue
				}

				if down&(^posBit) == 0 {
					s.avoidVer(r.row+pd, r.col)
				}
				break
			}
			return
		}

		if up&down != 0 {
			// it _could_ go up and down; can't infer anything to continue.
			return
		}

		negBit = uint32(1 << (v - 2))
		for nd = 1; ; {
			s.lineHor(r.row, r.col-nd)
			nd++

			if negBit&left != negBit {
				// we can't stop here; must continue
				negBit <<= 1
				continue
			}

			if left&(^negBit) == 0 {
				s.avoidHor(r.row, r.col-nd)
			}
			break
		}

		posBit = uint32(1)
		for pd = 0; ; {
			s.lineHor(r.row, r.col+pd)
			pd++

			if posBit&right == 0 {
				// we can't stop here; must continue
				posBit <<= 1
				continue
			}

			if right&(^posBit) == 0 {
				s.avoidHor(r.row, r.col+pd)
			}
			break
		}
	}
}
