package fetch

type size uint8

const (
	invalidSize size = 0

	six        size = 6
	eight      size = 8
	ten        size = 10
	fifteen    size = 15
	twenty     size = 20
	twentyfive size = 25
	daily      size = 30
	weekly     size = 35
	monthly    size = 40
)
