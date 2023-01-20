package solve

import (
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

type state struct {
	nodes []model.Node

	size model.Size

	horizontalLines  [model.MaxPointsPerLine]uint64
	horizontalAvoids [model.MaxPointsPerLine]uint64

	verticalLines  [model.MaxPointsPerLine]uint64
	verticalAvoids [model.MaxPointsPerLine]uint64
}

func newState(
	size model.Size,
	ns []model.Node,
) state {
	s := state{
		nodes: make([]model.Node, len(ns)),
		size:  size,
	}

	for i := range ns {
		s.nodes[i] = ns[i]
		s.nodes[i].Row++
		s.nodes[i].Col++
	}

	s.initialize()

	return s
}

func (s *state) toSolution() (model.Solution, bool, bool) {
	if !s.isValid() {
		return model.Solution{}, false, false
	}

	if !s.isValidPath() {
		return model.Solution{}, false, true
	}

	sol := model.Solution{
		Size: s.size,
	}

	for i := 0; i < int(s.size); i++ {
		sol.Horizontals[i] = (s.horizontalLines[i+1]) >> 1
		sol.Verticals[i] = (s.verticalLines[i+1]) >> 1
	}

	return sol, true, true
}

func (s *state) isValidPath() bool {
	var horizontalLines [model.MaxPointsPerLine]uint64
	var verticalLines [model.MaxPointsPerLine]uint64

	start := s.nodes[0].Coord
	cur := start
	prev := cur
	var l bool
	for {
		if l, _ = s.horAt(cur.Row, cur.Col); l && prev.Col != cur.Col+1 {
			prev = cur
			horizontalLines[cur.Row] |= (1 << cur.Col)
			cur.Col++
		} else if l, _ = s.horAt(cur.Row, cur.Col-1); l && prev.Col != cur.Col-1 {
			prev = cur
			horizontalLines[cur.Row] |= (1 << (cur.Col - 1))
			cur.Col--
		} else if l, _ = s.verAt(cur.Row, cur.Col); l && prev.Row != cur.Row+1 {
			prev = cur
			verticalLines[cur.Col] |= (1 << cur.Row)
			cur.Row++
		} else if l, _ = s.verAt(cur.Row-1, cur.Col); l && prev.Row != cur.Row-1 {
			prev = cur
			verticalLines[cur.Col] |= (1 << (cur.Row - 1))
			cur.Row--
		} else {
			return false
		}
		if cur == start {
			break
		}
	}
	if cur == prev || cur != start {
		return false
	}

	return horizontalLines == s.horizontalLines && verticalLines == s.verticalLines
}

func (s *state) isValid() bool {
	for i := 0; i <= int(s.size); i++ {
		if s.horizontalAvoids[i]&s.horizontalLines[i] != 0 {
			return false
		}
		if s.verticalLines[i]&s.verticalAvoids[i] != 0 {
			return false
		}
	}

	return true
}

func (s *state) initialize() {
	avoid := uint64(1 | (1 << s.size))
	for i := 1; i <= int(s.size); i++ {
		s.verticalAvoids[i] |= avoid
		s.horizontalAvoids[i] |= avoid
	}
	s.horizontalAvoids[0] = 0xFFFFFFFFFFFFFFFF
	s.verticalAvoids[0] = 0xFFFFFFFFFFFFFFFF
	s.horizontalAvoids[s.size+1] = 0xFFFFFFFFFFFFFFFF
	s.verticalAvoids[s.size+1] = 0xFFFFFFFFFFFFFFFF

	s.settleNodes()
	if !s.isValid() {
		panic(`dev error`)
	}
}

func (s *state) settleNodes() {
	prevhorizontalLines := s.horizontalLines
	prevhorizontalAvoids := s.horizontalAvoids

	prevverticalLines := s.verticalLines
	prevverticalAvoids := s.verticalAvoids

	i := 0
	for {
		i++

		s.checkNodes()
		s.checkPins()

		if prevhorizontalLines == s.horizontalLines &&
			prevhorizontalAvoids == s.horizontalAvoids &&
			prevverticalLines == s.verticalLines &&
			prevverticalAvoids == s.verticalAvoids {
			break
		}

		prevhorizontalLines = s.horizontalLines
		prevhorizontalAvoids = s.horizontalAvoids
		prevverticalLines = s.verticalLines
		prevverticalAvoids = s.verticalAvoids
	}
}

func (s *state) checkNodes() {
	for _, n := range s.nodes {
		if n.IsBlack {
			s.checkBlack(n.Row, n.Col)
		} else {
			s.checkWhite(n.Row, n.Col)
		}
	}
}

func (s *state) checkBlack(
	r, c int,
) {
	rl, ra := s.horAt(r, c)
	if rl {
		s.lineHor(r, c+1)
		s.avoidHor(r, c-1)
		s.avoidVer(r-1, c+1)
		s.avoidVer(r, c+1)
	} else if ra {
		s.lineHor(r, c-1)
		s.lineHor(r, c-2)
		s.avoidVer(r-1, c-1)
		s.avoidVer(r, c-1)
	}

	dl, da := s.verAt(r, c)
	if dl {
		s.lineVer(r+1, c)
		s.avoidVer(r-1, c)
		s.avoidHor(r+1, c-1)
		s.avoidHor(r+1, c)
	} else if da {
		s.lineVer(r-1, c)
		s.lineVer(r-2, c)
		s.avoidHor(r-1, c)
		s.avoidHor(r-1, c-1)
	}

	ll, la := s.horAt(r, c-1)
	if ll {
		s.lineHor(r, c-2)
		s.avoidHor(r, c)
		s.avoidVer(r-1, c-1)
		s.avoidVer(r, c-1)
	} else if la {
		s.lineHor(r, c)
		s.lineHor(r, c+1)
		s.avoidVer(r-1, c+1)
		s.avoidVer(r, c+1)
	}

	ul, ua := s.verAt(r-1, c)
	if ul {
		s.lineVer(r-2, c)
		s.avoidVer(r, c)
		s.avoidHor(r-1, c-1)
		s.avoidHor(r-1, c)
	} else if ua {
		s.lineVer(r, c)
		s.lineVer(r+1, c)
		s.avoidHor(r+1, c-1)
		s.avoidHor(r+1, c)
	}
}

func (s *state) checkWhite(
	r, c int,
) {
	l1, a1 := s.horAt(r, c-1)
	l2, a2 := s.horAt(r, c)
	hl := l1 || l2
	vl := a1 || a2

	l1, a1 = s.verAt(r-1, c)
	l2, a2 = s.verAt(r, c)
	hl = hl || a1 || a2
	vl = vl || l1 || l2

	if hl {
		s.lineHor(r, c-1)
		s.lineHor(r, c)
		s.avoidVer(r-1, c)
		s.avoidVer(r, c)
		if c < 2 {
			// error state
			s.lineVer(r, c)
			return
		}
		l1, a1 = s.horAt(r, c-2)
		if l1 {
			s.avoidHor(r, c+1)
		} else if !a1 {
			l2, _ = s.horAt(r, c+1)
			if l2 {
				s.avoidHor(r, c-2)
			}
		}
	} else if vl {
		s.avoidHor(r, c-1)
		s.avoidHor(r, c)
		s.lineVer(r-1, c)
		s.lineVer(r, c)
		if r < 2 {
			// error state
			s.lineHor(r, c)
			return
		}
		l1, a1 = s.verAt(r-2, c)
		if l1 {
			s.avoidVer(r+1, c)
		} else if !a1 {
			l2, _ = s.verAt(r+1, c)
			if l2 {
				s.avoidVer(r-2, c)
			}
		}
	}
}

func (s *state) checkPins() {
	for r := 1; r <= int(s.size); r++ {
		for c := 1; c <= int(s.size); c++ {
			s.checkPin(r, c)
		}
	}
}

func (s *state) checkPin(r, c int) {
	rl, ra := s.horAt(r, c)
	dl, da := s.verAt(r, c)
	ll, la := s.horAt(r, c-1)
	ul, ua := s.verAt(r-1, c)

	var nl, na, dir uint8
	if rl {
		nl++
		dir |= 1
	}
	if ra {
		na++
		dir |= 1
	}
	if dl {
		nl++
		dir |= 1 << 1
	}
	if da {
		na++
		dir |= 1 << 1
	}
	if ll {
		nl++
		dir |= 1 << 2
	}
	if la {
		na++
		dir |= 1 << 2
	}
	if ul {
		nl++
		dir |= 1 << 3
	}
	if ua {
		na++
		dir |= 1 << 3
	}

	if nl != 2 && nl+na != 3 {
		return
	}

	if dir&1 == 0 {
		if nl == 1 {
			s.lineHor(r, c)
		} else {
			s.avoidHor(r, c)
		}
	}
	if dir&(1<<1) == 0 {
		if nl == 1 {
			s.lineVer(r, c)
		} else {
			s.avoidVer(r, c)
		}
	}
	if dir&(1<<2) == 0 {
		if nl == 1 {
			s.lineHor(r, c-1)
		} else {
			s.avoidHor(r, c-1)
		}
	}
	if dir&(1<<3) == 0 {
		if nl == 1 {
			s.lineVer(r-1, c)
		} else {
			s.avoidVer(r-1, c)
		}
	}
}

func (s *state) horAt(r, c int) (bool, bool) {
	return s.horizontalLines[r]&(1<<c) != 0, s.horizontalAvoids[r]&(1<<c) != 0
}

func (s *state) avoidHor(r, c int) {
	s.horizontalAvoids[r] |= (1 << c)
}

func (s *state) lineHor(r, c int) {
	s.horizontalLines[r] |= (1 << c)
}

func (s *state) verAt(r, c int) (bool, bool) {
	return s.verticalLines[c]&(1<<r) != 0, s.verticalAvoids[c]&(1<<r) != 0
}

func (s *state) avoidVer(r, c int) {
	s.verticalAvoids[c] |= (1 << r)
}

func (s *state) lineVer(r, c int) {
	s.verticalLines[c] |= (1 << r)
}

func (s *state) String() string {
	var sb strings.Builder

	var isLine, isAvoid bool

	for r := 0; r <= int(s.size+1); r++ {
		for c := 0; c <= int(s.size+1); c++ {
			sb.WriteByte(s.getNode(r, c))
			sb.WriteByte(' ')
			isLine, isAvoid = s.horAt(r, c)
			if isLine && isAvoid {
				sb.WriteByte('@')
			} else if isLine {
				sb.WriteByte('-')
			} else if isAvoid {
				sb.WriteByte('X')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := 0; c <= int(s.size+1); c++ {
			isLine, isAvoid = s.verAt(r, c)
			if isLine && isAvoid {
				sb.WriteByte('@')
			} else if isLine {
				sb.WriteByte('|')
			} else if isAvoid {
				sb.WriteByte('X')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
			sb.WriteByte(' ')
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (s *state) getNode(r, c int) byte {
	for _, n := range s.nodes {
		if n.Row != r || n.Col != c {
			continue
		}
		if n.IsBlack {
			return 'B'
		}
		return 'W'
	}
	if r == 0 || c == 0 || r > int(s.size) || c > int(s.size) {
		return ' '
	}
	return '*'
}
