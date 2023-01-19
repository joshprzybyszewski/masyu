package model

type Size uint8

const (
	invalidSize Size = 0

	six        Size = 6
	eight      Size = 8
	ten        Size = 10
	fifteen    Size = 15
	twenty     Size = 20
	twentyfive Size = 25
	daily      Size = 30
	weekly     Size = 35
	monthly    Size = 40
)
