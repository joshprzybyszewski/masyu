package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type rule struct {
	check func(*state)

	row model.Dimension
	col model.Dimension
}
