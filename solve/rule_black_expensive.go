package solve

import "github.com/joshprzybyszewski/masyu/model"

func newBlackExpensiveRule(
	node model.Node,
	nodes *[maxPinsPerLine][maxPinsPerLine]model.Node,
) rule {
	r := rule{
		affects: int(node.Value) + 4,
		row:     node.Row,
		col:     node.Col,
	}
	bounds := getBlackBounds(node, nodes)
	r.check = r.getExpensiveBlackRule(node.Value, bounds)
	return r
}

type bounds struct {
	maxRight model.Dimension
	maxDown  model.Dimension
	maxLeft  model.Dimension
	maxUp    model.Dimension
}

func getBlackBounds(
	node model.Node,
	nodes *[maxPinsPerLine][maxPinsPerLine]model.Node,
) bounds {
	vm1 := model.Dimension(node.Value - 1)
	b := bounds{
		maxRight: node.Col + vm1 - 1,
		maxDown:  node.Row + vm1 - 1,
	}
	if vm1 < node.Col {
		b.maxLeft = node.Col - vm1
	}
	if vm1 < node.Row {
		b.maxUp = node.Row - vm1
	}
	// TODO change this if-condition to be <= 2 once we get the rest of the logic in this function correct
	if node.Value <= 2 {
		// if node.Value > 1 {
		return b
	}

	var otherVal model.Value

	// check the right
	for c := node.Col + 1; c <= b.maxRight; c++ {
		if nodes[node.Row][c].Value == 0 {
			continue
		}
		otherVal = nodes[node.Row][c].Value

		if nodes[node.Row][c].IsBlack {
			// if we found a black node at this pin, then we either need to stop here,
			// or at the column before here.
			if c-node.Col > model.Dimension(otherVal-1) {
				b.maxRight = c - 2
			} else {
				b.maxRight = c - 1
			}
			break
		}

		// we found a white node at this pin.
		if c-node.Col > model.Dimension(otherVal-1) ||
			b.maxRight < node.Col+model.Dimension(otherVal-1) {
			// the distance to get to this node is more than the white node can handle OR
			// we cannot go far enough to satisfy this white node.
			b.maxRight = c - 2
		} else {
			// if we're going to go through this white node, then we cannot go beyond the limitation
			// that it puts on us.
			b.maxRight = node.Col + model.Dimension(otherVal-1)
		}
		break
	}

	// check the left
	for c := node.Col - 1; c > 0 && c >= b.maxLeft; c-- {
		if nodes[node.Row][c].Value == 0 {
			continue
		}
		otherVal = nodes[node.Row][c].Value

		if nodes[node.Row][c].IsBlack {
			// if we found a black node at this pin, then we either need to stop here,
			// or at the column before here.
			if node.Col-c > model.Dimension(otherVal-1) {
				b.maxLeft = c + 1
			} else {
				b.maxLeft = c
			}
			break
		}

		// we found a white node at this pin.
		if node.Col-c > model.Dimension(otherVal-1) ||
			model.Dimension(otherVal) > node.Col ||
			b.maxLeft > node.Col-model.Dimension(otherVal) {
			// the distance to get to this node is more than the white node can handle OR
			// we cannot go far enough to satisfy this white node.
			b.maxLeft = c + 1
		} else {
			// if we're going to go through this white node, then we cannot go beyond the limitation
			// that it puts on us.
			b.maxLeft = node.Col - model.Dimension(otherVal)
		}
		break
	}

	// check down
	for r := node.Row + 1; r <= b.maxDown; r++ {
		if nodes[r][node.Col].Value == 0 {
			continue
		}
		otherVal = nodes[r][node.Col].Value

		if nodes[r][node.Col].IsBlack {
			// if we found a black node at this pin, then we either need to stop here,
			// or at the column before here.
			if r-node.Row > model.Dimension(otherVal-1) {
				b.maxDown = r - 2
			} else {
				b.maxDown = r - 1
			}
			break
		}

		// we found a white node at this pin.
		if r-node.Row > model.Dimension(otherVal-1) ||
			b.maxDown < node.Row+model.Dimension(otherVal-1) {
			// the distance to get to this node is more than the white node can handle OR
			// we cannot go far enough to satisfy this white node.
			b.maxDown = r - 2
		} else {
			// if we're going to go through this white node, then we cannot go beyond the limitation
			// that it puts on us.
			b.maxDown = node.Row + model.Dimension(otherVal-1)
		}
		break
	}

	// check up
	for r := node.Row - 1; r > 0 && r >= b.maxUp; r-- {
		if nodes[r][node.Col].Value == 0 {
			continue
		}
		otherVal = nodes[r][node.Col].Value

		if nodes[r][node.Col].IsBlack {
			// if we found a black node at this pin, then we either need to stop here,
			// or at the column before here.
			if node.Row-r > model.Dimension(otherVal-1) {
				b.maxUp = r + 1
			} else {
				b.maxUp = r
			}
			break
		}

		// we found a white node at this pin.
		if node.Row-r > model.Dimension(otherVal-1) ||
			model.Dimension(otherVal) > node.Row ||
			b.maxUp > node.Row-model.Dimension(otherVal) {
			// the distance to get to this node is more than the white node can handle OR
			// we cannot go far enough to satisfy this white node.
			b.maxUp = r + 1
		} else {
			// if we're going to go through this white node, then we cannot go beyond the limitation
			// that it puts on us.
			b.maxUp = node.Row - model.Dimension(otherVal)
		}
		break
	}

	return b
}

func (r *rule) getExpensiveBlackRule(
	v model.Value,
	bounds bounds,
) func(*state) {
	if v > 32 {
		// I use 32 bits to keep track of how far it can go. If the puzzle is one
		// of the monthly specials (41 pins) or weekly specials (36 pins), then I'm
		// just gonna say "nope".
		// TODO a black node on a corner can hit > 32 on 25x25. In that case, I just need to send out
		// arms the appropriate direction.
		panic(`unsupported puzzle`)
	}
	// TODO is it better to scope these vars once up here?
	// var right, down, left, up uint32
	// var cr, cd, cl, cu bool
	return func(s *state) {
		var right, down, left, up uint32 // := 0, 0, 0, 0
		var goal uint32
		cr, cd, cl, cu := true, true, true, true
		horBit := uint32(1 << 0)
		verBit := uint32(1 << (v - 2))
		pd, nd := model.Dimension(0), model.Dimension(1)

		for {
			// check right
			if cr {
				if r.col+pd > bounds.maxRight || s.horAvoidAt(r.row, r.col+pd) {
					cr = false
				} else {
					if !s.horLineAt(r.row, r.col+pd+1) {
						// if there's a line at the next spot, then I can't end here.
						right |= horBit
					}

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
				// TODO i think this first condition might be unnecessary?
				if nd >= r.col || r.col-nd < bounds.maxLeft || s.horAvoidAt(r.row, r.col-nd) {
					cl = false
				} else {
					if nd+1 < r.col || !s.horLineAt(r.row, r.col-nd-1) {
						// if there's a line at the next spot, then I can't end here.
						left |= horBit
					}

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
				if r.row+pd > bounds.maxDown || s.verAvoidAt(r.row+pd, r.col) {
					cd = false
				} else {
					if !s.verLineAt(r.row+pd+1, r.col) {
						// if there's a line at the next spot, then I can't end here.
						down |= verBit
					}

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
				// TODO i think this first condition might be unnecessary?
				if nd >= r.row || r.row-nd < bounds.maxUp || s.verAvoidAt(r.row-nd, r.col) {
					cu = false
				} else {
					if nd+1 < r.row || !s.verLineAt(r.row-nd-1, r.col) {
						// if there's a line at the next spot, then I can't end here.
						up |= verBit
					}

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
			if pd >= model.Dimension(v) {
				break
			}
			horBit <<= 1
			verBit >>= 1
			pd++
			nd++
		}

		if right&(up|down) == 0 {
			// must go left
			s.avoidHor(r.row, r.col)
			bit := uint32(1)
			goal = left & (up | down)
			for nd = 1; ; {
				s.lineHor(r.row, r.col-nd)
				nd++

				if bit&goal != bit {
					// we can't stop here; must continue
					bit <<= 1
					continue
				}

				if goal&(^bit) == 0 {
					s.avoidHor(r.row, r.col-nd)
				}
				break
			}
		} else if left&(up|down) == 0 {
			// must go right
			s.avoidHor(r.row, r.col-1)
			bit := uint32(1)
			goal = right & (up | down)
			for pd = 0; ; {
				s.lineHor(r.row, r.col+pd)
				pd++

				if bit&goal == 0 {
					// we can't stop here; must continue
					bit <<= 1
					continue
				}

				if goal&(^bit) == 0 {
					s.avoidHor(r.row, r.col+pd)
				}
				break
			}
		}

		if down&(right|left) == 0 {
			// must go up
			s.avoidVer(r.row, r.col)
			bit := uint32(1 << (v - 2))
			goal = up & (right | left)
			for nd = 1; ; {
				s.lineVer(r.row-nd, r.col)
				nd++

				if bit&goal != bit {
					// we can't stop here; must continue
					bit >>= 1
					continue
				}

				if goal&(^bit) == 0 {
					s.avoidVer(r.row-nd, r.col)
				}
				break
			}
		} else if up&(right|left) == 0 {
			// must go down
			s.avoidVer(r.row-1, r.col)
			bit := uint32(1 << (v - 2))
			goal = down & (right | left)
			for pd = 0; ; {
				s.lineVer(r.row+pd, r.col)
				pd++

				if bit&goal == 0 {
					// we can't stop here; must continue
					bit >>= 1
					continue
				}

				if goal&(^bit) == 0 {
					s.avoidVer(r.row+pd, r.col)
				}
				break
			}
		}

		/*
			// check right and down
			// if right&down == 0 { // TODO I think this case can be expanded
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

				// TODO send out a minimum horizontal lines
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
			// if down&left == 0 { // TODO I think this case can be expanded
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
			// if left&up == 0 { // TODO I think this case can be expanded
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
			// if up&right == 0 { // TODO I think this case can be expanded
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
		*/
	}
}
