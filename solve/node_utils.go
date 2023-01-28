package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

func isBlackNodeClear(
	s *state,
	n model.Node,
) bool {

	if !n.IsBlack {
		return false
	}

	if n.Row < 2 || n.Col < 2 {
		return false
	}

	if n.Row >= model.Dimension(s.size)-1 || n.Col >= model.Dimension(s.size)-1 {
		return false
	}

	l, a := s.horAt(n.Row, n.Col-2)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col+1)
	if l || a {
		return false
	}

	l, a = s.verAt(n.Row-2, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row+1, n.Col)
	if l || a {
		return false
	}

	return true
}

func isWhiteNodeClear(
	s *state,
	n model.Node,
) bool {

	if n.IsBlack {
		return false
	}

	if n.Row < 2 || n.Col < 2 {
		return false
	}

	if n.Row >= model.Dimension(s.size)-1 || n.Col >= model.Dimension(s.size)-1 {
		return false
	}

	l, a := s.verAt(n.Row-1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col+1)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col+1)
	if l || a {
		return false
	}

	l, a = s.horAt(n.Row-1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row+1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row+1, n.Col)
	if l || a {
		return false
	}

	return true
}
