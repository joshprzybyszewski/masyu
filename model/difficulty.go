package model

type Difficulty uint8

const (
	invalidDifficulty Difficulty = 0

	easy   Difficulty = 1
	medium Difficulty = 2
	hard   Difficulty = 3
)
