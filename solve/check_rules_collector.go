package solve

import "github.com/joshprzybyszewski/masyu/model"

type ruleCheckCollector struct {
	rules *rules

	hor [model.MaxPointsPerLine]uint64
	ver [model.MaxPointsPerLine]uint64

	hasPending bool
}

func newRuleCheckCollector(
	r *rules,
) ruleCheckCollector {
	return ruleCheckCollector{
		rules: r,
	}
}

func (c *ruleCheckCollector) checkHorizontal(
	row, col model.Dimension,
	s *state,
) {
	c.hasPending = true
	c.hor[row] |= col.Bit()
}

func (c *ruleCheckCollector) checkVertical(
	row, col model.Dimension,
	s *state,
) {
	c.hasPending = true
	c.ver[col] |= row.Bit()
}

func (c *ruleCheckCollector) runAllChecks(
	s *state,
) bool {

	var dim1, dim2 model.Dimension
	var bit uint64
	im := model.Dimension(int(s.size) + 2)

	for c.hasPending {
		c.hasPending = false

		for dim1 = 0; dim1 < im; dim1++ {
			if c.hor[dim1] != 0 {
				bit = 1
				for dim2 = 0; dim2 < im; dim2++ {
					if c.hor[dim1]&bit != 0 {
						c.hor[dim1] = c.hor[dim1] ^ bit
						c.rules.checkHorizontal(dim1, dim2, s)
						if !s.isValid() {
							return false
						}
					}
					bit <<= 1
				}
			}

			if c.ver[dim1] != 0 {
				bit = 1
				for dim2 = 0; dim2 < im; dim2++ {
					if c.ver[dim1]&bit != 0 {
						c.ver[dim1] = c.ver[dim1] ^ bit
						c.rules.checkVertical(dim2, dim1, s)
						if !s.isValid() {
							return false
						}
					}
					bit <<= 1
				}
			}
		}

		if !s.isValid() {
			return false
		}
	}

	return s.isValid()
}
