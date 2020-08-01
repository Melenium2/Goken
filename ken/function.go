package ken

import "github.com/goken/kenerrs"

type Function struct {
	funcReceiver  *FuncReceiver
	funcSignature *FuncSignature
	statements    []State
	caller        string
}

func NewFunction(receiver *FuncReceiver, signature *FuncSignature, states ...State) *Function {
	return &Function{
		funcReceiver:  receiver,
		funcSignature: signature,
		statements:    states,
		caller:        fetchCallerLine(),
	}
}

func (f *Function) AddStatements(states ...State) *Function {
	return &Function{
		funcReceiver:  f.funcReceiver,
		funcSignature: f.funcSignature,
		statements:    append(f.statements, states...),
		caller:        f.caller,
	}
}

func (f *Function) Statements(states ...State) *Function {
	return &Function{
		funcReceiver:  f.funcReceiver,
		funcSignature: f.funcSignature,
		statements:    states,
		caller:        f.caller,
	}
}

func (f *Function) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)

	stmt := indentLvl + "func"
	receiver := ""
	if f.funcReceiver != nil {
		var err error
		receiver, err = f.funcReceiver.Generate(0)
		if err != nil {
			return "", err
		}
	}
	if receiver != "" {
		stmt += receiver + " "
	}

	if f.funcSignature == nil {
		return "", kenerrs.FuncSignatureIsNilError(f.caller)
	}
	sig, err := f.funcSignature.Generate(0)
	if err != nil {
		return "", err
	}
	stmt += sig + "{\n"

	nextIndent := indent + 1
	for _, s := range f.statements {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", nil
		}
		stmt += gen
	}

	stmt += indentLvl + "}\n"

	return stmt, nil
}
