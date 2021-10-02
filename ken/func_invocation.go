package ken

import (
	"github.com/Melenium2/goken/kenerrs"
	"strings"
)

type FuncInvocation struct {
	params  []string
	callers []string
}

func NewFuncInvocation(params ...string) *FuncInvocation {
	return &FuncInvocation{
		params:  params,
		callers: fetchCallerLineAsSlice(len(params)),
	}
}

func (f *FuncInvocation) AddParams(params ...string) *FuncInvocation {
	return &FuncInvocation{
		params:  append(f.params, params...),
		callers: append(f.callers, fetchCallerLineAsSlice(len(params))...),
	}
}

func (f *FuncInvocation) Params(params ...string) *FuncInvocation {
	return &FuncInvocation{
		params:  params,
		callers: fetchCallerLineAsSlice(len(params)),
	}
}

func (f *FuncInvocation) Generate(indent int) (string, error) {
	for i, p := range f.params {
		if p == "" {
			return "", kenerrs.FuncInvocationParameterIsEmptyError(f.callers[i])
		}
	}

	return "(" + strings.Join(f.params, ", ") + ")", nil
}
