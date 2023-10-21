package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackExpensiveRule(
	nodeRow, nodeCol model.Dimension,
	value model.Value,
) rule {
	r := rule{
		affects: int(value) + 4,
		row:     nodeRow,
		col:     nodeCol,
	}
	r.check = r.getExpensiveBlackRule(value)
	return r
}

func (r *rule) getExpensiveBlackRule(
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
		horBit := uint32(1 << 0)
		verBit := uint32(1 << (v - 1))
		pd, nd := model.Dimension(0), model.Dimension(1)

		for {
			// check right
			if cr {
				if s.horAvoidAt(r.row, r.col+pd) {
					cr = false
				} else {
					// TODO if there is a white node at [r.row, r.col+pd+1],
					// then I CANNOT place a bit here.
					right |= horBit

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
					left |= horBit

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
					down |= verBit

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
				if nd >= r.col || s.verAvoidAt(r.row-nd, r.col) {
					cu = false
				} else {
					// TODO if there is a white node at [r.row-nd, r.col],
					// then I CANNOT place a bit here.
					up |= verBit

					// There is a spur coming into my path. I should not continue checking this direction
					if s.horLineAt(r.row-nd, r.col) ||
						s.horLineAt(r.row-nd, r.col+1) {
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
			horBit <<= 1
			verBit >>= 1
			pd++
			nd++
		}

		// check right and down
		if right == 0 && down == 0 { // TODO I think this case can be expanded
			// cannot be place right and down. must go left and up.
			lu := left & up
			if lu == 0 {
				// cannot go left and up. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid right and down
			s.avoidHor(r.row, r.col)
			s.avoidVer(r.row, r.col)

			numHor := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&lu == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking lu & !bit will be zero.
					// If that's the case, then we know how many to set.
					lu &= (^bit)
					if lu == 0 {
						numHor = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numHor == 0 {
				// we don't know how many in each direction it'll be.
				// just set left and up.
				s.lineHor(r.row, r.col-1)
				s.lineVer(r.row-1, r.col)
				return
			}

			// set a line out horizontally, then avoid the one after it.
			for nd = model.Dimension(1); nd <= numHor; nd++ {
				s.lineHor(r.row, r.col-nd)
			}
			s.avoidHor(r.row, r.col-nd)

			// set a line out vertically, then avoid the one after it.
			numVer := model.Dimension(v) - numHor
			for nd = model.Dimension(1); nd <= numVer; nd++ {
				s.lineVer(r.row-nd, r.col)
			}
			s.avoidVer(r.row-nd, r.col)
			return
		}

		// check down and left
		if down == 0 && left == 0 { // TODO I think this case can be expanded
			// cannot be placed down and left. must go up and right.
			ur := up & right
			if ur == 0 {
				// cannot go up and right. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid down and left
			s.avoidVer(r.row, r.col)
			s.avoidHor(r.row, r.col-1)

			numHor := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&ur == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking ur & !bit will be zero.
					// If that's the case, then we know how many to set.
					ur &= (^bit)
					if ur == 0 {
						numHor = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numHor == 0 {
				// we don't know how many in each direction it'll be.
				// just set up and right.
				s.lineVer(r.row-1, r.col)
				s.lineHor(r.row, r.col)
				return
			}

			// set a line out horizontally, then avoid the one after it.
			for pd = model.Dimension(0); pd < numHor; pd++ {
				s.lineHor(r.row, r.col+pd)
			}
			s.avoidHor(r.row, r.col+pd)

			// set a line out vertically, then avoid the one after it.
			numVer := model.Dimension(v) - numHor
			for nd = model.Dimension(1); nd <= numVer; nd++ {
				s.lineVer(r.row-nd, r.col)
			}
			s.avoidVer(r.row-nd, r.col)
			return
		}

		// check left and up
		if left == 0 && up == 0 { // TODO I think this case can be expanded
			// cannot be placed left and up. must go right and down.
			rd := right & down
			if rd == 0 {
				// cannot go right and down. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid left and up
			s.avoidHor(r.row, r.col-1)
			s.avoidVer(r.row-1, r.col)

			numHor := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&rd == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking rd & !bit will be zero.
					// If that's the case, then we know how many to set.
					rd &= (^bit)
					if rd == 0 {
						numHor = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numHor == 0 {
				// we don't know how many in each direction it'll be.
				// just set right and down.
				s.lineHor(r.row, r.col)
				s.lineVer(r.row, r.col)
				return
			}

			// set a line out horizontally, then avoid the one after it.
			for pd = model.Dimension(0); pd < numHor; pd++ {
				s.lineHor(r.row, r.col+pd)
			}
			s.avoidHor(r.row, r.col+pd)

			// set a line out vertically, then avoid the one after it.
			numVer := model.Dimension(v) - numHor
			for pd = model.Dimension(0); pd < numVer; pd++ {
				s.lineVer(r.row+pd, r.col)
			}
			s.avoidVer(r.row+pd, r.col)
			return
		}

		// check up and right
		if up == 0 && right == 0 { // TODO I think this case can be expanded
			// cannot be placed up and right. must go down and left.
			dl := down & left
			if dl == 0 {
				// cannot go right and down. invalid!
				r.setInvalid(s)
				return
			}

			// make sure we avoid up and right
			s.avoidVer(r.row-1, r.col)
			s.avoidHor(r.row, r.col)

			numHor := model.Dimension(0)
			for n, bit := model.Value(1), uint32(1); ; {
				if bit&dl == bit {
					// we found one option for the number of horizontal placements.
					// If this is the only option, then taking dl & !bit will be zero.
					// If that's the case, then we know how many to set.
					dl &= (^bit)
					if dl == 0 {
						numHor = model.Dimension(n)
					}
					break
				}
				bit <<= 1
				n++
			}
			if numHor == 0 {
				// we don't know how many in each direction it'll be.
				// just set down and left.
				s.lineVer(r.row, r.col)
				s.lineHor(r.row, r.col-1)
				return
			}

			// set a line out horizontally, then avoid the one after it.
			for nd = model.Dimension(1); nd <= numHor; nd++ {
				s.lineHor(r.row, r.col-nd)
			}
			s.avoidHor(r.row, r.col-nd)

			// set a line out vertically, then avoid the one after it.
			numVer := model.Dimension(v) - numHor
			for pd = model.Dimension(0); pd < numVer; pd++ {
				s.lineVer(r.row+pd, r.col)
			}
			s.avoidVer(r.row+pd, r.col)
			return
		}
	}
}
