package ken

import "fmt"

type Switch struct {
	condition    string
	caseStates   []*Case
	defaultState *DefaultCase
}

func NewSwitch(condition string) *Switch {
	return &Switch{
		condition: condition,
	}
}

func (s *Switch) AddCase(cases ...*Case) *Switch {
	return &Switch{
		condition: s.condition,
		caseStates: append(s.caseStates, cases...),
		defaultState: s.defaultState,
	}
}

func (s *Switch) Cases(cases ...*Case) *Switch {
	return &Switch{
		condition: s.condition,
		caseStates: cases,
		defaultState: s.defaultState,
	}
}

func (s *Switch) Default(def *DefaultCase) *Switch {
	return &Switch{
		condition: s.condition,
		caseStates: s.caseStates,
		defaultState: def,
	}
}

func (s *Switch) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)

	stmt := fmt.Sprintf("%sswitch %s {\n", indentLvl, s.condition)
	for _, s := range s.caseStates {
		if s == nil {
			continue
		}
		gen, err := s.Generate(indent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	if defaultStatement := s.defaultState; defaultStatement != nil {
		gen, err := defaultStatement.Generate(indent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}\n", indentLvl)
	return stmt, nil
}