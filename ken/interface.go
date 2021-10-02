package ken

import (
	"fmt"
	"github.com/Melenium2/goken/kenerrs"
)

type Interface struct {
	name          string
	funcSignature []*FuncSignature
	caller        string
}

func NewInterface(name string, signatures ...*FuncSignature) *Interface {
	return &Interface{
		name:          name,
		funcSignature: signatures,
		caller:        fetchCallerLine(),
	}
}

func (i *Interface) AddSignatures(signatures ...*FuncSignature) *Interface {
	return &Interface{
		name:          i.name,
		funcSignature: append(i.funcSignature, signatures...),
		caller:        i.caller,
	}
}

func (i *Interface) Signatures(signatures ...*FuncSignature) *Interface {
	return &Interface{
		name:          i.name,
		funcSignature: signatures,
		caller:        i.caller,
	}
}

func (i *Interface) Generate(indent int) (string, error) {
	if i.name == "" {
		return "", kenerrs.InterfaceNameEmptyError(i.caller)
	}

	indentLvl := Indent(indent)
	nextIndent := indent + 1

	stmt := fmt.Sprintf("%stype %s interface {\n", indentLvl, i.name)
	for _, s := range i.funcSignature {
		gen, err := s.Generate(nextIndent)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indentLvl)

	return stmt, nil
}
