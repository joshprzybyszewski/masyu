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
	fmt.Printf("state:\n%s\n", s)

	s.settleNodes()

	fmt.Printf("state:\n%s\n", s)
}

func (s *state) settleNodes() {
	for s.checkNodes() {
		fmt.Printf("state:\n%s\n\n", s)
	}
}

func (s *state) checkNodes() bool {
	changed := false

	for _, n := range s.nodes {
		if n.IsBlack {
			changed = changed || s.checkBlack(n.Row, n.Col)
		} else {
			changed = changed || s.checkWhite(n.Row, n.Col)
		}
	}

	return changed
}

func (s *state) checkBlack(
	r, c int,
) bool {
	changed := false
	rl, ra := s.horAt(r, c)
	if rl {
		changed = changed || s.avoidHor(r, c-1)
	} else if ra {
		changed = changed || s.lineHor(r, c-1)
	}

	dl, da := s.verAt(r, c)
	if dl {
		changed = changed || s.avoidVer(r-1, c)
	} else if da {
		changed = changed || s.lineVer(r-1, c)
	}

	ll, la := s.horAt(r, c-1)
	if ll {
		changed = changed || s.avoidHor(r, c)
	} else if la {
		changed = changed || s.lineHor(r, c)
	}

	ul, ua := s.verAt(r-1, c)
	if ul {
		changed = changed || s.avoidVer(r, c)
	} else if ua {
		changed = changed || s.lineVer(r, c)
	}

	return changed
}

func (s *state) checkWhite(
	r, c int,
) bool {
	changed := false
	rl, ra := s.horAt(r, c)
	if rl {
		changed = changed || s.lineHor(r, c-1)
	} else if ra {
		changed = changed || s.avoidHor(r, c-1)
	}

	dl, da := s.verAt(r, c)
	if dl {
		changed = changed || s.lineVer(r-1, c)
	} else if da {
		changed = changed || s.avoidVer(r-1, c)
	}

	ll, la := s.horAt(r, c-1)
	if ll {
		changed = changed || s.lineHor(r, c)
	} else if la {
		changed = changed || s.avoidHor(r, c)
	}

	ul, ua := s.verAt(r-1, c)
	if ul {
		changed = changed || s.lineVer(r, c)
	} else if ua {
		changed = changed || s.avoidVer(r, c)
	}

	return changed
}

func (s *state) horAt(r, c int) (bool, bool) {
	return s.horizontalLines[r]&(1<<c) != 0, s.horizontalAvoids[r]&(1<<c) != 0
}

func (s *state) avoidHor(r, c int) bool {
	l, a := s.horAt(r, c)
	if l || a {
		return false
	}
	s.horizontalAvoids[r] |= (1 << c)
	return true
}

func (s *state) lineHor(r, c int) bool {
	l, a := s.horAt(r, c)
	if l || a {
		return false
	}
	s.horizontalLines[r] |= (1 << c)
	return true
}

func (s *state) verAt(r, c int) (bool, bool) {
	return s.verticalLines[c]&(1<<r) != 0, s.verticalAvoids[c]&(1<<r) != 0
}

func (s *state) avoidVer(r, c int) bool {
	l, a := s.verAt(r, c)
	if l || a {
		return false
	}
	s.verticalAvoids[c] |= (1 << r)
	return true
}

func (s *state) lineVer(r, c int) bool {
	l, a := s.verAt(r, c)
	if l || a {
		return false
	}
	s.verticalLines[c] |= (1 << r)
	return true
}

func (s *state) String() string {
	var sb strings.Builder

	var isLine, isAvoid bool

	for r := 0; r <= int(s.size+1); r++ {
		for c := 0; c <= int(s.size+1); c++ {
			sb.WriteByte(s.getNode(r, c))
			sb.WriteByte(' ')
			isLine, isAvoid = s.horAt(r, c)
			if isLine {
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
			if isLine {
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
