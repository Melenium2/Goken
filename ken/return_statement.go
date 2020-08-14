package ken

import "strings"

type ReturnStatement struct {
	items []State
}

func NewReturnStatement(items ...State) *ReturnStatement {
	return &ReturnStatement{
		items: items,
	}
}

func (r *ReturnStatement) AddReturnItems(items ...State) *ReturnStatement {
	return &ReturnStatement{
		items: append(r.items, items...),
	}
}

func (r *ReturnStatement) ReturnItems(items ...State) *ReturnStatement {
	return &ReturnStatement{
		items: items,
	}
}

func (r *ReturnStatement) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := indentLvl + "return"
	for _, s := range r.items {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += " " + gen + ", "
	}
	stmt = strings.TrimRight(stmt, ", ")

	return stmt, nil
}
