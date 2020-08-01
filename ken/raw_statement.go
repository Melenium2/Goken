package ken

type RawStatement struct {
	state string
	line  bool
}

func NewRawStatement(st string) *RawStatement {
	return &RawStatement{
		state: st,
		line:  true,
	}
}

func (rs *RawStatement) WithLine(line bool) *RawStatement {
	return &RawStatement{
		state: rs.state,
		line:  line,
	}
}

func (rs *RawStatement) Generate(indentlvl int) (string, error) {
	ind := Indent(indentlvl)

	var newline string
	{
		if rs.line {
			newline += "\n"
		}
	}

	return ind + rs.state + newline, nil
}
