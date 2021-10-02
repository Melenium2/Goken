package ken

import (
	"fmt"
	"github.com/Melenium2/goken/kenerrs"
)

type Case struct {
	condition  string
	statements []State
	caller     string
}

func NewCase(condition string, states ...State) *Case {
	return &Case{
		condition:  condition,
		statements: states,
		caller:     fetchCallerLine(),
	}
}

func (c *Case) AddSates(states ...State) *Case {
	return &Case{
		condition:  c.condition,
		statements: append(c.statements, states...),
		caller:     c.caller,
	}
}

func (c *Case) States(states ...State) *Case {
	return &Case{
		condition:  c.condition,
		statements: states,
		caller:     c.caller,
	}
}

func (c *Case) Generate(indent int) (string, error) {
	if c.condition == "" {
		return "", kenerrs.CaseConditionIsEmptyError(c.caller)
	}

	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := fmt.Sprintf("%scase %s:\n", indentLvl, c.condition)
	for _, s := range c.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	return stmt, nil
}
