package solve

type applyFn func(*state)

func getSimpleNextPermutations(
	s *state,
) []applyFn {

	c, isHor, ok := s.getMostInterestingPath()
	if !ok {
		return nil
	}

	if isHor {
		return []applyFn{
			func(s *state) {
				s.lineHor(c.Row, c.Col)
			},
			func(s *state) {
				s.avoidHor(c.Row, c.Col)
			},
		}
	}

	return []applyFn{
		func(s *state) {
			s.lineVer(c.Row, c.Col)
		},
		func(s *state) {
			s.avoidVer(c.Row, c.Col)
		},
	}
}
