package solve

import (
	"fmt"
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	all64Bits uint64 = 0xFFFFFFFFFFFFFFFF
)

type state struct {
	rules *ruleCheckCollector
	nodes []model.Node

	size       model.Size
	hasInvalid bool

	paths pathCollector

	crossings crossings

	// [row]colBitMask
	horizontalLines  [model.MaxPointsPerLine]uint64
	horizontalAvoids [model.MaxPointsPerLine]uint64

	// [col]rowBitMask
	verticalLines  [model.MaxPointsPerLine]uint64
	verticalAvoids [model.MaxPointsPerLine]uint64
}

func newState(
	size model.Size,
	ns []model.Node,
) state {

	r := newRules(size)
	rcc := newRuleCheckCollector(r)

	s := state{
		nodes:     make([]model.Node, len(ns)),
		size:      size,
		crossings: newCrossings(size),
		rules:     &rcc,
	}

	// offset all of the input nodes by positive one
	for i := range ns {
		s.nodes[i] = ns[i]
		s.nodes[i].Row++
		s.nodes[i].Col++
	}
	s.paths = newPathCollector(s.nodes)

	findGimmes(&s)

	r.populateRules(&s)

	s.initialize()

	r.populateUnknowns(&s)

	return s
}

func (s *state) initialize() {
	avoid := uint64(1 | (1 << s.size))
	for i := 1; i <= int(s.size); i++ {
		s.verticalAvoids[i] |= avoid
		s.horizontalAvoids[i] |= avoid
	}
	s.horizontalAvoids[0] = all64Bits
	s.verticalAvoids[0] = all64Bits
	s.horizontalAvoids[s.size+1] = all64Bits
	s.verticalAvoids[s.size+1] = all64Bits

	if !checkEntireRuleset(s) {
		fmt.Printf("Invalid State:\n%s\n", s)
		panic(`state initialization is not valid?`)
	}
}

func (s *state) toSolution() model.Solution {
	sol := model.Solution{
		Size: s.size,
	}

	// each line needs to be shifted by one.
	for i := 0; i < int(s.size); i++ {
		sol.Horizontals[i] = (s.horizontalLines[i+1]) >> 1
		sol.Verticals[i] = (s.verticalLines[i+1]) >> 1
	}

	return sol
}

func (s *state) isValid() bool {

	if s.hasInvalid {
		return false
	}

	if s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes) {
		return false
	}

	return true
}

func (s *state) getMostInterestingPath() (model.Coord, bool, bool) {
	c, isHor, ok := s.paths.getInteresting(s)
	if ok {
		return c, isHor, true
	}

	var l, a bool
	for _, pp := range s.rules.rules.unknowns {
		if pp.IsHorizontal {
			if l, a = s.horAt(pp.Row, pp.Col); !l && !a {
				return pp.Coord, pp.IsHorizontal, true
			}
		} else {
			if l, a = s.verAt(pp.Row, pp.Col); !l && !a {
				return pp.Coord, pp.IsHorizontal, true
			}
		}
	}
	// there are no more interesting paths left. Likely this means that there's
	// an error in the state and we need to abort.
	return model.Coord{}, false, false
}

func (s *state) horAt(r, c model.Dimension) (bool, bool) {
	return s.horLineAt(r, c), s.horAvoidAt(r, c)
}

func (s *state) horLineAt(r, c model.Dimension) bool {
	return s.horizontalLines[r]&c.Bit() != 0
}

func (s *state) horAvoidAt(r, c model.Dimension) bool {
	return s.horizontalAvoids[r]&c.Bit() != 0
}

func (s *state) avoidHor(r, c model.Dimension) {
	b := c.Bit()
	if s.horizontalAvoids[r]&b == b {
		// already avoided
		return
	}
	s.horizontalAvoids[r] |= b
	if s.horizontalLines[r]&s.horizontalAvoids[r] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.crossings.avoidHor(c)
	s.rules.checkHorizontal(r, c, s)
}

func (s *state) lineHor(r, c model.Dimension) {
	b := c.Bit()
	if s.horizontalLines[r]&b == b {
		// already a line
		return
	}
	s.horizontalLines[r] |= b
	if s.horizontalLines[r]&s.horizontalAvoids[r] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.crossings.lineHor(c)
	s.rules.checkHorizontal(r, c, s)
	s.paths.addHorizontal(r, c, s)
}

func (s *state) verAt(r, c model.Dimension) (bool, bool) {
	return s.verLineAt(r, c), s.verAvoidAt(r, c)
}

func (s *state) verLineAt(r, c model.Dimension) bool {
	return s.verticalLines[c]&r.Bit() != 0
}

func (s *state) verAvoidAt(r, c model.Dimension) bool {
	return s.verticalAvoids[c]&r.Bit() != 0
}

func (s *state) avoidVer(r, c model.Dimension) {
	b := r.Bit()
	if s.verticalAvoids[c]&b == b {
		// already avoided
		return
	}
	s.verticalAvoids[c] |= b
	if s.verticalLines[c]&s.verticalAvoids[c] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.crossings.avoidVer(r)
	s.rules.checkVertical(r, c, s)
}

func (s *state) lineVer(r, c model.Dimension) {
	b := r.Bit()
	if s.verticalLines[c]&b == b {
		// already avoided
		return
	}
	s.verticalLines[c] |= b
	if s.verticalLines[c]&s.verticalAvoids[c] != 0 {
		// invalid
		s.hasInvalid = true
		return
	}

	s.crossings.lineVer(r)
	s.rules.checkVertical(r, c, s)
	s.paths.addVertical(r, c, s)
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
	if r == 0 {
		return '0' + byte(c%10)
	}
	if c == 0 {
		return '0' + byte(r%10)
	}
	if r > model.Dimension(s.size) || c > model.Dimension(s.size) {
		return ' '
	}
	return '*'
}
