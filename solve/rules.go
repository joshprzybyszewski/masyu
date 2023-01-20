package solve

import "github.com/joshprzybyszewski/masyu/model"

type rules struct {
	// if "this" row/col changes, then run these other checks
	// [row][col]
	horizontals [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule
	verticals   [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule
}

func newRules(
	size model.Size,
) *rules {
	r := rules{}

	r.addDefault(size)

	return &r
}

func (r *rules) addDefault(
	size model.Size,
) {
	var pins [model.MaxPointsPerLine][model.MaxPointsPerLine]rule
	for row := model.Dimension(1); row <= model.Dimension(size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(size); col++ {
			pins[row][col] = newDefaultRule(row, col)
		}
	}

	for long := model.Dimension(1); long <= model.Dimension(size); long++ {
		for short := model.Dimension(1); short < model.Dimension(size); short++ {
			r.horizontals[long][short] = append(r.horizontals[long][short],
				&pins[long][short],
				&pins[long][short+1],
			)

			r.verticals[short][long] = append(r.verticals[short][long],
				&pins[short][long],
				&pins[short+1][long],
			)
		}
	}
}

func (r *rules) checkHorizontal(
	row, col model.Dimension,
	s *state,
) {
	for i := range r.horizontals[row][col] {
		r.horizontals[row][col][i].check(s)
	}
}

func (r *rules) checkVertical(
	row, col model.Dimension,
	s *state,
) {
	for i := range r.verticals[row][col] {
		r.verticals[row][col][i].check(s)
	}
}

func (r *rules) addBlackNode(
	row, col model.Dimension,
) {
	left := newBlackL1Rule(row, col)
	r.horizontals[row][col-1] = append(r.horizontals[row][col-1],
		&left,
	)
	right := newBlackR1Rule(row, col)
	r.horizontals[row][col] = append(r.horizontals[row][col],
		&right,
	)
	up := newBlackU1Rule(row, col)
	r.verticals[row-1][col] = append(r.verticals[row-1][col],
		&up,
	)
	down := newBlackD1Rule(row, col)
	r.verticals[row][col] = append(r.verticals[row][col],
		&down,
	)
	// TODO add advanced rules
}

func (r *rules) addWhiteNode(
	row, col model.Dimension,
) {
	left := newWhiteL1Rule(row, col)
	r.horizontals[row][col-1] = append(r.horizontals[row][col-1],
		&left,
	)
	right := newWhiteR1Rule(row, col)
	r.horizontals[row][col] = append(r.horizontals[row][col],
		&right,
	)
	up := newWhiteU1Rule(row, col)
	r.verticals[row-1][col] = append(r.verticals[row-1][col],
		&up,
	)
	down := newWhiteD1Rule(row, col)
	r.verticals[row][col] = append(r.verticals[row][col],
		&down,
	)

	ah := newAdvancedHorizontalWhite(row, col)
	if col > 1 {
		r.horizontals[row][col-2] = append(r.horizontals[row][col-2],
			&ah,
		)
	}
	r.horizontals[row][col+1] = append(r.horizontals[row][col+1],
		&ah,
	)

	av := newAdvancedVerticalWhite(row, col)
	if row > 1 {
		r.verticals[row-2][col] = append(r.verticals[row-2][col],
			&av,
		)
	}
	r.verticals[row+1][col] = append(r.verticals[row+1][col],
		&av,
	)

	ih := newInvertHorizontalWhite(row, col)
	if col > 1 {
		r.horizontals[row][col-2] = append(r.horizontals[row][col-2],
			&ih,
		)
	}
	r.horizontals[row][col+1] = append(r.horizontals[row][col+1],
		&ih,
	)

	iv := newInvertVerticalWhite(row, col)
	if row > 1 {
		r.verticals[row-2][col] = append(r.verticals[row-2][col],
			&iv,
		)
	}
	r.verticals[row+1][col] = append(r.verticals[row+1][col],
		&iv,
	)
}
