package ken

import (
	"fmt"
	"github.com/Melenium2/goken/kenerrs"
)

type If struct {
	condition    string
	statements   []State
	elseIfBlocks []*ElseIf
	elseBlocks   *Else
	caller       string
}

func NewIf(condition string, states ...State) *If {
	return &If{
		condition:  condition,
		statements: states,
		caller:     fetchCallerLine(),
	}
}

func (i *If) AddStatements(states ...State) *If {
	return &If{
		condition:    i.condition,
		statements:   append(i.statements, states...),
		elseIfBlocks: i.elseIfBlocks,
		elseBlocks:   i.elseBlocks,
		caller:       i.caller,
	}
}

func (i *If) Statements(states ...State) *If {
	return &If{
		condition:    i.condition,
		statements:   states,
		elseIfBlocks: i.elseIfBlocks,
		elseBlocks:   i.elseBlocks,
		caller:       i.caller,
	}
}

func (i *If) AddElseIfBlocks(blocks ...*ElseIf) *If {
	return &If{
		condition:    i.condition,
		statements:   i.statements,
		elseIfBlocks: append(i.elseIfBlocks, blocks...),
		elseBlocks:   i.elseBlocks,
		caller:       i.caller,
	}
}

func (i *If) ElseIfBlocks(blocks ...*ElseIf) *If {
	return &If{
		condition:    i.condition,
		statements:   i.statements,
		elseIfBlocks: blocks,
		elseBlocks:   i.elseBlocks,
		caller:       i.caller,
	}
}

func (i *If) Else(block *Else) *If {
	return &If{
		condition:    i.condition,
		statements:   i.statements,
		elseIfBlocks: i.elseIfBlocks,
		elseBlocks:   block,
		caller:       i.caller,
	}
}

func (i *If) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextIndent := indent + 1

	if i.condition == "" {
		return "", kenerrs.IfConditionIsEmptyError(i.caller)
	}

	stmt := fmt.Sprintf("%sif %s {\n", indentLvl, i.condition)
	for _, s := range i.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}", indentLvl)

	for _, elseIfBlocks := range i.elseIfBlocks {
		if elseIfBlocks == nil {
			continue
		}

		gen, err := elseIfBlocks.Generate(indent)
		if err != nil {
			return "", err
		}

		stmt += gen
	}

	if elseBlock := i.elseBlocks; elseBlock != nil {
		gen, err := elseBlock.Generate(indent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += "\n"

	return stmt, nil
}
