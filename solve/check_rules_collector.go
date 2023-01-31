package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

type ruleCheckCollector struct {
	rules *rules

	hor [model.MaxPointsPerLine]model.DimensionBit
	ver [model.MaxPointsPerLine]model.DimensionBit

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
	row model.Dimension,
	col model.DimensionBit,
) {
	c.hasPending = true
	c.hor[row] |= col
}

func (c *ruleCheckCollector) checkVertical(
	row model.DimensionBit,
	col model.Dimension,
) {
	c.hasPending = true
	c.ver[col] |= row
}

func (c *ruleCheckCollector) runAllChecks(
	s *state,
) settledState {
	if s.hasInvalid || (s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes)) {
		return invalid
	}

	var dim1, dim2 model.Dimension
	var tmp model.DimensionBit
	im := model.Dimension(int(s.size) + 2)

	for c.hasPending {
		c.hasPending = false

		for dim1 = 0; dim1 < im; dim1++ {
			if c.hor[dim1] != 0 {
				tmp = c.hor[dim1]
				c.hor[dim1] = 0
				for dim2 = 0; tmp != 0 && dim2 < im; dim2++ {
					if tmp&1 == 1 {
						c.rules.horizontals[dim1][dim2].fn(s)
					}
					tmp >>= 1
				}
			}

			if c.ver[dim1] != 0 {
				tmp = c.ver[dim1]
				c.ver[dim1] = 0
				for dim2 = 0; tmp != 0 && dim2 < im; dim2++ {
					if tmp&1 == 1 {
						c.rules.verticals[dim2][dim1].fn(s)
					}
					tmp >>= 1
				}
			}
			if s.hasInvalid || (s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes)) {
				return invalid
			}
		}

		if s.hasInvalid || (s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes)) {
			return invalid
		}
	}

	if s.hasInvalid || (s.paths.hasCycle && s.paths.cycleSeenNodes != len(s.nodes)) {
		return invalid
	}

	return validUnsolved
}
