package solve

import (
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

type pendingPath struct {
	model.Coord
	IsHorizontal bool
}

type rules struct {
	// if "this" row/col changes, then run these other checks
	// [row][col]
	horizontals [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule
	verticals   [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule

	pendingPath []pendingPath
}

func newRules(
	size model.Size,
) *rules {
	r := rules{
		pendingPath: make([]pendingPath, 0, 2*int(size)*int(size-1)),
	}

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

func (r *rules) initializePending(
	s *state,
) {
	var l, a bool
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {

			if l, a = s.horAt(row, col); !l && !a {
				r.pendingPath = append(r.pendingPath, pendingPath{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: true,
				})
			}

			if l, a = s.verAt(row, col); !l && !a {
				r.pendingPath = append(r.pendingPath, pendingPath{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: false,
				})
			}
		}
	}
}

func (r *rules) sortPending(
	size model.Size,
) {
	sort.Slice(r.pendingPath, func(i, j int) bool {
		var ni, nj int
		if r.pendingPath[i].IsHorizontal {
			ni = len(r.horizontals[r.pendingPath[i].Row][r.pendingPath[i].Col])
		} else {
			ni = len(r.verticals[r.pendingPath[i].Row][r.pendingPath[i].Col])
		}

		if r.pendingPath[j].IsHorizontal {
			nj = len(r.horizontals[r.pendingPath[j].Row][r.pendingPath[j].Col])
		} else {
			nj = len(r.verticals[r.pendingPath[j].Row][r.pendingPath[j].Col])
		}

		if ni != nj {
			return ni < nj
		}

		// TODO be smarter
		if r.pendingPath[i].Row != r.pendingPath[j].Row {
			return r.pendingPath[i].Row < r.pendingPath[j].Row
		}
		if r.pendingPath[i].Col != r.pendingPath[j].Col {
			return r.pendingPath[i].Col < r.pendingPath[j].Col
		}
		return r.pendingPath[i].IsHorizontal
	})
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

	// Look at extended "avoids"
	if col > 1 {
		left2 := newBlackL2Rule(row, col)
		r.horizontals[row][col-2] = append(r.horizontals[row][col-2],
			&left2,
		)
	}
	right2 := newBlackR2Rule(row, col)
	r.horizontals[row][col+1] = append(r.horizontals[row][col+1],
		&right2,
	)
	if row > 1 {
		up2 := newBlackU2Rule(row, col)
		r.verticals[row-2][col] = append(r.verticals[row-2][col],
			&up2,
		)
	}
	down2 := newBlackD2Rule(row, col)
	r.verticals[row+1][col] = append(r.verticals[row+1][col],
		&down2,
	)

	// Look at branches off the adjacencies.
	leftBranch := newBlackLBranchRule(row, col)
	r.verticals[row-1][col-1] = append(r.verticals[row-1][col-1],
		&leftBranch,
	)
	r.verticals[row][col-1] = append(r.verticals[row][col-1],
		&leftBranch,
	)
	rightBranch := newBlackRBranchRule(row, col)
	r.verticals[row-1][col+1] = append(r.verticals[row-1][col+1],
		&rightBranch,
	)
	r.verticals[row][col+1] = append(r.verticals[row][col+1],
		&rightBranch,
	)
	upBranch := newBlackUBranchRule(row, col)
	r.horizontals[row-1][col] = append(r.horizontals[row-1][col],
		&upBranch,
	)
	r.horizontals[row-1][col-1] = append(r.horizontals[row-1][col-1],
		&upBranch,
	)
	downBranch := newBlackDBranchRule(row, col)
	r.horizontals[row+1][col] = append(r.horizontals[row+1][col],
		&downBranch,
	)
	r.horizontals[row+1][col-1] = append(r.horizontals[row+1][col-1],
		&downBranch,
	)
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
