package ken

import (
	"fmt"
	"github.com/Melenium2/goken/kenerrs"
)

type StructField struct {
	name string
	typ  string
	tag  string
}

type Struct struct {
	name         string
	fields       []*StructField
	nameCaller   string
	fieldCallers []string
}

func NewStruct(name string) *Struct {
	return &Struct{
		name:       name,
		nameCaller: fetchCallerLine(),
	}
}

func (s *Struct) AddField(name, typ string, tag ...string) *Struct {
	t := ""
	if len(tag) > 0 {
		t = tag[0]
	}

	return &Struct{
		name: s.name,
		fields: append(s.fields, &StructField{
			name: name,
			typ:  typ,
			tag:  t,
		}),
		nameCaller:   s.nameCaller,
		fieldCallers: append(s.fieldCallers, fetchCallerLine()),
	}
}

func (s *Struct) Generate(indent int) (string, error) {
	if s.name == "" {
		return "", kenerrs.StructNameIsEmptyError(s.nameCaller)
	}

	indentLvl := Indent(indent)
	stmt := fmt.Sprintf("%stype %s struct{\n", indentLvl, s.name)

	for i, field := range s.fields {
		if field.name == "" {
			return "", kenerrs.StructFieldNameIsEmptyError(s.fieldCallers[i])
		}
		if field.typ == "" {
			return "", kenerrs.StructFieldTypeIsEmptyError(s.fieldCallers[i])
		}

		stmt += fmt.Sprintf("%s\t%s %s", indentLvl, field.name, field.typ)
		if tag := field.tag; tag != "" {
			stmt += fmt.Sprintf(" `%s`", tag)
		}
		stmt += "\n"
	}
	stmt += fmt.Sprintf("%s}\n", indentLvl)

	return stmt, nil
}
