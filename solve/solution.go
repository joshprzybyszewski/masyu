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

func (s Solution) String() string {
	var sb strings.Builder

	for r := model.Size(0); r < s.size; r++ {
		for c := model.Size(0); c < s.size; c++ {
			sb.WriteByte(' ')
			if (s.hor[r] & 1 << c) != 0 {
				sb.WriteByte('-')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')

		if r == s.size-1 {
			break
		}

		for c := model.Size(0); c < s.size; c++ {
			if (s.ver[c] & 1 << r) != 0 {
				sb.WriteByte('|')
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
