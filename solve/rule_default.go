package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func newDefaultRule(
	row, col model.Dimension,
) rule {
	r := rule{
		row: row,
		col: col,
	}
	r.check = r.checkDefault
	return r
}

func (r *rule) checkDefault(
	s *state,
) {

	l, a := s.horAt(r.row, r.col)
	if l {
		// RIGHT
		l, a = s.verAt(r.row, r.col)
		if l {
			// RIGHT + DOWN
			s.avoidHor(r.row, r.col-1)
			s.avoidVer(r.row-1, r.col)
		} else if a {
			// RIGHT + !DOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// RIGHT + !DOWN + LEFT
				s.avoidVer(r.row-1, r.col)
			} else if a {
				// RIGHT + !DOWN + !LEFT
				s.lineVer(r.row-1, r.col)
			} else {
				// RIGHT + !DOWN + uLEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// RIGHT + !DOWN + UP
					s.avoidHor(r.row, r.col-1)
				} else if a {
					// RIGHT + !DOWN + !UP
					s.lineHor(r.row, r.col-1)
				}
			}
		} else {
			// RIGHT + uDOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// RIGHT + LEFT
				s.avoidVer(r.row-1, r.col)
				s.avoidVer(r.row, r.col)
			} else if a {
				// RIGHT + uDOWN + !LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// RIGHT + uDOWN + !LEFT + UP
					s.avoidVer(r.row, r.col)
				} else if a {
					// RIGHT + uDOWN + !LEFT + !UP
					s.lineVer(r.row, r.col)
				}
			} else {
				// RIGHT + uDOWN + uLEFT
				if s.verLineAt(r.row-1, r.col) {
					// RIGHT + uDOWN + uLEFT + UP
					s.avoidVer(r.row, r.col)
					s.avoidHor(r.row, r.col-1)
				}
			}
		}
	} else if a {
		// !RIGHT
		l, a = s.verAt(r.row, r.col)
		if l {
			// !RIGHT + DOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// !RIGHT + DOWN + LEFT
				s.avoidVer(r.row-1, r.col)
			} else if a {
				// !RIGHT + DOWN + !LEFT
				s.lineVer(r.row-1, r.col)
			} else {
				// !RIGHT + DOWN + uLEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// !RIGHT + DOWN + UP
					s.avoidHor(r.row, r.col-1)
				} else if a {
					// !RIGHT + DOWN + !UP
					s.lineHor(r.row, r.col-1)
				}
			}
		} else if a {
			// !RIGHT + !DOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// !RIGHT + !DOWN + LEFT
				s.lineVer(r.row-1, r.col)
			} else if a {
				// !RIGHT + !DOWN + !LEFT
				s.avoidVer(r.row-1, r.col)
			} else {
				// !RIGHT + !DOWN + uLEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// !RIGHT + !DOWN + UP
					s.lineHor(r.row, r.col-1)
				} else if a {
					// !RIGHT + !DOWN + !UP
					s.avoidHor(r.row, r.col-1)
				}
			}
		} else {
			// !RIGHT + uDOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// !RIGHT + uDOWN + LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// !RIGHT + uDOWN + LEFT + UP
					s.avoidVer(r.row, r.col)
				} else if a {
					// !RIGHT + uDOWN + LEFT + !UP
					s.lineVer(r.row, r.col)
				}
			} else if a {
				// !RIGHT + uDOWN + !LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// !RIGHT + uDOWN + !LEFT + UP
					s.lineVer(r.row, r.col)
				} else if a {
					// !RIGHT + uDOWN + !LEFT + !UP
					s.avoidVer(r.row, r.col)
				}
			}
		}
	} else {
		// uRIGHT
		l, a = s.verAt(r.row, r.col)
		if l {
			// uRIGHT + DOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// uRIGHT + DOWN + LEFT
				s.avoidHor(r.row, r.col)
				s.avoidVer(r.row-1, r.col)
			} else if a {
				// uRIGHT + DOWN + !LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// uRIGHT + DOWN + !LEFT + UP
					s.avoidHor(r.row, r.col)
				} else if a {
					// uRIGHT + DOWN + !LEFT + !UP
					s.lineHor(r.row, r.col)
				} // else uRIGHT + DOWN + !LEFT + uUP
			} else {
				// else uRIGHT + DOWN + uLEFT
				if s.verLineAt(r.row-1, r.col) {
					// else uRIGHT + DOWN + uLEFT + UP
					s.avoidHor(r.row, r.col)
					s.avoidHor(r.row, r.col-1)
				}
			}
		} else if a {
			// uRIGHT + !DOWN
			l, a = s.horAt(r.row, r.col-1)
			if l {
				// uRIGHT + !DOWN + LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// uRIGHT + !DOWN + LEFT + UP
					s.avoidHor(r.row, r.col)
				} else if a {
					// uRIGHT + !DOWN + LEFT + !UP
					s.lineHor(r.row, r.col)
				}
			} else if a {
				// uRIGHT + !DOWN + !LEFT
				l, a = s.verAt(r.row-1, r.col)
				if l {
					// uRIGHT + !DOWN + !LEFT + UP
					s.lineHor(r.row, r.col)
				} else if a {
					// uRIGHT + !DOWN + !LEFT + !UP
					s.avoidHor(r.row, r.col)
				}
			} // else uRIGHT + !DOWN + uLEFT
		} else {
			// uRIGHT + uDOWN
			if s.horLineAt(r.row, r.col-1) &&
				s.verLineAt(r.row-1, r.col) {
				// uRIGHT + uDOWN + LEFT + UP
				s.avoidHor(r.row, r.col)
				s.avoidVer(r.row, r.col)
			}
		}
	}
}
