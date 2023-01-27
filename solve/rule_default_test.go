package solve

import (
	"testing"

	"github.com/joshprzybyszewski/masyu/model"
	"github.com/stretchr/testify/assert"
)

func TestRuleDefault(t *testing.T) {

	const (
		right uint8 = 1 << iota
		down
		left
		up

		none uint8 = 0
		all        = right | down | left | up
	)

	testCases := []struct {
		name        string
		startLines  uint8
		startAvoids uint8
		expInvalid  bool
		expLines    uint8
		expAvoids   uint8
	}{{
		name: `none set`,
	}, {
		name:       `one line: right`,
		startLines: right,
		expLines:   right,
	}, {
		name:       `one line: down`,
		startLines: down,
		expLines:   down,
	}, {
		name:       `one line: left`,
		startLines: left,
		expLines:   left,
	}, {
		name:       `one line: up`,
		startLines: up,
		expLines:   up,
	}, {
		name:        `one avoid: right`,
		startAvoids: right,
		expAvoids:   right,
	}, {
		name:        `one avoid: down`,
		startAvoids: down,
		expAvoids:   down,
	}, {
		name:        `one avoid: left`,
		startAvoids: left,
		expAvoids:   left,
	}, {
		name:        `one avoid: up`,
		startAvoids: up,
		expAvoids:   up,
	}, {
		name:       `two lines: right & down`,
		startLines: right | down,
		expLines:   right | down,
		expAvoids:  left | up,
	}, {
		name:       `two lines: right & left`,
		startLines: right | left,
		expLines:   right | left,
		expAvoids:  down | up,
	}, {
		name:       `two lines: right & up`,
		startLines: right | up,
		expLines:   right | up,
		expAvoids:  down | left,
	}, {
		name:       `two lines: down & left`,
		startLines: down | left,
		expLines:   down | left,
		expAvoids:  up | right,
	}, {
		name:       `two lines: down & up`,
		startLines: down | up,
		expLines:   down | up,
		expAvoids:  left | right,
	}, {
		name:       `two lines: left & up`,
		startLines: left | up,
		expLines:   left | up,
		expAvoids:  down | right,
	}, {
		name:        `two Avoids: right & down`,
		startAvoids: right | down,
		expAvoids:   right | down,
	}, {
		name:        `two Avoids: right & left`,
		startAvoids: right | left,
		expAvoids:   right | left,
	}, {
		name:        `two Avoids: right & up`,
		startAvoids: right | up,
		expAvoids:   right | up,
	}, {
		name:        `two Avoids: down & left`,
		startAvoids: down | left,
		expAvoids:   down | left,
	}, {
		name:        `two Avoids: down & up`,
		startAvoids: down | up,
		expAvoids:   down | up,
	}, {
		name:        `two Avoids: left & up`,
		startAvoids: left | up,
		expAvoids:   left | up,
	}, {
		name:        `one line and two Avoids: left vs. right & down`,
		startLines:  left,
		startAvoids: right | down,
		expLines:    left | up,
		expAvoids:   right | down,
	}, {
		name:        `one line and two Avoids: up vs. right & down`,
		startLines:  up,
		startAvoids: right | down,
		expLines:    left | up,
		expAvoids:   right | down,
	}, {
		name:        `one line and two Avoids: down vs. right & left`,
		startLines:  down,
		startAvoids: right | left,
		expLines:    up | down,
		expAvoids:   right | left,
	}, {
		name:        `one line and two Avoids: up vs. right & left`,
		startLines:  up,
		startAvoids: right | left,
		expLines:    up | down,
		expAvoids:   right | left,
	}, {
		name:        `one line and two Avoids: left vs. right & up`,
		startLines:  left,
		startAvoids: right | up,
		expLines:    left | down,
		expAvoids:   right | up,
	}, {
		name:        `one line and two Avoids: down vs. right & up`,
		startLines:  down,
		startAvoids: right | up,
		expLines:    left | down,
		expAvoids:   right | up,
	}, {
		name:        `one line and two Avoids: right vs. down & left`,
		startLines:  right,
		startAvoids: down | left,
		expLines:    up | right,
		expAvoids:   down | left,
	}, {
		name:        `one line and two Avoids: up vs. down & left`,
		startLines:  up,
		startAvoids: down | left,
		expLines:    up | right,
		expAvoids:   down | left,
	}, {
		name:        `one line and two Avoids: left vs. down & up`,
		startLines:  left,
		startAvoids: down | up,
		expLines:    left | right,
		expAvoids:   down | up,
	}, {
		name:        `one line and two Avoids: right vs. down & up`,
		startLines:  right,
		startAvoids: down | up,
		expLines:    left | right,
		expAvoids:   down | up,
	}, {
		name:        `one line and two Avoids: right vs. left & up`,
		startLines:  right,
		startAvoids: left | up,
		expLines:    down | right,
		expAvoids:   left | up,
	}, {
		name:        `one line and two Avoids: down vs. left & up`,
		startLines:  down,
		startAvoids: left | up,
		expLines:    down | right,
		expAvoids:   left | up,
	}, {
		name:       `three lines: RDL`,
		startLines: right | down | left,
		expInvalid: true,
	}, {
		name:       `three lines: RDU`,
		startLines: right | down | up,
		expInvalid: true,
	}, {
		name:       `three lines: RLU`,
		startLines: right | left | up,
		expInvalid: true,
	}, {
		name:       `three lines: DLU`,
		startLines: down | left | up,
		expInvalid: true,
	}, {
		name:        `three Avoids: RDL`,
		startAvoids: right | down | left,
		expAvoids:   all,
	}, {
		name:        `three Avoids: RDU`,
		startAvoids: right | down | up,
		expAvoids:   all,
	}, {
		name:        `three Avoids: RLU`,
		startAvoids: right | left | up,
		expAvoids:   all,
	}, {
		name:        `three Avoids: DLU`,
		startAvoids: down | left | up,
		expAvoids:   all,
	}, {
		name:       `four lines: RDLU`,
		startLines: all,
		expInvalid: true,
	}, {
		name:        `four avoid: RDLU`,
		startAvoids: all,
		expAvoids:   all,
	}}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s := newState(
				6,
				nil,
			)
			c := model.Coord{
				Row: 3,
				Col: 3,
			}

			if tc.startLines&right == right {
				s.lineHor(c.Row, c.Col)
			}
			if tc.startAvoids&right == right {
				s.avoidHor(c.Row, c.Col)
			}

			if tc.startLines&down == down {
				s.lineVer(c.Row, c.Col)
			}
			if tc.startAvoids&down == down {
				s.avoidVer(c.Row, c.Col)
			}

			if tc.startLines&left == left {
				s.lineHor(c.Row, c.Col-1)
			}
			if tc.startAvoids&left == left {
				s.avoidHor(c.Row, c.Col-1)
			}

			if tc.startLines&up == up {
				s.lineVer(c.Row-1, c.Col)
			}
			if tc.startAvoids&up == up {
				s.avoidVer(c.Row-1, c.Col)
			}

			assert.Equal(t, tc.startLines|tc.startAvoids != 0, s.rules.hasPending)

			ss := settle(&s)

			if tc.expInvalid {
				assert.Equal(t, invalid, ss)
				return
			}
			assert.False(t, s.rules.hasPending)
			assert.Equal(t, validUnsolved, ss)

			l, a := s.horAt(c.Row, c.Col)
			assert.Equal(t, tc.expLines&right == right, l, `right should be a Line`)
			assert.Equal(t, tc.expAvoids&right == right, a, `right should be an Avoid`)

			l, a = s.verAt(c.Row, c.Col)
			assert.Equal(t, tc.expLines&down == down, l, `down should be a Line`)
			assert.Equal(t, tc.expAvoids&down == down, a, `down should be an Avoid`)

			l, a = s.horAt(c.Row, c.Col-1)
			assert.Equal(t, tc.expLines&left == left, l, `left should be a Line`)
			assert.Equal(t, tc.expAvoids&left == left, a, `left should be an Avoid`)

			l, a = s.verAt(c.Row-1, c.Col)
			assert.Equal(t, tc.expLines&up == up, l, `up should be a Line`)
			assert.Equal(t, tc.expAvoids&up == up, a, `up should be an Avoid`)
		})
	}
}
