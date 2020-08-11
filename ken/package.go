package ken

import "fmt"

type Package struct {
	name string
}

func NewPackage(name string) *Package {
	return &Package{
		name: name,
	}
}

func (p *Package) Generate(indent int) (string, error) {
	indentLvl := Indent(indent)
	return fmt.Sprintf("%spackage %s\n", indentLvl, p.name), nil
}
