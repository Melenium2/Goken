package astparser

import (
	"bytes"
	"fmt"
	"github.com/goken/utils"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

type GoFileParser struct {
	astfile        *ast.File
	representation *ASTFile
}

func (p *GoFileParser) Parse(src []byte) *ASTFile {
	var err error
	fset := token.NewFileSet()
	p.astfile, err = parser.ParseFile(fset, "./astparser/parser.go", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	if p.representation == nil {
		p.representation = &ASTFile{}
	}
	p.parsePackageName()
	if len(p.astfile.Comments) > 0 {
		p.parseComments()
	}
	if len(p.astfile.Decls) > 0 {
		p.parseDecl()
	}

	return p.representation
}

func (p *GoFileParser) parsePackageName() {
	p.representation.PackageName = p.astfile.Name.Name
}

func (p *GoFileParser) parseComments() {
	comments := make([]string, 0)
	for _, c := range p.astfile.Comments {
		text := c.Text()
		if text != "" {
			comments = append(comments, text)
		}
	}
	p.representation.Comments = comments
}

func (p *GoFileParser) parseDecl() {
	for _, d := range p.astfile.Decls {
		if dec, ok := d.(*ast.FuncDecl); ok {
			f := FuncType{
				Name: dec.Name.Name,
				Docs: dec.Doc.Text(),
			}
			if dec.Recv != nil {
				receiver := parseFieldListAsPairs(dec.Recv)
				if len(receiver) > 0 {
					for _, r := range receiver {
						if r.First != "" && r.Second != "" {
							f.Receiver = r
							break
						}
					}
				}
			}
			if dec.Type != nil {
				params := parseFieldListAsPairs(dec.Type.Params)
				if len(params) > 0 {
					f.Parameters = append(f.Parameters, params...)
				}
				results := parseFieldListAsPairs(dec.Type.Results)
				if len(results) > 0 {
					f.Results = append(f.Results, results...)
				}
			}
			if dec.Body != nil {
				fset := token.NewFileSet()
				bts := bytes.NewBufferString("")
				err := format.Node(bts, fset, dec.Body)
				if err != nil {
					log.Fatal(err)
				}
				btsStr := bts.String()
				f.Body = btsStr[1 : len(btsStr)-1]
			}

			p.representation.Functions = append(p.representation.Functions, f)
		}
		if dec, ok := d.(*ast.GenDecl); ok {
			switch dec.Tok {
			case token.IMPORT:
				p.representation.Imports = parseImports(dec.Specs)
			case token.CONST:
				p.representation.Constants = append(p.representation.Constants, parseConstants(dec.Specs)...)
			case token.VAR:
				p.representation.Vars = append(p.representation.Vars, parseVars(dec.Specs)...)
			case token.TYPE:
				parseType(dec.Specs, p.representation)
			default:
				log.Print("Unsupported type")
			}
		}
	}
}

func parseType(specs []ast.Spec, file *ASTFile) {
	for _, s := range specs {
		tsp, ok := s.(*ast.TypeSpec)
		if !ok {
			log.Print("Type is not typeSpec")
			continue
		}
		switch t := tsp.Type.(type) {
		case *ast.InterfaceType:
			funcs := parseFieldListAsMethods(t.Methods)
			iterf := InterfaceType{
				Name:      tsp.Name.Name,
				Functions: funcs,
			}
			file.Interfaces = append(file.Interfaces, iterf)
		case *ast.StructType:
			file.Structures = append(file.Structures, StructType{
				Name:   tsp.Name.Name,
				Fields: parseFieldListAsPairs(t.Fields),
			})
		case *ast.FuncType:
			file.AnonFunctions = FuncType{
				Name:       tsp.Name.Name,
				Parameters: parseFieldListAsPairs(t.Params),
				Results:    parseFieldListAsPairs(t.Results),
			}
		default:
			log.Print("Skip unsupported type")
		}
	}
}

func parseVars(specs []ast.Spec) []TypeTuple {
	var vars []TypeTuple
	for _, s := range specs {
		vsp, ok := s.(*ast.ValueSpec)
		if !ok {
			log.Print("Value is not value spec")
			continue
		}
		tp, ok := vsp.Type.(*ast.Ident)
		if len(vsp.Values) > 0 {
			fset := token.NewFileSet()
			bts := bytes.NewBufferString("")
			err := format.Node(bts, fset, vsp.Values[0])
			if err != nil {
				log.Fatal(err)
			}
			btsStr := bts.String()
			if !ok {
				vars = append(vars, TypeTuple{
					First:  vsp.Names[0].Name,
					Second: "",
					Value:  btsStr,
				})
				continue
			}
			vars = append(vars, TypeTuple{
				First:  tp.Name,
				Second: vsp.Names[0].Name,
				Value:  btsStr,
			})
		} else {
			if !ok {
				vars = append(vars, TypeTuple{
					First:  vsp.Names[0].Name,
				})
				continue
			} else {
				vars = append(vars, TypeTuple{
					First:  tp.Name,
					Second: vsp.Names[0].Name,
				})
			}
		}
	}

	return vars
}

func parseConstants(specs []ast.Spec) []TypeTuple {
	var constants []TypeTuple
	for _, s := range specs {
		vsp, ok := s.(*ast.ValueSpec)
		if !ok {
			log.Print("Value is not values spec")
			continue
		}
		fset := token.NewFileSet()
		bts := bytes.NewBufferString("")
		err := format.Node(bts, fset, vsp.Values[0])
		if err != nil {
			log.Fatal(err)
		}
		btsStr := bts.String()
		iden, ok := vsp.Type.(*ast.Ident)
		if !ok {
			constants = append(constants, TypeTuple{
				First:  vsp.Names[0].Name,
				Second: "",
				Value:  btsStr,
			})
			continue
		}
		constants = append(constants, TypeTuple{
			First:  iden.Name,
			Second: vsp.Names[0].Name,
			Value:  btsStr,
		})
	}

	return constants
}

func parseImports(specs []ast.Spec) []TypeTuple {
	var imports []TypeTuple
	for _, s := range specs {
		is, ok := s.(*ast.ImportSpec)
		if !ok {
			log.Print("Not import spec")
			continue
		}
		pair := TypeTuple{}
		if is.Name != nil {
			pair.First = is.Name.Name
		}
		if is.Path != nil {
			pair.Second = is.Path.Value
		}
		imports = append(imports, pair)
	}

	return imports
}

func parseFieldListAsPairs(list *ast.FieldList) []TypeTuple {
	var pairs []TypeTuple
	if list != nil {
		for i, item := range list.List {
			typ := getType(item.Type)
			names := make([]string, len(item.Names))
			for _, n := range item.Names {
				names = append(names, n.Name)
			}
			if len(names) == 0 {
				if strings.HasPrefix(typ, "[]") {
					names = append(names, utils.ToLowerFirstCamelCase(fmt.Sprintf("%s%d", typ[2:3], i)))
				} else if strings.HasPrefix(typ, "*") {
					names = append(names, utils.ToLowerFirstCamelCase(fmt.Sprintf("%s%d", typ[1:2], i)))
				} else {
					names = append(names, utils.ToLowerFirstCamelCase(fmt.Sprintf("%s%d", typ[:1], i)))
				}
			}

			for _, n := range names {
				if n != "" && typ != "" {
					pairs = append(pairs, TypeTuple{
						First:  n,
						Second: typ,
					})
				}
			}
		}
	}

	return pairs
}

func parseFieldListAsMethods(list *ast.FieldList) []FuncType {
	var funcs []FuncType
	if list != nil {
		for _, item := range list.List {
			switch t := item.Type.(type) {
			case *ast.FuncType:
				f := FuncType{
					Name: item.Names[0].Name,
				}
				f.Parameters = parseFieldListAsPairs(t.Params)
				f.Results = parseFieldListAsPairs(t.Results)
				funcs = append(funcs, f)
			default:
				log.Print("Unsupported type")
			}
		}
	}
	return funcs
}

func getType(t ast.Expr) string {
	typ := ""
	switch k := t.(type) {
	case *ast.Ident:
		typ = k.Name
	case *ast.SelectorExpr:
		selector := getType(k.X)
		typ = fmt.Sprintf("%s.%s", selector, k.Sel.Name)
	case *ast.StarExpr:
		star := getType(k.X)
		typ = fmt.Sprintf("*%s", star)
	case *ast.ArrayType:
		array := getType(k.Elt)
		typ = fmt.Sprintf("[]%s", array)
	case *ast.MapType:
		key := getType(k.Key)
		value := getType(k.Value)
		typ = fmt.Sprintf("map[%s]%s", key, value)
	case *ast.InterfaceType:
		typ = "interface{}"
	case *ast.Ellipsis:
		el := getType(k.Elt)
		typ = fmt.Sprintf("...%s", el)
	default:
		log.Print("Not supported type")
		return ""
	}

	return typ
}

func New() *GoFileParser {
	return &GoFileParser{}
}
