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
) {
	c.hasPending = true
	c.hor[row] |= col.Bit()
}

func (c *ruleCheckCollector) checkVertical(
	row, col model.Dimension,
) {
	c.hasPending = true
	c.ver[col] |= row.Bit()
}

func (c *ruleCheckCollector) runAllChecks(
	s *state,
) settledState {
	if !s.isValid() {
		return invalid
	}

	var dim1, dim2 model.Dimension
	var bit, tmp uint64
	var r *rule
	im := model.Dimension(int(s.size) + 2)

	for c.hasPending {
		c.hasPending = false

		for dim1 = 0; dim1 < im; dim1++ {
			if c.hor[dim1] != 0 {
				bit = 1
				tmp = c.hor[dim1]
				c.hor[dim1] = 0
				for dim2 = 0; dim2 < im; dim2++ {
					if tmp&bit == bit {
						for _, r = range c.rules.horizontals[dim1][dim2] {
							r.check(s)
						}
					}
					bit <<= 1
				}
			}

			if c.ver[dim1] != 0 {
				bit = 1
				tmp = c.ver[dim1]
				c.ver[dim1] = 0
				for dim2 = 0; dim2 < im; dim2++ {
					if tmp&bit == bit {
						for _, r = range c.rules.verticals[dim2][dim1] {
							r.check(s)
						}
					}
					bit <<= 1
				}
			}
		}

		if !s.isValid() {
			return invalid
		}
	}

	if !s.isValid() {
		return invalid
	}

	return validUnsolved
}
