package solve

import "github.com/joshprzybyszewski/masyu/model"

type crossings struct {
	// [col]numLines/AvoidsInThatCol
	cols      [model.MaxPointsPerLine]model.Dimension
	colsAvoid [model.MaxPointsPerLine]model.Dimension

	// [row]numLines/AvoidsInThatRow
	rows      [model.MaxPointsPerLine]model.Dimension
	rowsAvoid [model.MaxPointsPerLine]model.Dimension

	// This is the "target" for an _almost_ empty row/col
	target            model.Dimension
	hasNearlyComplete bool
}

func newCrossings(
	size model.Size,
) crossings {
	return crossings{
		target: model.Dimension(size) - 1,
	}
}

func (c *crossings) lineHor(
	col model.Dimension,
) {
	c.cols[col]++
	if c.colsAvoid[col]+c.cols[col] == c.target {
		c.hasNearlyComplete = true
	}
}

func (c *crossings) avoidHor(
	col model.Dimension,
) {
	c.colsAvoid[col]++
	if c.colsAvoid[col]+c.cols[col] == c.target {
		c.hasNearlyComplete = true
	}
}

func (c *crossings) lineVer(
	row model.Dimension,
) {
	c.rows[row]++
	if c.rowsAvoid[row]+c.rows[row] == c.target {
		c.hasNearlyComplete = true
	}
}

func (c *crossings) avoidVer(
	row model.Dimension,
) {
	c.rowsAvoid[row]++
	if c.rowsAvoid[row]+c.rows[row] == c.target {
		c.hasNearlyComplete = true
	}
}

func (c *crossings) complete(
	s *state,
) bool {
	if !c.hasNearlyComplete {
		return false
	}

	c.hasNearlyComplete = false
	changed := false

	var other model.Dimension

	for i := model.Dimension(1); i < model.Dimension(s.size); i++ {
		if c.cols[i]+c.colsAvoid[i] == c.target {
			changed = true
			other = getEmptyCrossingInColumn(s, i)
			if c.cols[i]%2 == 0 {
				s.avoidHor(other, i)
			} else {
				s.lineHor(other, i)
			}
		}
		if c.rows[i]+c.rowsAvoid[i] == c.target {
			changed = true
			other = getEmptyCrossingInRow(s, i)
			if c.rows[i]%2 == 0 {
				s.avoidVer(i, other)
			} else {
				s.lineVer(i, other)
			}
		}
	}

	return changed
}

func getEmptyCrossingInColumn(
	s *state,
	col model.Dimension,
) model.Dimension {
	var l, a bool
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		if l, a = s.horAt(row, col); !l && !a {
			return row
		}
	}
	return 0
}

func getEmptyCrossingInRow(
	s *state,
	row model.Dimension,
) model.Dimension {
	var l, a bool
	for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {
		if l, a = s.verAt(row, col); !l && !a {
			return col
		}
	}
	return 0
}
