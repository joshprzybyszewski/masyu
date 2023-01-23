package solve

import (
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

type state struct {
	rules *ruleCheckCollector
	nodes []model.Node

	size           model.Size
	lastLinePlaced model.Coord
	hasInvalid     bool

	paths pathCollector

	horizontalLines  [model.MaxPointsPerLine]uint64
	horizontalAvoids [model.MaxPointsPerLine]uint64

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
		nodes: make([]model.Node, len(ns)),
		size:  size,
		rules: &rcc,
	}

	// offset all of the input nodes by positive one
	for i := range ns {
		s.nodes[i] = ns[i]
		s.nodes[i].Row++
		s.nodes[i].Col++
	}
	s.paths = newPathCollector(s.nodes)

	s.lastLinePlaced = s.nodes[0].Coord

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

	if !s.rules.runAllChecks(s) {
		panic(`state initialization is not valid?`)
	}
}

func (s *state) isValidAndSolved() (bool, bool) {
	if !s.rules.runAllChecks(s) {
		return false, false
	}

	if !s.paths.hasCycle {
		// is not completed yet
		return true, false
	}

	if s.paths.cycleSeenNodes != len(s.nodes) {
		// there's a cycle, but it doesn't include all of the nodes.
		// it's invalid.
		return false, false
	}

	return true, true
}

func (s *state) toSolution() (model.Solution, bool) {
	// we found a state that includes a cycle with all of the nodes.
	// avoid all of the remaining spots in the state, and see if it's
	// still valid: this eliminates the bad state of having tertiary paths set.
	for row := model.Dimension(0); row <= model.Dimension(s.size+1); row++ {
		for col := model.Dimension(0); col <= model.Dimension(s.size+1); col++ {
			if !s.horLineAt(row, col) {
				s.avoidHor(row, col)
			}
			if !s.verLineAt(row, col) {
				s.avoidVer(row, col)
			}
		}
	}

	// and force a re-check of all the rules.
	for row := model.Dimension(0); row <= model.Dimension(s.size+1); row++ {
		for col := model.Dimension(0); col <= model.Dimension(s.size+1); col++ {
			s.rules.checkHorizontal(row, col, s)
			s.rules.checkVertical(row, col, s)
		}
	}

	if !s.rules.runAllChecks(s) {
		return model.Solution{}, false
	}

	sol := model.Solution{
		Size: s.size,
	}

	for i := 0; i < int(s.size); i++ {
		sol.Horizontals[i] = (s.horizontalLines[i+1]) >> 1
		sol.Verticals[i] = (s.verticalLines[i+1]) >> 1
	}

	return sol, true
}

func (s *state) isValid() bool {

	if s.hasInvalid {
		return false
	}

	if s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes) {
		return false
	}

	// for i := 0; i <= int(s.size); i++ {
	// 	if (s.horizontalAvoids[i])&(s.horizontalLines[i]) != 0 ||
	// 		(s.verticalLines[i])&(s.verticalAvoids[i]) != 0 {
	// 		return false
	// 	}
	// }

	return true
}

func (s *state) getMostInterestingPath() (model.Coord, bool, bool) {
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

	s.lastLinePlaced.Row = r
	s.lastLinePlaced.Col = c

	s.rules.checkHorizontal(r, c, s)
	s.paths.addHorizontal(r, c)
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

	s.lastLinePlaced.Row = r
	s.lastLinePlaced.Col = c

	s.rules.checkVertical(r, c, s)
	s.paths.addVertical(r, c)
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
