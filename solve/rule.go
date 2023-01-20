package solve

import "github.com/joshprzybyszewski/masyu/model"

type ruleKind uint8

const (
	defaultRuleKind ruleKind = iota
	blackL1RuleKind
	blackR1RuleKind
	blackU1RuleKind
	blackD1RuleKind
	whiteL1RuleKind
	whiteR1RuleKind
	whiteU1RuleKind
	whiteD1RuleKind
)

type rule struct {
	kind ruleKind

	row model.Dimension
	col model.Dimension
}

func (r *rule) check(
	s *state,
) {
	switch r.kind {
	case defaultRuleKind:
		r.checkDefault(s)
	case blackL1RuleKind:
		r.checkBlackL1(s)
	case blackR1RuleKind:
		r.checkBlackR1(s)
	case blackU1RuleKind:
		r.checkBlackU1(s)
	case blackD1RuleKind:
		r.checkBlackD1(s)
	case whiteL1RuleKind:
		r.checkWhiteL1(s)
	case whiteR1RuleKind:
		r.checkWhiteR1(s)
	case whiteU1RuleKind:
		r.checkWhiteU1(s)
	case whiteD1RuleKind:
		r.checkWhiteD1(s)
	}
}
