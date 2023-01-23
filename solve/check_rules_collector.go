package solve

import "github.com/joshprzybyszewski/masyu/model"

type ruleCheckCollector struct {
	rules *rules

	hor [model.MaxPointsPerLine]uint64
	ver [model.MaxPointsPerLine]uint64
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
	c.hor[row] |= col.Bit()
}

func (c *ruleCheckCollector) checkVertical(
	row, col model.Dimension,
) {
	c.ver[col] |= row.Bit()
}

func (c *ruleCheckCollector) runAllChecks(
	s *state,
) bool {

	for c.hasPending() {
		c.flush(s)
		if !s.isValid() {
			return false
		}
	}

	return s.isValid()
}

func (c *ruleCheckCollector) hasPending() bool {
	for i := 0; i < model.MaxPointsPerLine; i++ {
		if c.hor[i] != 0 || c.ver[i] != 0 {
			return true
		}
	}
	return false
}

func (c *ruleCheckCollector) flush(
	s *state,
) {
	hor := c.hor
	ver := c.ver

	im := int(s.size) + 3

	for row := model.Dimension(0); row < model.Dimension(im); row++ {
		if hor[row] == 0 {
			continue
		}
		c.hor[row] = 0
		for col := model.Dimension(0); col < model.Dimension(im); col++ {
			if hor[row]&col.Bit() != 0 {
				c.rules.checkHorizontal(row, col, s)
			}
		}
	}

	for col := model.Dimension(0); col < model.Dimension(im); col++ {
		if ver[col] == 0 {
			continue
		}
		c.ver[col] = 0
		for row := model.Dimension(0); row < model.Dimension(im); row++ {
			if ver[col]&row.Bit() != 0 {
				c.rules.checkVertical(row, col, s)
			}
		}
	}
}
