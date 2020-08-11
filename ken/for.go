package ken

import "fmt"

type For struct {
	condition  string
	statements []State
}

func NewFor(condi string, states ...State) *For {
	return &For {
		condition: condi,
		statements: states,
	}
}

func (f *For) AddStatements(states ...State) *For {
	return &For{
		condition: f.condition,
		statements: append(f.statements, states...),
	}
}

func (f *For) Statements(states ...State) *For {
	return &For{
		condition: f.condition,
		statements: states,
	}
}

func (f *For) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	cond := f.condition
	stmt := fmt.Sprintf("%sfor %s", indentLvl, cond)
	if cond != "" {
		stmt += " "
	}
	stmt += "{\n"

	for _, s := range f.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indentLvl)

	return stmt, nil
}
