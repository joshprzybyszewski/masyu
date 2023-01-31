package model

type Dimension uint8

type DimensionBit uint64

func (d Dimension) Bit() DimensionBit {
	return 1 << d
}

type Coord struct {
	Row Dimension
	Col Dimension
}
