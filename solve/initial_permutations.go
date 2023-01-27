package solve

import (
	"github.com/joshprzybyszewski/masyu/model"
)

const (
	totalInitialPermutationsFactorySpace = int(1 << 15)
	permutationsFactoryMoreSpace         = totalInitialPermutationsFactorySpace - permutationsFactoryNumVals
)

type intialPermutationsFactory struct {
	permutationsFactory
	moreSpace [permutationsFactoryMoreSpace]applyFn
}

func newIntialPermutationsFactory() intialPermutationsFactory {
	return intialPermutationsFactory{}
}

func (pf *intialPermutationsFactory) save(
	a applyFn,
) {
	if pf.numVals < uint16(len(pf.vals)) {
		pf.vals[pf.numVals] = a
	} else {
		pf.moreSpace[pf.numVals-uint16(len(pf.vals))] = a
	}
	pf.numVals++
}

func (pf *intialPermutationsFactory) hasRoomForNumEmpty(
	numEmpty int,
) bool {
	numPerms := 1 << (numEmpty - 1)
	rem := len(pf.vals) + len(pf.moreSpace) - int(pf.numVals)
	return numPerms <= rem
}

func (pf *intialPermutationsFactory) chooseStart(
	byNumEmpty [40]model.Dimension,
) (int, model.Dimension) {
	for numEmpty := len(byNumEmpty); numEmpty > 1; numEmpty-- {
		if byNumEmpty[numEmpty] > 0 && pf.hasRoomForNumEmpty(numEmpty) {
			return numEmpty, byNumEmpty[numEmpty]
		}
	}

	return 0, 0
}

func (pf *intialPermutationsFactory) populateFallback(
	initial *state,
) {
	if pf.numVals > 0 {
		return
	}

	pf.buildClearBlackNodePermutations(
		initial,
		permutationsFactorySubstate{
			known:        0,
			numCrossings: 1,
			perms: func(s *state) {
			},
		},
	)
}

func (pf *intialPermutationsFactory) buildClearBlackNodePermutations(
	s *state,
	cur permutationsFactorySubstate,
) {

	myNode, ok := getNthClearBlackNode(s, cur.known)
	if !ok {
		// did not find nth
		pf.buildClearWhiteNodePermutations(
			s,
			permutationsFactorySubstate{
				known:        0,
				numCrossings: cur.numCrossings,
				perms: func(s *state) {
					cur.perms(s)
				},
			},
		)
		return
	}

	permsWithMyNode := cur.numCrossings * 4

	if permsWithMyNode > totalInitialPermutationsFactorySpace {
		pf.save(cur.perms)
		return
	}

	rd := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col+1)
			s.lineVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row+1, myNode.Col)
			cur.perms(s)
		},
	}

	dl := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col-2)
			s.lineVer(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row+1, myNode.Col)
			cur.perms(s)
		},
	}

	lu := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineHor(myNode.Row, myNode.Col-2)
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row-2, myNode.Col)
			cur.perms(s)
		},
	}

	ur := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row, myNode.Col+1)
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineVer(myNode.Row-2, myNode.Col)
			cur.perms(s)
		},
	}

	pf.buildClearBlackNodePermutations(s, rd)
	pf.buildClearBlackNodePermutations(s, dl)
	pf.buildClearBlackNodePermutations(s, lu)
	pf.buildClearBlackNodePermutations(s, ur)
}

func (pf *intialPermutationsFactory) buildClearWhiteNodePermutations(
	s *state,
	cur permutationsFactorySubstate,
) {

	myNode, ok := getNthClearWhiteNode(s, cur.known)
	if !ok {
		// did not find nth
		pf.save(cur.perms)
		return
	}

	permsWithMyNode := cur.numCrossings * 8

	if permsWithMyNode > totalInitialPermutationsFactorySpace {
		pf.save(cur.perms)
		return
	}

	rd := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row, myNode.Col+1)
			cur.perms(s)
		},
	}

	ru := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col)
			s.lineVer(myNode.Row-1, myNode.Col+1)
			cur.perms(s)
		},
	}

	dr := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineVer(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row+1, myNode.Col)
			cur.perms(s)
		},
	}

	dl := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineVer(myNode.Row, myNode.Col)
			s.lineHor(myNode.Row+1, myNode.Col-1)
			cur.perms(s)
		},
	}

	ld := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineVer(myNode.Row, myNode.Col-1)
			cur.perms(s)
		},
	}

	lu := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineHor(myNode.Row, myNode.Col-1)
			s.lineVer(myNode.Row-1, myNode.Col-1)
			cur.perms(s)
		},
	}

	ul := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineHor(myNode.Row-1, myNode.Col-1)
			cur.perms(s)
		},
	}

	ur := permutationsFactorySubstate{
		known:        cur.known + 1,
		numCrossings: permsWithMyNode,
		perms: func(s *state) {
			s.lineVer(myNode.Row-1, myNode.Col)
			s.lineHor(myNode.Row-1, myNode.Col)
			cur.perms(s)
		},
	}

	pf.buildClearBlackNodePermutations(s, ru)
	pf.buildClearBlackNodePermutations(s, rd)
	pf.buildClearBlackNodePermutations(s, dr)
	pf.buildClearBlackNodePermutations(s, dl)
	pf.buildClearBlackNodePermutations(s, ld)
	pf.buildClearBlackNodePermutations(s, lu)
	pf.buildClearBlackNodePermutations(s, ul)
	pf.buildClearBlackNodePermutations(s, ur)
}

func getNthClearBlackNode(
	s *state,
	nth model.Dimension,
) (model.Node, bool) {
	var seen model.Dimension
	for _, n := range s.nodes {
		if !isBlackNodeClear(s, n) {
			continue
		}

		if seen == nth {
			return n, true
		}
		seen++
	}
	return model.Node{}, false
}

func isBlackNodeClear(
	s *state,
	n model.Node,
) bool {

	if !n.IsBlack {
		return false
	}

	if n.Row < 2 || n.Col < 2 {
		return false
	}

	if n.Row >= model.Dimension(s.size)-1 || n.Col >= model.Dimension(s.size)-1 {
		return false
	}

	l, a := s.horAt(n.Row, n.Col-2)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col+1)
	if l || a {
		return false
	}

	l, a = s.verAt(n.Row-2, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row+1, n.Col)
	if l || a {
		return false
	}

	return true
}

func getNthClearWhiteNode(
	s *state,
	nth model.Dimension,
) (model.Node, bool) {
	var seen model.Dimension
	for _, n := range s.nodes {
		if !isWhiteNodeClear(s, n) {
			continue
		}

		if seen == nth {
			return n, true
		}
		seen++
	}
	return model.Node{}, false
}

func isWhiteNodeClear(
	s *state,
	n model.Node,
) bool {

	if n.IsBlack {
		return false
	}

	l, a := s.verAt(n.Row-1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col+1)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col+1)
	if l || a {
		return false
	}

	l, a = s.horAt(n.Row-1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row-1, n.Col)
	if l || a {
		return false
	}
	l, a = s.verAt(n.Row, n.Col)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row+1, n.Col-1)
	if l || a {
		return false
	}
	l, a = s.horAt(n.Row+1, n.Col)
	if l || a {
		return false
	}

	return true
}
