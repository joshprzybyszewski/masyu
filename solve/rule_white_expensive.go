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

// TODO figure this one out!!!
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
		posBit := uint32(1 << 0)
		negBit := uint32(1 << (v - 2))
		pd, nd := model.Dimension(0), model.Dimension(1)

		for {
			// check right
			if cr {
				if s.horAvoidAt(r.row, r.col+pd) {
					cr = false
				} else {
					// TODO if there is a white node at [r.row, r.col+pd+1],
					// then I CANNOT place a bit here.
					right |= posBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.verLineAt(r.row, r.col+pd+1) ||
						s.verLineAt(r.row-1, r.col+pd+1) {
						// cannot continue
						cr = false
					}
					// TODO if there is a black node at [r.row, r.col+pd+1],
					// then I CANNOT continue
				}
			}
			// check left
			if cl {
				if nd >= r.col || s.horAvoidAt(r.row, r.col-nd) {
					cl = false
				} else {
					// TODO if there is a white node at [r.row, r.col-nd],
					// then I CANNOT place a bit here.
					left |= negBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.verLineAt(r.row, r.col-nd) ||
						s.verLineAt(r.row-1, r.col-nd) {
						// cannot continue
						cl = false
					}
					// TODO if there is a black node at [r.row, r.col-nd],
					// then I CANNOT continue
				}
			}
			// check down
			if cd {
				if s.verAvoidAt(r.row+pd, r.col) {
					cd = false
				} else {
					// TODO if there is a white node at [r.row+pd+1, r.col],
					// then I CANNOT place a bit here.
					down |= posBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.horLineAt(r.row+pd+1, r.col) ||
						s.horLineAt(r.row+pd+1, r.col-1) {
						// cannot continue
						cd = false
					}
					// TODO if there is a black node at [r.row, r.col+pd+1],
					// then I CANNOT continue
				}
			}
			// check up
			if cu {
				if nd >= r.row || s.verAvoidAt(r.row-nd, r.col) {
					cu = false
				} else {
					// TODO if there is a white node at [r.row-nd, r.col],
					// then I CANNOT place a bit here.
					up |= negBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.horLineAt(r.row-nd, r.col) ||
						s.horLineAt(r.row-nd, r.col-1) {
						// cannot continue
						cu = false
					}
					// TODO if there is a black node at [r.row-nd, r.col],
					// then I CANNOT continue
				}
			}

			if !cr && !cd && !cl && !cu {
				break
			}
			posBit <<= 1
			negBit >>= 1
			if negBit == 0 {
				break
			}
			pd++
			nd++
		}

		// check right and left
		// if right&left == 0 { // TODO I think this case can be expanded
		if right == 0 || left == 0 { // TODO I think this case can be expanded
			// cannot be place right and left. must go down and up.
			du := down & up
			if du == 0 {
				// cannot go left and up. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid left and right
			s.avoidHor(r.row, r.col-1)
			s.avoidHor(r.row, r.col)

			numPos := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&du == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking du & !bit will be zero.
					// If that's the case, then we know how many to set.
					du &= (^bit)
					if du == 0 {
						numPos = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numPos == 0 {
				// we don't know how many in each direction it'll be.
				// just set up and down.
				s.lineVer(r.row-1, r.col)
				s.lineVer(r.row, r.col)
				return
			}

			// set a line out positive vertically, then avoid the one after it.
			for pd = model.Dimension(0); pd < numPos; pd++ {
				s.lineVer(r.row+pd, r.col)
			}
			s.avoidVer(r.row+pd, r.col)

			// set a line out negative vertically, then avoid the one after it.
			numNeg := model.Dimension(v) - numPos
			for nd = model.Dimension(1); nd <= numNeg; nd++ {
				s.lineVer(r.row-nd, r.col)
			}
			s.avoidVer(r.row-nd, r.col)
			return
		}

		// check up and down
		// if up&down == 0 { // TODO I think this case can be expanded
		if up == 0 || down == 0 { // TODO I think this case can be expanded
			// cannot be placed up and down. must go left and right.
			lr := left & right
			if lr == 0 {
				// cannot go up and right. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid up and down
			s.avoidVer(r.row-1, r.col)
			s.avoidVer(r.row, r.col)

			numPos := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&lr == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking lr & !bit will be zero.
					// If that's the case, then we know how many to set.
					lr &= (^bit)
					if lr == 0 {
						numPos = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numPos == 0 {
				// we don't know how many in each direction it'll be.
				// just set left and right.
				s.lineHor(r.row, r.col-1)
				s.lineHor(r.row, r.col)
				return
			}

			// set a line out positive horizontally, then avoid the one after it.
			for pd = model.Dimension(0); pd < numPos; pd++ {
				s.lineHor(r.row, r.col+pd)
			}
			s.avoidHor(r.row, r.col+pd)

			// set a line out negative horizontally, then avoid the one after it.
			numNeg := model.Dimension(v) - numPos
			for nd = model.Dimension(1); nd <= numNeg; nd++ {
				s.lineHor(r.row, r.col-nd)
			}
			s.avoidHor(r.row, r.col-nd)
			return
		}
	}
}
