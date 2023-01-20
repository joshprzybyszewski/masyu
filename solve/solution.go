package solve

import (
	"strings"

	"github.com/joshprzybyszewski/masyu/model"
)

type Solution struct {
	size model.Size

	hor [model.MaxPointsPerLine]uint64
	ver [model.MaxPointsPerLine]uint64
}

func (s *Solution) String() string {
	var sb strings.Builder

	for r := 0; r < int(s.size); r++ {
		for c := 0; c < int(s.size); c++ {
			sb.WriteByte('*')
			sb.WriteByte(' ')
			if s.hor[r]&(1<<c) != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')

		for c := 0; c < int(s.size); c++ {
			if s.ver[c]&(1<<r) != 0 {
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
