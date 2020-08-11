package ken

import "fmt"

type Import struct {
	names []string
}

func NewImport(names ...string) *Import {
	return &Import{
		names: names,
	}
}

func (i *Import) AddImports(imports ...string) *Import {
	return &Import{
		names: append(i.names, imports...),
	}
}

func (i *Import) Imports(imports ...string) *Import {
	return &Import{
		names: imports,
	}
}

func (i *Import) Generate(indent int) (string, error) {
	if len(i.names) == 0 {
		return "", nil
	}

	indentLvl := Indent(indent)
	stmt := fmt.Sprintf("%simport (\n", indentLvl)
	for _, name := range i.names {
		if name == "" {
			continue
		}
		stmt += fmt.Sprintf("%s\t\"%s\"\n", indentLvl, name)
	}
	stmt += fmt.Sprintf("%s)\n", indentLvl)

	return stmt, nil
}
