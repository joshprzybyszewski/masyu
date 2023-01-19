package fetch

const (
	MinIterator = 0
	MaxIterator = 18
)

type Iterator int

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) size() size {
	if i < MinIterator || i > MaxIterator {
		return invalidSize
	}

	if i == 0 {
		return six
	}
	if i <= 3 {
		return eight
	}
	if i <= 6 {
		return ten
	}
	if i <= 9 {
		return fifteen
	}
	if i <= 12 {
		return twenty
	}
	if i <= 15 {
		return twentyfive
	}
	if i == 16 {
		return daily
	}
	if i == 17 {
		return weekly
	}
	if i == 18 {
		return monthly
	}

	return invalidSize
}

func (i Iterator) difficulty() difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i == 0 || i >= 16 {
		return easy
	}

	return difficulty((i - 1) % 3)
}
