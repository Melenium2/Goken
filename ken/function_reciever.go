package ken

import (
	"fmt"
	"github.com/goken/kenerrs"
)

type FuncReceiver struct {
	name   string
	typ    string
	caller string
}

func NewFuncReceiver(name, typ string) *FuncReceiver {
	return &FuncReceiver{
		name:   name,
		typ:    typ,
		caller: fetchCallerLine(),
	}
}

func (f *FuncReceiver) Generate(indent int) (string, error) {
	name := f.name
	typ := f.typ

	if name == "" && typ == "" {
		return "", nil
	}

	if name == "" {
		return "", kenerrs.FuncReceiverEmptyNameError(f.caller)
	}

	if typ == "" {
		return "", kenerrs.FuncReceiverEmptyTypeError(f.caller)
	}

	return fmt.Sprintf("(%s %s)", name, typ), nil
}
