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

	if s.hasHorDefined(n.Row, n.Col-2) {
		return false
	}
	if s.hasHorDefined(n.Row, n.Col-1) {
		return false
	}
	if s.hasHorDefined(n.Row, n.Col) {
		return false
	}
	if s.hasHorDefined(n.Row, n.Col+1) {
		return false
	}

	if s.hasVerDefined(n.Row-2, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row-1, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row+1, n.Col) {
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

	if s.hasVerDefined(n.Row-1, n.Col-1) {
		return false
	}
	if s.hasVerDefined(n.Row, n.Col-1) {
		return false
	}
	if s.hasHorDefined(n.Row, n.Col-1) {
		return false
	}
	if s.hasHorDefined(n.Row, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row-1, n.Col+1) {
		return false
	}
	if s.hasVerDefined(n.Row, n.Col+1) {
		return false
	}

	if s.hasHorDefined(n.Row-1, n.Col-1) {
		return false
	}
	if s.hasHorDefined(n.Row-1, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row-1, n.Col) {
		return false
	}
	if s.hasVerDefined(n.Row, n.Col) {
		return false
	}
	if s.hasHorDefined(n.Row+1, n.Col-1) {
		return false
	}
	if s.hasHorDefined(n.Row+1, n.Col) {
		return false
	}

	return true
}
