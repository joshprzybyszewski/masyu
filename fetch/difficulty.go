package fetch

type difficulty uint8

const (
	invalidDifficulty difficulty = 0

	easy   difficulty = 1
	medium difficulty = 2
	hard   difficulty = 3
)

func GetDifficulty(
	i int,
) (difficulty, bool) {
	if i < 0 || i > int(hard)-1 {
		return 0, false
	}
	return difficulty(i + 1), true
}
