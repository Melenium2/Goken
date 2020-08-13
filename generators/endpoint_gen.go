package generators

import (
	"fmt"
	"github.com/goken/astparser"
	"github.com/goken/ken"
	"github.com/spf13/viper"
	"strings"
	"unicode"
)

func Endpoints(path string) error {
	servicePath := path
	if servicePath == "" {
		servicePath = fmt.Sprintf(viper.GetString("ken_service_path"), "./"+viper.GetString("ken_package"), viper.GetString("ken_service_name"))
	}
	astfile, err := parseService(servicePath)
	if err != nil {
		return err
	}
	inter := findInterface(astfile)
	gen := newEndpointGenerator(inter)
	s, err := gen.raw()
	println(s)

	return nil
}

func findInterface(astfile *astparser.ASTFile) astparser.InterfaceType {
	var inter astparser.InterfaceType
	for _, items := range astfile.Interfaces {
		if strings.Contains(items.Name, "Service") {
			inter = items
		}
	}

	for i, f := range inter.Functions {
		if unicode.IsLower([]rune(f.Name)[0]) {
			inter.Functions = append(inter.Functions[:i], inter.Functions[i+1:]...)
		}
	}
	return inter
}

type endpointsGenerator struct {
	serviceInterface astparser.InterfaceType
	endpointsName    string
}

func (e *endpointsGenerator) declareStruct() *ken.Struct {
	st := ken.NewStruct(e.endpointsName)

	for _, f := range e.serviceInterface.Functions {
		st = st.AddField(f.Name+"Endpoint", "endpoint.Endpoint")
	}

	return st
}

func (e *endpointsGenerator) declareNewFunc() *ken.Function {
	r := ken.NewCompositeLiteral("&"+e.endpointsName)
	for _, f := range e.serviceInterface.Functions{
		r = r.AddFieldRaw(f.Name+"Endpoint", fmt.Sprintf("Make%sEndpoint(s)", f.Name))
	}
	f := ken.NewFunction(
		nil,
		ken.NewFuncSignature("New"+e.endpointsName).
			Params(ken.NewFuncParam("s", e.serviceInterface.Name)).
			ReturnType("*"+e.endpointsName),
	).AddStatements(
		ken.NewReturnStatement(r),
	)

	return f
}

func (e *endpointsGenerator) raw() (string, error) {
	return root(e.declareStruct(), e.declareNewFunc())
}

func newEndpointGenerator(inter astparser.InterfaceType) Generator {
	name := inter.Name
	name = strings.Replace(name, "Service", "", -1)

	return &endpointsGenerator{
		serviceInterface: inter,
		endpointsName:    name + "Endpoints",
	}
}
