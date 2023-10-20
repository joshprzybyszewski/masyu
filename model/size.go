package model

type Size uint8

const (
	invalidSize Size = 0

	five       Size = 6
	seven      Size = 8
	ten        Size = 11
	fifteen    Size = 16
	twenty     Size = 21
	twentyfive Size = 26
	daily      Size = 31
	weekly     Size = 36
	monthly    Size = 41
)
