package fetch

import (
	"fmt"

	"github.com/joshprzybyszewski/masyu/model"
)

type input struct {
	ID    string
	param string
	task  string

	iter model.Iterator
}

func (i input) String() string {
	return fmt.Sprintf("Puzzle %s (Iter: %d, Size: %d, Difficulty: %s)",
		i.ID,
		i.iter,
		i.iter.GetSize(),
		i.iter.GetDifficulty(),
	)
}

func (i input) Task() string {
	return i.task
}

func (i input) ToNodes() []model.Node {
	var r, c model.Dimension
	max := model.Dimension(i.iter.GetSize())
	output := make([]model.Node, 0, len(i.task)/2)

	// TODO consider replacing with a strings.NewReader(i.task)?
	taskBytes := []byte(i.task)
	oi, vi := 0, 0
	maxI := len(taskBytes)
	var b byte
	for i := 0; i < maxI; i++ {
		b = taskBytes[i]
		if b == 'B' || b == 'W' {
			output = append(output, model.Node{
				Coord: model.Coord{
					Row: r,
					Col: c,
				},
				IsBlack: b == 'B',
			})

			// get the value of the node too
			for vi = i + 1; vi < maxI; vi++ {
				b = taskBytes[vi]
				if b < '0' || b > '9' {
					break
				}
				// we're using this byte for the node's value. Increment the standard i too.
				i++
				if output[oi].Value != 0 {
					// optimization: don't multiply if the value is zero.
					// this is definitely a pre-optimization:#
					output[oi].Value *= 10
				}
				output[oi].Value += model.Value(b - '0')
			}

			// increase the index of the output node
			oi++
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
