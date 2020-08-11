package ken

import "fmt"

type ElseIf struct {
	condition string
	statements []State
}

func NewElseIf(condi string, states ...State) *ElseIf {
	return &ElseIf{
		condition:  condi,
		statements: states,
	}
}

func (e *ElseIf) AddStatements(states ...State) *ElseIf {
	return &ElseIf{
		condition:  e.condition,
		statements: append(e.statements, states...),
	}
}

func (e *ElseIf) Statements(states ...State) *ElseIf {
	return &ElseIf{
		condition:  e.condition,
		statements: states,
	}
}

func (e *ElseIf) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := fmt.Sprintf(" else if %s {\n", e.condition)
	for _, s := range e.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}", indentLvl)

	return stmt, nil
}
