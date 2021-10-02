package ken

import "github.com/Melenium2/goken/kenerrs"

type AnonFunc struct {
	async          bool `json:"async"`
	funcSignature  *AnonFuncSignature
	statements     []State
	funcInvocation *FuncInvocation
	caller         string
}

func NewAnonFunc(async bool, signature *AnonFuncSignature, states ...State) *AnonFunc {
	return &AnonFunc{
		async:         async,
		funcSignature: signature,
		statements:    states,
		caller:        fetchCallerLine(),
	}
}

func (f *AnonFunc) AddStatements(states ...State) *AnonFunc {
	return &AnonFunc{
		async:          f.async,
		funcSignature:  f.funcSignature,
		statements:     append(f.statements, states...),
		funcInvocation: f.funcInvocation,
		caller:         fetchCallerLine(),
	}
}

func (f *AnonFunc) Statements(states ...State) *AnonFunc {
	return &AnonFunc{
		async:          f.async,
		funcSignature:  f.funcSignature,
		statements:     states,
		funcInvocation: f.funcInvocation,
		caller:         fetchCallerLine(),
	}
}

func (f *AnonFunc) Invocation(invocation *FuncInvocation) *AnonFunc {
	return &AnonFunc{
		async:          f.async,
		funcSignature:  f.funcSignature,
		statements:     f.statements,
		funcInvocation: invocation,
		caller:         fetchCallerLine(),
	}
}

func (f *AnonFunc) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	stmt := indentLvl

	if f.async {
		stmt += "go "
	}
	stmt += "func"

	if f.funcSignature == nil {
		return "", kenerrs.AnonFuncSignatureIsNilError(f.caller)
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

	stmt += indentLvl + "}"

	if invoke := f.funcInvocation; invoke != nil {
		i, err := invoke.Generate(0)
		if err != nil {
			return "", err
		}
		stmt += i
	}

	return stmt, nil
}
