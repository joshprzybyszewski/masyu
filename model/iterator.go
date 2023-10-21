package model

import "fmt"

type Iterator int

func (i Iterator) String() string {
	return fmt.Sprintf("%dx%d %s", i.GetSize()-1, i.GetSize()-1, i.GetDifficulty())
}

func (i Iterator) Valid() bool {
	return MinIterator <= i && i <= MaxIterator
}

func (i Iterator) GetSize() Size {
	if i < MinIterator || i > MaxIterator {
		return invalidSize
	}

	if i <= 1 {
		return five
	}
	if i <= 4 {
		return seven
	}
	if i <= 7 {
		return ten
	}
	if i <= 10 {
		return fifteen
	}
	if i <= 13 {
		return twenty
	}
	if i <= 16 {
		return twentyfive
	}
	if i == 17 {
		return daily
	}
	if i == 18 {
		return weekly
	}
	if i == 19 {
		return monthly
	}

	return invalidSize
}

func (i Iterator) GetDifficulty() Difficulty {
	if i < MinIterator || i > MaxIterator {
		return invalidDifficulty
	}

	if i <= 1 {
		return Difficulty(i + 1)
	}
	if i >= 17 && i <= 19 {
		return hard
	}
	if i > 15 {
		return Difficulty(i - 15)
	}

	// 2 is the easy 7x7, and then it increments by threes
	return Difficulty((i-2)%3) + 1
}
