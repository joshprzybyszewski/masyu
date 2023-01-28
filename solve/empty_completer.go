package solve

import (
	"context"

	"github.com/joshprzybyszewski/masyu/model"
)

const (
	maxDepth = 50
)

type emptyCompleter struct {
	emptyBlacks []model.Node
	emptyWhites []model.Node

	otherBlacks []model.Node
	otherWhites []model.Node

	state     state
	isTesting bool
}

func newEmptyCompleter(
	s *state,
) emptyCompleter {
	ec := emptyCompleter{
		emptyBlacks: make([]model.Node, 0, len(s.nodes)),
		emptyWhites: make([]model.Node, 0, len(s.nodes)),
	}

	for _, n := range s.nodes {
		if isBlackNodeClear(s, n) {
			ec.emptyBlacks = append(ec.emptyBlacks, n)
		} else if isWhiteNodeClear(s, n) {
			ec.emptyWhites = append(ec.emptyWhites, n)
		} else if n.IsBlack {
			ec.otherBlacks = append(ec.otherBlacks, n)
		} else {
			ec.otherWhites = append(ec.otherWhites, n)
		}
	}

	return ec
}

func (ec *emptyCompleter) complete(
	ctx context.Context,
	initial *state,
	send applyFn,
) {
	if len(ec.emptyBlacks) == 0 && len(ec.emptyWhites) == 0 {
		return
	}

	ec.buildClearBlackNodePermutations(
		ctx,
		initial,
		permutationsFactorySubstate{
			known: 0,
			perms: func(s *state) {
				if ec.isTesting || s.hasInvalid {
					return
				}
				send(s)
			},
		},
	)
}

func (ec *emptyCompleter) shouldContinue(
	s *state,
	apply applyFn,
) bool {
	ec.isTesting = true
	defer func() {
		ec.isTesting = false
	}()
	ec.state = *s
	apply(&ec.state)
	return ec.state.hasInvalid
}

func (ec *emptyCompleter) applyAndSend(
	s *state,
	apply applyFn,
) {
	ec.state = *s
	apply(&ec.state)
}

func (ec *emptyCompleter) buildClearBlackNodePermutations(
	ctx context.Context,
	s *state,
	cur permutationsFactorySubstate,
) {
	if ctx.Err() != nil {
		return
	}

	if !ec.shouldContinue(s, cur.perms) {
		return
	}

	if int(cur.known) >= maxDepth {
		ec.applyAndSend(s, cur.perms)
		return
	}

	if int(cur.known) >= len(ec.emptyBlacks) {
		// no more clear back nodes to fill
		ec.buildClearWhiteNodePermutations(
			ctx,
			s,
			permutationsFactorySubstate{
				known: 0,
				perms: func(s *state) {
					if s.hasInvalid {
						return
					}
					cur.perms(s)
				},
			},
		)
		return
	}
	myNode := ec.emptyBlacks[cur.known]

	rd := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			s.avoidHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row-1, myNode.Col+1)
			s.avoidVer(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row+1, myNode.Col)

			s.avoidHor(myNode.Row+1, myNode.Col-1)
			s.avoidHor(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	dl := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			s.avoidHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col-2)

			s.avoidVer(myNode.Row-1, myNode.Col-1)
			s.avoidVer(myNode.Row, myNode.Col-1)

			s.avoidVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row+1, myNode.Col)

			s.avoidHor(myNode.Row+1, myNode.Col-1)
			s.avoidHor(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	lu := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			s.avoidHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col-2)

			s.avoidVer(myNode.Row-1, myNode.Col-1)
			s.avoidVer(myNode.Row, myNode.Col-1)

			s.avoidVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row-2, myNode.Col)

			s.avoidHor(myNode.Row-1, myNode.Col-1)
			s.avoidHor(myNode.Row-1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ur := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			s.avoidHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row-1, myNode.Col+1)
			s.avoidVer(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row-2, myNode.Col)

			s.avoidHor(myNode.Row-1, myNode.Col-1)
			s.avoidHor(myNode.Row-1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ec.buildClearBlackNodePermutations(ctx, s, rd)
	ec.buildClearBlackNodePermutations(ctx, s, dl)
	ec.buildClearBlackNodePermutations(ctx, s, lu)
	ec.buildClearBlackNodePermutations(ctx, s, ur)
}

func (ec *emptyCompleter) buildClearWhiteNodePermutations(
	ctx context.Context,
	s *state,
	cur permutationsFactorySubstate,
) {
	if ctx.Err() != nil {
		return
	}

	if !ec.shouldContinue(s, cur.perms) {
		return
	}

	if len(ec.emptyBlacks)+int(cur.known) >= maxDepth {
		ec.applyAndSend(s, cur.perms)
		return
	}

	if int(cur.known) >= len(ec.emptyWhites) {
		ec.buildOtherBlackNodePermutations(
			ctx,
			s,
			permutationsFactorySubstate{
				known: 0,
				perms: func(s *state) {
					if s.hasInvalid {
						return
					}
					cur.perms(s)
				},
			},
		)
	}

	myNode := ec.emptyWhites[cur.known]

	setHorizontalNode := func(s *state) {
		s.lineHor(myNode.Row, myNode.Col-1)
		s.lineHor(myNode.Row, myNode.Col)

		s.avoidVer(myNode.Row-1, myNode.Col)
		s.avoidVer(myNode.Row, myNode.Col)
	}

	setVerticalNode := func(s *state) {
		s.avoidHor(myNode.Row, myNode.Col-1)
		s.avoidHor(myNode.Row, myNode.Col)

		s.lineVer(myNode.Row-1, myNode.Col)
		s.lineVer(myNode.Row, myNode.Col)
	}

	rd := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row, myNode.Col+1)
			s.avoidVer(myNode.Row-1, myNode.Col+1)
			s.avoidHor(myNode.Row, myNode.Col+1)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ru := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row-1, myNode.Col+1)
			s.avoidVer(myNode.Row, myNode.Col+1)
			s.avoidHor(myNode.Row, myNode.Col+1)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	dr := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row+1, myNode.Col)
			s.avoidHor(myNode.Row+1, myNode.Col-1)
			s.avoidVer(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	dl := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row+1, myNode.Col-1)
			s.avoidHor(myNode.Row+1, myNode.Col)
			s.avoidVer(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ld := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row, myNode.Col-1)
			s.avoidVer(myNode.Row-1, myNode.Col-1)
			s.avoidHor(myNode.Row, myNode.Col-2)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	lu := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row-1, myNode.Col-1)
			s.avoidVer(myNode.Row, myNode.Col-1)
			s.avoidHor(myNode.Row, myNode.Col-2)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ul := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row-1, myNode.Col-1)
			s.avoidHor(myNode.Row-1, myNode.Col)
			s.avoidVer(myNode.Row-2, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ur := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row-1, myNode.Col)
			s.avoidHor(myNode.Row-1, myNode.Col-1)
			s.avoidVer(myNode.Row-2, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}

	ec.buildClearWhiteNodePermutations(ctx, s, ru)
	ec.buildClearWhiteNodePermutations(ctx, s, rd)
	ec.buildClearWhiteNodePermutations(ctx, s, dr)
	ec.buildClearWhiteNodePermutations(ctx, s, dl)
	ec.buildClearWhiteNodePermutations(ctx, s, ld)
	ec.buildClearWhiteNodePermutations(ctx, s, lu)
	ec.buildClearWhiteNodePermutations(ctx, s, ul)
	ec.buildClearWhiteNodePermutations(ctx, s, ur)
}

func (ec *emptyCompleter) buildOtherBlackNodePermutations(
	ctx context.Context,
	s *state,
	cur permutationsFactorySubstate,
) {
	if ctx.Err() != nil {
		return
	}

	if !ec.shouldContinue(s, cur.perms) {
		return
	}

	if len(ec.emptyBlacks)+len(ec.emptyWhites)+int(cur.known) >= maxDepth {
		ec.applyAndSend(s, cur.perms)
		return
	}

	if int(cur.known) >= len(ec.otherBlacks) {
		// no more other back nodes to fill
		ec.buildOtherWhiteNodePermutations(
			ctx,
			s,
			permutationsFactorySubstate{
				known: 0,
				perms: func(s *state) {
					if s.hasInvalid {
						return
					}
					cur.perms(s)
				},
			},
		)
		return
	}
	myNode := ec.otherBlacks[cur.known]

	rd := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			s.avoidHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row-1, myNode.Col+1)
			s.avoidVer(myNode.Row, myNode.Col+1)

			s.avoidVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row+1, myNode.Col)

			s.avoidHor(myNode.Row+1, myNode.Col-1)
			s.avoidHor(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}
	ec.buildClearBlackNodePermutations(ctx, s, rd)

	if myNode.Col > 1 {
		dl := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				s.avoidHor(myNode.Row, myNode.Col)
				s.lineHor(myNode.Row, myNode.Col-1)
				s.lineHor(myNode.Row, myNode.Col-2)

				s.avoidVer(myNode.Row-1, myNode.Col-1)
				s.avoidVer(myNode.Row, myNode.Col-1)

				s.avoidVer(myNode.Row-1, myNode.Col)
				s.lineVer(myNode.Row, myNode.Col)
				s.lineVer(myNode.Row+1, myNode.Col)

				s.avoidHor(myNode.Row+1, myNode.Col-1)
				s.avoidHor(myNode.Row+1, myNode.Col)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearBlackNodePermutations(ctx, s, dl)
	}

	if myNode.Col > 1 && myNode.Row > 1 {
		lu := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				s.avoidHor(myNode.Row, myNode.Col)
				s.lineHor(myNode.Row, myNode.Col-1)
				s.lineHor(myNode.Row, myNode.Col-2)

				s.avoidVer(myNode.Row-1, myNode.Col-1)
				s.avoidVer(myNode.Row, myNode.Col-1)

				s.avoidVer(myNode.Row, myNode.Col)
				s.lineVer(myNode.Row-1, myNode.Col)
				s.lineVer(myNode.Row-2, myNode.Col)

				s.avoidHor(myNode.Row-1, myNode.Col-1)
				s.avoidHor(myNode.Row-1, myNode.Col)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearBlackNodePermutations(ctx, s, lu)
	}

	if myNode.Row > 1 {
		ur := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				s.avoidHor(myNode.Row, myNode.Col-1)
				s.lineHor(myNode.Row, myNode.Col)
				s.lineHor(myNode.Row, myNode.Col+1)

				s.avoidVer(myNode.Row-1, myNode.Col+1)
				s.avoidVer(myNode.Row, myNode.Col+1)

				s.avoidVer(myNode.Row, myNode.Col)
				s.lineVer(myNode.Row-1, myNode.Col)
				s.lineVer(myNode.Row-2, myNode.Col)

				s.avoidHor(myNode.Row-1, myNode.Col-1)
				s.avoidHor(myNode.Row-1, myNode.Col)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearBlackNodePermutations(ctx, s, ur)
	}
}

func (ec *emptyCompleter) buildOtherWhiteNodePermutations(
	ctx context.Context,
	s *state,
	cur permutationsFactorySubstate,
) {
	if ctx.Err() != nil {
		return
	}

	if !ec.shouldContinue(s, cur.perms) {
		return
	}

	if int(cur.known) >= len(ec.otherWhites) ||
		len(ec.emptyBlacks)+len(ec.emptyWhites)+len(ec.otherBlacks)+int(cur.known) >= maxDepth {
		ec.applyAndSend(s, cur.perms)
		return
	}

	myNode := ec.otherWhites[cur.known]

	setHorizontalNode := func(s *state) {
		s.lineHor(myNode.Row, myNode.Col-1)
		s.lineHor(myNode.Row, myNode.Col)

		s.avoidVer(myNode.Row-1, myNode.Col)
		s.avoidVer(myNode.Row, myNode.Col)
	}

	setVerticalNode := func(s *state) {
		s.avoidHor(myNode.Row, myNode.Col-1)
		s.avoidHor(myNode.Row, myNode.Col)

		s.lineVer(myNode.Row-1, myNode.Col)
		s.lineVer(myNode.Row, myNode.Col)
	}

	rd := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row, myNode.Col+1)
			s.avoidVer(myNode.Row-1, myNode.Col+1)
			s.avoidHor(myNode.Row, myNode.Col+1)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}
	ec.buildClearWhiteNodePermutations(ctx, s, rd)

	ru := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setHorizontalNode(s)
			s.lineVer(myNode.Row-1, myNode.Col+1)
			s.avoidVer(myNode.Row, myNode.Col+1)
			s.avoidHor(myNode.Row, myNode.Col+1)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}
	ec.buildClearWhiteNodePermutations(ctx, s, ru)

	dr := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row+1, myNode.Col)
			s.avoidHor(myNode.Row+1, myNode.Col-1)
			s.avoidVer(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}
	ec.buildClearWhiteNodePermutations(ctx, s, dr)

	dl := permutationsFactorySubstate{
		known: cur.known + 1,
		perms: func(s *state) {
			setVerticalNode(s)
			s.lineHor(myNode.Row+1, myNode.Col-1)
			s.avoidHor(myNode.Row+1, myNode.Col)
			s.avoidVer(myNode.Row+1, myNode.Col)

			if s.hasInvalid {
				return
			}
			cur.perms(s)
		},
	}
	ec.buildClearWhiteNodePermutations(ctx, s, dl)

	if myNode.Col > 1 {
		ld := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				setHorizontalNode(s)
				s.lineVer(myNode.Row, myNode.Col-1)
				s.avoidVer(myNode.Row-1, myNode.Col-1)
				s.avoidHor(myNode.Row, myNode.Col-2)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearWhiteNodePermutations(ctx, s, ld)

		lu := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				setHorizontalNode(s)
				s.lineVer(myNode.Row-1, myNode.Col-1)
				s.avoidVer(myNode.Row, myNode.Col-1)
				s.avoidHor(myNode.Row, myNode.Col-2)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearWhiteNodePermutations(ctx, s, lu)
	}

	if myNode.Row > 1 {
		ul := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				setVerticalNode(s)
				s.lineHor(myNode.Row-1, myNode.Col-1)
				s.avoidHor(myNode.Row-1, myNode.Col)
				s.avoidVer(myNode.Row-2, myNode.Col)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearWhiteNodePermutations(ctx, s, ul)

		ur := permutationsFactorySubstate{
			known: cur.known + 1,
			perms: func(s *state) {
				setVerticalNode(s)
				s.lineHor(myNode.Row-1, myNode.Col)
				s.avoidHor(myNode.Row-1, myNode.Col-1)
				s.avoidVer(myNode.Row-2, myNode.Col)

				if s.hasInvalid {
					return
				}
				cur.perms(s)
			},
		}
		ec.buildClearWhiteNodePermutations(ctx, s, ur)
	}
}
