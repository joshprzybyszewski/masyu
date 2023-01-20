package model

import "strings"

type Solution struct {
	Size Size

	Horizontals [MaxPointsPerLine]uint64
	Verticals   [MaxPointsPerLine]uint64
}

func (s *Solution) String() string {
	var sb strings.Builder

	for r := Dimension(0); r < Dimension(s.Size); r++ {
		for c := Dimension(0); c < Dimension(s.Size); c++ {
			sb.WriteByte('*')
			sb.WriteByte(' ')
			if s.Horizontals[r]&c.Bit() != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := Dimension(0); c < Dimension(s.Size); c++ {
			if s.Verticals[c]&r.Bit() != 0 {
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

func (s *Solution) Pretty(
	nodes []Node,
) string {
	var sb strings.Builder

	for r := Dimension(0); r < Dimension(s.Size); r++ {
		for c := Dimension(0); c < Dimension(s.Size); c++ {
			sb.WriteByte(getNodeChar(nodes, r, c, s.Size))
			sb.WriteByte(' ')
			if s.Horizontals[r]&c.Bit() != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := Dimension(0); c < Dimension(s.Size); c++ {
			if s.Verticals[c]&r.Bit() != 0 {
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

func getNodeChar(
	nodes []Node,
	r, c Dimension,
	size Size,
) byte {
	for _, n := range nodes {
		if n.Row != r || n.Col != c {
			continue
		}
		if n.IsBlack {
			return 'B'
		}
		return 'W'
	}
	return '*'
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
