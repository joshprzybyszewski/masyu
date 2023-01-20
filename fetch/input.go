package fetch

import "github.com/joshprzybyszewski/masyu/model"

type input struct {
	id    string
	param string
	task  string

	iter       model.Iterator
	size       model.Size
	difficulty model.Difficulty
}
