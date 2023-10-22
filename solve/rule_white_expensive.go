package solve

import "github.com/joshprzybyszewski/masyu/model"

func newWhiteExpensiveRule(
	node model.Node,
	nodes *[maxPinsPerLine][maxPinsPerLine]model.Node,
) rule {
	r := rule{
		affects: int(node.Value) + 4,
		row:     node.Row,
		col:     node.Col,
	}
	bounds := getWhiteBounds(node, nodes)
	r.check = r.getExpensiveWhiteRule(node.Value, bounds)
	return r
}

func getWhiteBounds(
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
	if node.Value <= 2 {
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
			if node.Value >= otherVal {
				// our straight line is too much for this black node. stop before we get there.
				b.maxRight = c - 2
			} else {
				b.maxRight = c - 1
			}
			break
		}

		// we found a white node at this pin.
		if node.Value != otherVal {
			// If we cannot continue on through this white node together
			b.maxRight = c - 2
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
			if node.Value >= otherVal {
				b.maxLeft = c + 1
			} else {
				b.maxLeft = c
			}
			break
		}

		// we found a white node at this pin.
		if node.Value != otherVal {
			b.maxLeft = c + 1
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
			if node.Value >= otherVal {
				b.maxDown = r - 2
			} else {
				b.maxDown = r - 1
			}
			break
		}

		// we found a white node at this pin.
		if node.Value != otherVal {
			b.maxDown = r - 2
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
			if node.Value >= otherVal {
				b.maxUp = r + 1
			} else {
				b.maxUp = r
			}
			break
		}

		// we found a white node at this pin.
		if node.Value != otherVal {
			b.maxUp = r + 1
		}
		break
	}

	return b
}

func (r *rule) getExpensiveWhiteRule(
	v model.Value,
	bounds bounds,
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
		var goal uint32
		cr, cd, cl, cu := true, true, true, true
		posBit := uint32(1)
		negBit := uint32(1 << (v - 2))
		pd, nd := model.Dimension(0), model.Dimension(1)

		for {
			// check right
			if cr {
				if r.col+pd > bounds.maxRight || s.horAvoidAt(r.row, r.col+pd) {
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
				if nd >= r.col || r.col-nd < bounds.maxLeft || s.horAvoidAt(r.row, r.col-nd) {
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
				if r.row+pd > bounds.maxDown || s.verAvoidAt(r.row+pd, r.col) {
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
				if nd >= r.row || r.row-nd < bounds.maxUp || s.verAvoidAt(r.row-nd, r.col) {
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
			goal = down & up
			if goal == 0 {
				// cannot go up and down? invalid!
				r.setInvalid(s)
				return
			}

			negBit = uint32(1 << (v - 2))
			for nd = 1; ; {
				s.lineVer(r.row-nd, r.col)
				nd++

				if negBit&goal != negBit {
					// we can't stop here; must continue
					negBit >>= 1
					continue
				}

				if goal&(^negBit) == 0 {
					s.avoidVer(r.row-nd, r.col)
				}
				break
			}

			posBit = uint32(1)
			for pd = 0; ; {
				s.lineVer(r.row+pd, r.col)
				pd++

				if posBit&goal == 0 {
					// we can't stop here; must continue
					posBit <<= 1
					continue
				}

				if goal&(^posBit) == 0 {
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

		goal = right & left
		negBit = uint32(1 << (v - 2))
		for nd = 1; ; {
			s.lineHor(r.row, r.col-nd)
			nd++

			if negBit&goal != negBit {
				// we can't stop here; must continue
				negBit >>= 1
				continue
			}

			if goal&(^negBit) == 0 {
				s.avoidHor(r.row, r.col-nd)
			}
			break
		}

		posBit = uint32(1)
		for pd = 0; ; {
			s.lineHor(r.row, r.col+pd)
			pd++

			if posBit&goal == 0 {
				// we can't stop here; must continue
				posBit <<= 1
				continue
			}

			if goal&(^posBit) == 0 {
				s.avoidHor(r.row, r.col+pd)
			}
			break
		}
	}
}
