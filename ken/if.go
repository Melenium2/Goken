package ken

type If struct {
	condition    string
	statements   []State
	elseIfBlocks []*ElseIf
	elseBlocks   []*Else
	caller       string
}

func NewIf(condition string, states ...State) *If {
	return &If{
		condition: condition,
		statements: states,
	}
}