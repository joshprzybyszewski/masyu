package model

import "strconv"

type Difficulty uint8

const (
	invalidDifficulty Difficulty = 0

	easy   Difficulty = 1
	medium Difficulty = 2
	hard   Difficulty = 3
)

func (d Difficulty) String() string {
	switch d {
	case easy:
		return `easy`
	case medium:
		return `medium`
	case hard:
		return `hard`
	default:
		return strconv.Itoa(int(d))
	}
}
