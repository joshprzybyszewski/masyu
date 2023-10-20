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

	// example: aB4B3eW2gB3W5fB5gB2bB6
	// example: dB4lB8dB2aB2dB2B5bB6aW5cB5bB3eB2gB7gB3eB3bW2dB6gW4aW3B3cW4bW3fW3fW2gW4dB6aB5hB5B2B3cB5jW4aB6B4aB4dW2eW4bW2dB6bW3B6aB3bW2aB3bW3gB3jB4dB4bB4bB2W2bB2B3dW3aB2aW3bB4B2iB4rB3cB2dB2fB2aW2bB7eW5aB2aW3aB3cB2iB2aB2B2fB3bB2dB5W2aB2fB4eW2B3fB5aB2aW4cB7aB2B2cW9bB2hB2fB5gB3cB9dB10kB4W3gW7cB3gB2a
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
