package astparser

type ASTFile struct {
	PackageName   string
	Imports       []TypeTuple
	Comments      []string
	Interfaces    []InterfaceType
	Structures    []StructType
	Constants     []TypeTuple
	Vars          []TypeTuple
	Functions     []FuncType
	AnonFunctions FuncType
}

type TypeTuple struct {
	First  string
	Second string
	Value  string
}

type InterfaceType struct {
	Name      string
	Functions []FuncType
}

type StructType struct {
	Name   string
	Fields []TypeTuple
}

type FuncType struct {
	Name       string
	Docs       string
	Body       string
	Receiver   TypeTuple
	Parameters []TypeTuple
	Results    []TypeTuple
}
