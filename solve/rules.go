package solve

import (
	"sort"

	"github.com/joshprzybyszewski/masyu/model"
)

type path struct {
	model.Coord
	IsHorizontal bool
}

type rules struct {
	// if "this" row/col changes, then run these other checks
	// [row][col]
	horizontals [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule
	verticals   [model.MaxPointsPerLine][model.MaxPointsPerLine][]*rule

	// unknowns describes the paths that aren't initialized known.
	// They should exist in a sorted manner, where the first one has the most
	// rules associated with it, and so on. This can be used to find "the most
	// interesting space to investigate next."
	unknowns []path
}

func newRules(
	size model.Size,
) *rules {
	return &rules{
		unknowns: make([]path, 0, 2*int(size)*int(size-1)),
	}
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

	for row := model.Dimension(1); row <= model.Dimension(size); row++ {
		for col := model.Dimension(1); col < model.Dimension(size); col++ {
			r.horizontals[row][col] = append(r.horizontals[row][col],
				&pins[row][col],
				&pins[row][col+1],
			)
		}
	}

	for col := model.Dimension(1); col <= model.Dimension(size); col++ {
		for row := model.Dimension(1); row < model.Dimension(size); row++ {
			r.verticals[row][col] = append(r.verticals[row][col],
				&pins[row][col],
				&pins[row+1][col],
			)
		}
	}
}

func (r *rules) intializeUnknowns(
	s *state,
) {

	r.addDefault(s.size)

	var l, a bool
	for row := model.Dimension(1); row <= model.Dimension(s.size); row++ {
		for col := model.Dimension(1); col <= model.Dimension(s.size); col++ {

			if l, a = s.horAt(row, col); !l && !a {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: true,
				})
			}

			if l, a = s.verAt(row, col); !l && !a {
				r.unknowns = append(r.unknowns, path{
					Coord: model.Coord{
						Row: row,
						Col: col,
					},
					IsHorizontal: false,
				})
			}
		}
	}

	var ni, nj int
	var dri, dci, drj, dcj, tmp int
	is := int(s.size)

	sort.Slice(r.unknowns, func(i, j int) bool {
		if r.unknowns[i].IsHorizontal {
			ni = len(r.horizontals[r.unknowns[i].Row][r.unknowns[i].Col])
		} else {
			ni = len(r.verticals[r.unknowns[i].Row][r.unknowns[i].Col])
		}

		if r.unknowns[j].IsHorizontal {
			nj = len(r.horizontals[r.unknowns[j].Row][r.unknowns[j].Col])
		} else {
			nj = len(r.verticals[r.unknowns[j].Row][r.unknowns[j].Col])
		}

		if ni != nj {
			return ni < nj
		}

		// There are the same number of rules for this segment.
		// Prioritize a segment that is closer to the outer wall
		dri = int(r.unknowns[i].Row) - 1
		if tmp = is - int(r.unknowns[i].Row); tmp < dri {
			dri = tmp
		}
		dci = int(r.unknowns[i].Col) - 1
		if tmp = is - int(r.unknowns[i].Col); tmp < dci {
			dci = tmp
		}

		drj = int(r.unknowns[j].Row) - 1
		if tmp = is - int(r.unknowns[j].Row); tmp < drj {
			drj = tmp
		}
		dcj = int(r.unknowns[j].Col) - 1
		if tmp = is - int(r.unknowns[j].Col); tmp < dcj {
			dcj = tmp
		}
		ni = dri + dci
		nj = drj + dcj
		if ni != nj {
			return ni < nj
		}

		// They are equally close to the outer wall.
		// Prioritize the one in the top left.
		if r.unknowns[i].Row != r.unknowns[j].Row {
			return r.unknowns[i].Row < r.unknowns[j].Row
		}
		if r.unknowns[i].Col != r.unknowns[j].Col {
			return r.unknowns[i].Col < r.unknowns[j].Col
		}

		// Check horizontal first.
		return r.unknowns[i].IsHorizontal
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

	// look at inversions for black nodes
	if ih := newInvertHorizontalBlack(row, col); ih != nil {
		r.horizontals[row][col-2] = append(r.horizontals[row][col-2],
			ih,
		)
		r.horizontals[row][col+1] = append(r.horizontals[row][col+1],
			ih,
		)
		r.verticals[row-2][col] = append(r.verticals[row-2][col],
			ih,
		)
		r.verticals[row+1][col] = append(r.verticals[row+1][col],
			ih,
		)
	}
	if iv := newInvertVerticalBlack(row, col); iv != nil {
		r.horizontals[row][col-2] = append(r.horizontals[row][col-2],
			iv,
		)
		r.horizontals[row][col+1] = append(r.horizontals[row][col+1],
			iv,
		)
		r.verticals[row-2][col] = append(r.verticals[row-2][col],
			iv,
		)
		r.verticals[row+1][col] = append(r.verticals[row+1][col],
			iv,
		)
	}
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
