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

	rules *rules
}

func newState(
	size model.Size,
	ns []model.Node,
) state {
	s := state{
		nodes: make([]model.Node, len(ns)),
		size:  size,
		rules: newRules(size),
	}

	// offset all of the input nodes by positive one
	for i := range ns {
		s.nodes[i] = ns[i]
		s.nodes[i].Row++
		s.nodes[i].Col++

		if ns[i].IsBlack {
			s.rules.addBlackNode(s.nodes[i].Row, s.nodes[i].Col)
		} else {
			s.rules.addWhiteNode(s.nodes[i].Row, s.nodes[i].Col)
		}
	}

	s.initialize()

	return s
}

func (s *state) toSolution() (model.Solution, model.Coord, bool, bool) {
	if !s.isValid() {
		return model.Solution{}, model.Coord{}, false, false
	}

	eop, isValid, complete := s.isValidPath()
	if !isValid {
		return model.Solution{}, model.Coord{}, false, false
	}

	if !complete {
		return model.Solution{}, eop, false, true
	}

	sol := model.Solution{
		Size: s.size,
	}

	for i := 0; i < int(s.size); i++ {
		sol.Horizontals[i] = (s.horizontalLines[i+1]) >> 1
		sol.Verticals[i] = (s.verticalLines[i+1]) >> 1
	}

	return sol, model.Coord{}, true, true
}

func (s *state) isValidPath() (model.Coord, bool, bool) {
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
			return cur, true, false
		}
		if cur == start {
			break
		}
	}
	if cur == prev || cur != start {
		return cur, true, false
	}

	if horizontalLines != s.horizontalLines || verticalLines != s.verticalLines {
		return model.Coord{}, false, false
	}

	return model.Coord{}, true, true
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

	for row := model.Dimension(0); row <= model.Dimension(s.size+1); row++ {
		for col := model.Dimension(0); col <= model.Dimension(s.size+1); col++ {
			s.rules.checkHorizontal(row, col, s)
			s.rules.checkVertical(row, col, s)
		}
	}

	if !s.isValid() {
		panic(`dev error`)
	}
}

func (s *state) checkNodes() {
	return
}

func (s *state) horAt(r, c model.Dimension) (bool, bool) {
	return s.horizontalLines[r]&c.Bit() != 0, s.horizontalAvoids[r]&c.Bit() != 0
}

func (s *state) avoidHor(r, c model.Dimension) {
	b := c.Bit()
	if s.horizontalAvoids[r]&b == b {
		// already avoided
		return
	}
	s.horizontalAvoids[r] |= b
	if s.horizontalLines[r]&b == 0 {
		// still valid; check the rules
		s.rules.checkHorizontal(r, c, s)
	}
}

func (s *state) lineHor(r, c model.Dimension) {
	b := c.Bit()
	if s.horizontalLines[r]&b == b {
		// already avoided
		return
	}
	s.horizontalLines[r] |= b
	if s.horizontalAvoids[r]&b == 0 {
		// still valid; check the rules
		s.rules.checkHorizontal(r, c, s)
	}
}

func (s *state) verAt(r, c model.Dimension) (bool, bool) {
	return s.verticalLines[c]&r.Bit() != 0, s.verticalAvoids[c]&r.Bit() != 0
}

func (s *state) avoidVer(r, c model.Dimension) {
	b := r.Bit()
	if s.verticalAvoids[c]&b == b {
		// already avoided
		return
	}
	s.verticalAvoids[c] |= b
	if s.verticalLines[c]&b == 0 {
		// still valid; check the rules
		s.rules.checkVertical(r, c, s)
	}
}

func (s *state) lineVer(r, c model.Dimension) {
	b := r.Bit()
	if s.verticalLines[c]&b == b {
		// already avoided
		return
	}
	s.verticalLines[c] |= b
	if s.verticalAvoids[c]&b == 0 {
		// still valid; check the rules
		s.rules.checkVertical(r, c, s)
	}
}

const (
	confusedSpace       byte = '@'
	horizontalLineSpace byte = '-'
	verticalLineSpace   byte = '|'
	avoidSpace          byte = 'X'
)

func (s *state) String() string {
	var sb strings.Builder

	var isLine, isAvoid bool

	for r := 0; r <= int(s.size+1); r++ {
		for c := 0; c <= int(s.size+1); c++ {
			sb.WriteByte(s.getNode(model.Dimension(r), model.Dimension(c)))
			sb.WriteByte(' ')
			isLine, isAvoid = s.horAt(model.Dimension(r), model.Dimension(c))
			if isLine && isAvoid {
				sb.WriteByte(confusedSpace)
			} else if isLine {
				sb.WriteByte(horizontalLineSpace)
			} else if isAvoid {
				sb.WriteByte(avoidSpace)
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := 0; c <= int(s.size+1); c++ {
			isLine, isAvoid = s.verAt(model.Dimension(r), model.Dimension(c))
			if isLine && isAvoid {
				sb.WriteByte(confusedSpace)
			} else if isLine {
				sb.WriteByte(verticalLineSpace)
			} else if isAvoid {
				sb.WriteByte(avoidSpace)
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

func (s *state) getNode(r, c model.Dimension) byte {
	for _, n := range s.nodes {
		if n.Row != r || n.Col != c {
			continue
		}
		if n.IsBlack {
			return 'B'
		}
		return 'W'
	}
	if r == 0 || c == 0 || r > model.Dimension(s.size) || c > model.Dimension(s.size) {
		return ' '
	}
	return '*'
}
