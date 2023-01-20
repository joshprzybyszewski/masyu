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

func (i input) ToNodes() []model.Node {
	var r, c model.Dimension
	max := model.Dimension(i.size)
	output := make([]model.Node, 0, len(i.task)/2)

	for _, b := range i.task {
		if b == 'B' || b == 'W' {
			output = append(output, model.Node{
				Coord: model.Coord{
					Row: r,
					Col: c,
				},
				IsBlack: b == 'B',
			})
		} else {
			c += model.Dimension(b - 'a')
		}

		c++

		if c >= max {
			r += (c / max)
			c %= max
		}
	}

	return output
}
