package model

import "strings"

type Solution struct {
	Size Size

	Horizontals [MaxPointsPerLine]uint64
	Verticals   [MaxPointsPerLine]uint64
}

func (s *Solution) String() string {
	var sb strings.Builder

	for r := 0; r < int(s.Size); r++ {
		for c := 0; c < int(s.Size); c++ {
			sb.WriteByte('*')
			sb.WriteByte(' ')
			if s.Horizontals[r]&(1<<c) != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := 0; c < int(s.Size); c++ {
			if s.Verticals[c]&(1<<r) != 0 {
				sb.WriteByte('|')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
			sb.WriteByte(' ')
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (s *Solution) ToAnswer() string {

	numEdges := int(s.Size) - 1
	var rows, cols strings.Builder
	rows.Grow(numEdges * (numEdges - 1))
	cols.Grow(numEdges * (numEdges - 1))

	for dim1 := 0; dim1 <= numEdges; dim1++ {
		for dim2 := 0; dim2 <= numEdges; dim2++ {
			if dim2 < numEdges {
				if s.Horizontals[dim1]&(1<<dim2) != 0 {
					_ = rows.WriteByte('y')
				} else {
					_ = rows.WriteByte('n')
				}
			}

			if dim1 < numEdges {
				if s.Verticals[dim2]&(1<<dim1) != 0 {
					_ = cols.WriteByte('y')
				} else {
					_ = cols.WriteByte('n')
				}
			}
		}
	}

	return rows.String() + cols.String()
}
