package ken

type CodeBlock struct {
	statements []State
}

func NewCodeBlock(states ...State) *CodeBlock {
	return &CodeBlock{
		statements: states,
	}
}

func (c *CodeBlock) AddStates(states ...State) *CodeBlock {
	return &CodeBlock{
		statements: append(c.statements, states...),
	}
}

func (c *CodeBlock) States(states ...State) *CodeBlock {
	return &CodeBlock{
		statements: states,
	}
}

func (c *CodeBlock) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)

	stmt := indentLvl + "{\n"
	nextLvl := indent + 1
	for _, s := range c.statements {
		gen, err := s.Generate(nextLvl)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indentLvl + "}\n"

	return stmt, nil
}
