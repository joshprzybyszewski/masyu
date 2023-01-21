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
			sb.WriteByte(getNodeChar(nodes, r, c))
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

	for row := Dimension(0); row <= Dimension(numEdges); row++ {
		for col := Dimension(0); col < Dimension(numEdges); col++ {
			if s.Horizontals[row]&col.Bit() != 0 {
				_ = rows.WriteByte('y')
			} else {
				_ = rows.WriteByte('n')
			}
		}
	}

	for row := Dimension(0); row < Dimension(numEdges); row++ {
		for col := Dimension(0); col <= Dimension(numEdges); col++ {
			if s.Verticals[col]&row.Bit() != 0 {
				_ = cols.WriteByte('y')
			} else {
				_ = cols.WriteByte('n')
			}
		}
	}

	return rows.String() + cols.String()
}
