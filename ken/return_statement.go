package ken

import "strings"

type ReturnStatement struct {
	items []string
}

func NewReturnStatement(items ...string) *ReturnStatement {
	return &ReturnStatement{
		items: items,
	}
}

func (r *ReturnStatement) AddReturnItems(items ...string) *ReturnStatement {
	return &ReturnStatement{
		items: append(r.items, items...),
	}
}

func (r *ReturnStatement) ReturnItems(items ...string) *ReturnStatement {
	return &ReturnStatement{
		items: items,
	}
}

func (r *ReturnStatement) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)

	stmt := indentLvl + "return"
	if it := strings.Join(r.items, ", "); it != "" {
		stmt += " " + it
	}
	stmt += "\n"

	return stmt, nil
}
