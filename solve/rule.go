package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type rule struct {
	check func(*state)

	row model.Dimension
	col model.Dimension
}

func (r *rule) setInvalid(s *state) {
	s.lineHor(r.row, r.col)
	s.avoidHor(r.row, r.col)
}
