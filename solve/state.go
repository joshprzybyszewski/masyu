package solve

import (
	"fmt"
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
	fmt.Printf("state BEFORE settle:\n%s\n", s)

	s.settleNodes()
	if !s.isValid() {
		fmt.Printf("error state:\n%s\n\n", s)
		panic(`dev error`)
	}

	fmt.Printf("state AFTER settle:\n%s\n", s)
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
		fmt.Printf("state %d:\n%s\n\n", i, s)

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
	l, a := s.horAt(r, c)
	if l {
		s.lineHor(r, c-1)
		s.avoidVer(r-1, c)
		s.avoidVer(r, c)
		return
	} else if a {
		s.lineVer(r, c)
		s.lineVer(r-1, c)
		s.avoidHor(r, c-1)
		return
	}

	l, a = s.verAt(r, c)
	if l {
		s.lineVer(r-1, c)
		s.avoidHor(r, c-1)
		s.avoidHor(r, c)
		return
	} else if a {
		s.lineHor(r, c-1)
		s.lineHor(r, c)
		s.avoidVer(r-1, c)
		return
	}

	l, a = s.horAt(r, c-1)
	if l {
		s.lineHor(r, c)
		s.avoidVer(r-1, c)
		s.avoidVer(r, c)
		return
	} else if a {
		s.avoidHor(r, c)
		s.lineVer(r, c)
		s.lineVer(r-1, c)
		return
	}

	l, a = s.verAt(r-1, c)
	if l {
		s.lineVer(r, c)
		s.avoidHor(r, c-1)
		s.avoidHor(r, c)
		return
	} else if a {
		s.lineHor(r, c-1)
		s.lineHor(r, c)
		s.avoidVer(r, c)
		return
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
