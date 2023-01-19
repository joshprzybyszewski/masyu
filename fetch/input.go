package fetch

import "github.com/joshprzybyszewski/masyu/model"

type input struct {
	id    string
	param string
	task  string

	size       model.Size
	difficulty model.Difficulty
}
