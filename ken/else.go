package ken

import "fmt"

type Else struct {
	statements []State
}

func NewElse(states ...State) *Else {
	return &Else{
		statements: states,
	}
}

func (e *Else) AddStatements(states ...State) *Else {
	return &Else {
		statements: append(e.statements, states...),
	}
}

func (e *Else) Statements(states ...State) *Else {
	return &Else {
		statements: states,
	}
}

func (e *Else) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := fmt.Sprintf(" else {\n")
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
