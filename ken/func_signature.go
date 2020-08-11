package ken

import (
	"github.com/goken/kenerrs"
	"strings"
)

type FuncParam struct {
	name string
	typ  string
}

func NewFuncParam(name, typ string) *FuncParam {
	return &FuncParam{
		name: name,
		typ:  typ,
	}
}

type FuncReturnType struct {
	name string
	typ  string
}

func NewFuncReturnType(typ string, name ...string) *FuncReturnType {
	s := ""
	if len(name) > 0 {
		s = name[0]
	}

	return &FuncReturnType{
		name: s,
		typ:  typ,
	}
}

func (frt *FuncReturnType) Generate(indent int) (string, error) {
	name := frt.name
	typ := frt.typ

	stmt := name
	if name != "" && typ != "" {
		stmt += " "
	}
	stmt += typ

	return stmt, nil
}

type FuncSignature struct {
	funcName          string
	funcParams        []*FuncParam
	returnType        []*FuncReturnType
	paramCallers      []string
	funcNameCaller    string
	returnTypeCallers []string
}

func NewFuncSignature(name string) *FuncSignature {
	return &FuncSignature{
		funcName:       name,
		funcNameCaller: fetchCallerLine(),
	}
}

func (fs *FuncSignature) AddParams(funcParams ...*FuncParam) *FuncSignature {
	return &FuncSignature{
		funcName:          fs.funcName,
		funcParams:        append(fs.funcParams, funcParams...),
		returnType:        fs.returnType,
		paramCallers:      append(fs.paramCallers, fetchCallerLineAsSlice(len(funcParams))...),
		funcNameCaller:    fs.funcNameCaller,
		returnTypeCallers: fs.returnTypeCallers,
	}
}

func (fs *FuncSignature) Params(params ...*FuncParam) *FuncSignature {
	return &FuncSignature{
		funcName:          fs.funcName,
		funcParams:        params,
		returnType:        fs.returnType,
		paramCallers:      append(fs.paramCallers, fetchCallerLineAsSlice(len(params))...),
		funcNameCaller:    fs.funcNameCaller,
		returnTypeCallers: fs.returnTypeCallers,
	}
}

func (fs *FuncSignature) AddReturnType(returnType ...string) *FuncSignature {
	types := make([]*FuncReturnType, len(returnType))
	for i, typ := range returnType {
		types[i] = NewFuncReturnType(typ)
	}

	return fs.AddReturnTypeStatement(types...)
}

func (fs *FuncSignature) AddReturnTypeStatement(returnTypes ...*FuncReturnType) *FuncSignature {
	return &FuncSignature{
		funcName:          fs.funcName,
		funcParams:        fs.funcParams,
		returnType:        append(fs.returnType, returnTypes...),
		paramCallers:      fs.paramCallers,
		funcNameCaller:    fs.funcNameCaller,
		returnTypeCallers: append(fs.returnTypeCallers, fetchCallerLineAsSlice(len(returnTypes))...),
	}
}

func (fs *FuncSignature) ReturnType(returnType ...string) *FuncSignature {
	types := make([]*FuncReturnType, len(returnType))
	for i, typ := range returnType {
		types[i] = NewFuncReturnType(typ)
	}

	return fs.ReturnTypeStatement(types...)
}

func (fs *FuncSignature) ReturnTypeStatement(returnTypes ...*FuncReturnType) *FuncSignature {
	return &FuncSignature{
		funcName:          fs.funcName,
		funcParams:        fs.funcParams,
		returnType:        returnTypes,
		paramCallers:      fs.paramCallers,
		funcNameCaller:    fs.funcNameCaller,
		returnTypeCallers: append(fs.returnTypeCallers, fetchCallerLineAsSlice(len(returnTypes))...),
	}
}

func (fs *FuncSignature) Generate(indent int) (string, error) {
	if fs.funcName == "" {
		return "", kenerrs.FuncSignatureEmptyFuncNameError(fs.funcNameCaller)
	}

	stmt := " " + fs.funcName
	var typeBoundaries []int
	typeExisted := true
	typeMissingCaller := ""
	for i, param := range fs.funcParams {
		if param.name == "" {
			return "", kenerrs.FuncParamNameEmptyError(fs.paramCallers[i])
		}

		typeExisted = param.typ != ""
		if typeExisted {
			typeBoundaries = append(typeBoundaries, i)
		}
		if !typeExisted {
			typeMissingCaller = fs.paramCallers[i]
		}
	}

	if !typeExisted {
		return "", kenerrs.LastParameterTypeIsEmptyError(typeMissingCaller)
	}

	stmt += "("

	groups := make([]string, len(typeBoundaries))
	prevBound := 0

	for groupIndex, bound := range typeBoundaries {
		group := fs.funcParams[prevBound : bound+1]

		chunks := make([]string, len(group))
		for i, param := range group {
			chunk := param.name
			if param.typ != "" {
				chunk += " " + param.typ
			}
			chunks[i] = chunk
		}

		groups[groupIndex] = strings.Join(chunks, ", ")

		prevBound = bound + 1
	}

	if len(groups) > 1 {
		firstIndent := Indent(indent)
		nextIndent := Indent(indent + 1)

		stmt += "\n" + firstIndent + nextIndent + strings.Join(groups, ",\n"+nextIndent) + ",\n" + firstIndent
	} else if len(groups) == 1 {
		stmt += groups[0]
	}

	stmt += ")"

	returnType := fs.returnType
	switch len(returnType) {
	case 0:
	case 1:
		typ, _ := fs.returnType[0].Generate(0)

		opener := " "
		closer := ""
		if strings.Contains(typ, " ") {
			opener = " ("
			closer = ")"
		}
		stmt += opener + typ + closer
	default:
		namedRetrnType := false
		types := make([]string, len(returnType))
		for i, ret := range returnType {
			r, _ := ret.Generate(0)
			types[i] = r

			isNamed := strings.Contains(r, " ")
			if !namedRetrnType {
				namedRetrnType = isNamed
			}
			if namedRetrnType && !isNamed {
				return "", kenerrs.UnnamedReturnTypeAppearsAfterNamedReturnTypeError(fs.returnTypeCallers[i])
			}
		}
		stmt += " (" + strings.Join(types, ", ") + ")"
	}
	return stmt, nil
}
