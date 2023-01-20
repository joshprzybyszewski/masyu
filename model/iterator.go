package model

type Iterator int

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) GetSize() Size {
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
	if i == 13 {
		return daily
	}
	if i == 14 {
		return weekly
	}
	if i == 15 {
		return monthly
	}
	if i <= 18 {
		return twentyfive
	}

	return invalidSize
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i == 0 || (i >= 13 && i <= 15) {
		return easy
	}
	if i > 15 {
		return Difficulty(i - 15)
	}

	return Difficulty((i-1)%3) + 1
}
