package ken

import (
	"github.com/goken/kenerrs"
	"strings"
)

type AnonFuncSignature struct {
	funcParams []*FuncParam
	returnTyp  []string
	callers    []string
}

func NewAnonFuncSignature() *AnonFuncSignature {
	return &AnonFuncSignature{}
}

func (f *AnonFuncSignature) AddParams(params ...*FuncParam) *AnonFuncSignature {
	return &AnonFuncSignature{
		funcParams: append(f.funcParams, params...),
		returnTyp: f.returnTyp,
		callers: append(f.callers, fetchCallerLineAsSlice(len(params))...),
	}
}

func (f *AnonFuncSignature) Params(params ...*FuncParam) *AnonFuncSignature {
	return &AnonFuncSignature{
		funcParams: params,
		returnTyp: f.returnTyp,
		callers: fetchCallerLineAsSlice(len(params)),
	}
}

func (f *AnonFuncSignature) AddReturnType(types ...string) *AnonFuncSignature {
	return &AnonFuncSignature{
		funcParams: f.funcParams,
		returnTyp: append(f.returnTyp, types...),
		callers: f.callers,
	}
}

func (f *AnonFuncSignature) ReturnType(types ...string) *AnonFuncSignature {
	return &AnonFuncSignature{
		funcParams: f.funcParams,
		returnTyp: types,
		callers: f.callers,
	}
}

func (f *AnonFuncSignature) Generate(indent int) (string, error) {
	stmt := "("

	typeExisted := true
	typeMissingCaller := ""
	params := make([]string, len(f.funcParams))
	for i, p := range f.funcParams {
		if p.name == "" {
			return "", kenerrs.FuncParamNameEmptyError(f.callers[i])
		}

		paramSet := p.name
		typeExisted = p.typ != ""
		if typeExisted {
			paramSet += " " + p.typ
		}
		if !typeExisted {
			typeMissingCaller = f.callers[i]
		}
		params[i] = paramSet
	}

	if !typeExisted {
		return "", kenerrs.LastParameterTypeIsEmptyError(typeMissingCaller)
	}

	stmt += strings.Join(params, ", ") + ")"

	retType := f.returnTyp
	switch len(retType) {
	case 0:
	case 1:
		stmt += " " + retType[0]
	default:
		stmt += " (" + strings.Join(retType, ", ") + ")"
	}

	return stmt, nil
}

