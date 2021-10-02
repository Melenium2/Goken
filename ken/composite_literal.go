package ken

import (
	"fmt"
	"github.com/Melenium2/goken/kenerrs"
	"strings"
)

type compositeLiteralField struct {
	key   string
	value State
}

type CompositeLiteral struct {
	typ     string
	fields  []*compositeLiteralField
	callers []string
}

func NewCompositeLiteral(typ string) *CompositeLiteral {
	return &CompositeLiteral{
		typ: typ,
	}
}

func (c *CompositeLiteral) AddField(key string, value State) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: value,
		}),
		callers: append(c.callers, fetchCallerLine()),
	}
}

func (c *CompositeLiteral) AddFieldStr(key, value string) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: NewRawStatement(fmt.Sprintf(`"%s"`, value)),
		}),
	}
}

func (c *CompositeLiteral) AddFieldRaw(key string, value interface{}) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: NewRawStatement(fmt.Sprintf("%v", value)),
		}),
		callers: append(c.callers, fetchCallerLine()),
	}
}

func (c *CompositeLiteral) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	nextLvl := Indent(indent + 1)

	stmt := fmt.Sprintf("%s%s{\n", indentLvl, c.typ)
	for i, field := range c.fields {
		gen, err := field.value.Generate(indent + 1)
		if err != nil {
			return "", err
		}

		gen = strings.TrimSpace(gen)
		stmt += nextLvl

		if key := field.key; key != "" {
			stmt += key + ": "
		}
		if gen == "" {
			return "", kenerrs.ValueOfCompositeLiteralIsEmptyError(c.callers[i])
		}

		stmt += fmt.Sprintf("%s,\n", gen)
	}
	stmt += fmt.Sprintf("%s}\n", indentLvl)

	return stmt, nil
}
