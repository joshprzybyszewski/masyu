package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type rule struct {
	check   func(*state)
	affects int

	row model.Dimension
	col model.Dimension
}

func (r *rule) setInvalid(s *state) {
	// fmt.Printf("s.String(): %s\n", s.String())
	s.hasInvalid = true
}
