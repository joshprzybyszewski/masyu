package model

type Dimension uint8

func (d Dimension) Bit() uint64 {
	return 1 << d
}

type Coord struct {
	Row Dimension
	Col Dimension
}
