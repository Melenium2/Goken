package ken

import "fmt"

type DefaultCase struct {
	statements []State
}

func NewDefaultState(states ...State) *DefaultCase {
	return &DefaultCase{
		statements: states,
	}
}

func (s *DefaultCase) AddState(states ...State) *DefaultCase {
	return &DefaultCase{
		statements: append(s.statements, states...),
	}
}

func (s *DefaultCase) States(states ...State) *DefaultCase {
	return &DefaultCase{
		statements: states,
	}
}

func (s *DefaultCase) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := fmt.Sprintf("%sdefault:\n", indentLvl)
	for _, s := range s.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}

		stmt += gen
	}

	return stmt, nil
}