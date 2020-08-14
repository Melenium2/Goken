package generators

import (
	"fmt"
	"github.com/goken/astparser"
	"github.com/goken/fs"
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
	gen := newEndpointGenerator(astfile, findInterface(astfile))
	s, err := gen.raw()
	println(s)
	savePath := fmt.Sprintf(viper.GetString("ken_service_path"), "./"+viper.GetString("ken_package"), viper.GetString("ken_endpoints_name"))
	println(savePath)
	return fs.Get().WriteFile(savePath, s, false)
}

func findInterface(astfile *astparser.ASTFile) astparser.InterfaceType {
	var inter astparser.InterfaceType
	for _, items := range astfile.Interfaces {
		if strings.Contains(items.Name, "Service") {
			inter = items
		}
	}

	newFuncs := make([]astparser.FuncType, 0)
	for _, f := range inter.Functions {
		if unicode.IsUpper([]rune(f.Name)[0]) {
			newFuncs = append(newFuncs, f)
		}
	}
	inter.Functions = newFuncs
	return inter
}

type endpointsGenerator struct {
	serviceFile      *astparser.ASTFile
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
	r := ken.NewCompositeLiteral("&" + e.endpointsName)
	for _, f := range e.serviceInterface.Functions {
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

func (e *endpointsGenerator) declareMakeFuncs() []ken.State {
	funcs := make([]ken.State, len(e.serviceInterface.Functions))

	for i, item := range e.serviceInterface.Functions {
		f := ken.NewFunction(
			nil,
			ken.NewFuncSignature(fmt.Sprintf("Make%sEndpoint", item.Name)).
				Params(ken.NewFuncParam("s", e.serviceInterface.Name)).
				ReturnType("endpoint.Endpoint"),
		).AddStatements(
			ken.NewReturnStatement(
				ken.NewAnonFunc(
					false,
					ken.NewAnonFuncSignature().
						Params(ken.NewFuncParam("ctx", "context.Context"), ken.NewFuncParam("r", "interface{}")).
						ReturnType("interface{}", "error"),
				).AddStatements(
					ken.NewRawStatement("panic(\"Implement me!!\")"),
				),
			),
		)

		funcs[i] = f
	}

	return funcs
}

func (e *endpointsGenerator) declareServiceImpl() []ken.State {
	var inter astparser.InterfaceType
	for _, items := range e.serviceFile.Interfaces {
		if strings.Contains(items.Name, "Service") {
			inter = items
		}
	}

	funcs := make([]ken.State, len(inter.Functions))
	for i, m := range inter.Functions {
		signature := ken.NewFuncSignature(m.Name)
		for _, p := range m.Parameters {
			signature = signature.AddParams(ken.NewFuncParam(p.First, p.Second))
		}
		for _, r := range m.Results {
			signature = signature.AddReturnType(r.Second)
		}
		f := ken.NewFunction(
			ken.NewFuncReceiver("e", e.endpointsName),
			signature,
		).AddStatements(
			ken.NewRawStatement("panic(\"implement me pls...\")"),
		)

		funcs[i] = f
	}

	return funcs
}

func (e *endpointsGenerator) raw() (string, error) {
	s := e.declareStruct()
	nf := e.declareNewFunc()
	mf := e.declareMakeFuncs()
	si := e.declareServiceImpl()
	packageName := ken.NewPackage(viper.GetString("ken_package"))

	gComment := ken.NewComment(">> Things represented below is auto-generated by KEN Generator <<")
	siComment := ken.NewComment("Implementations of service for grpc client")
	sComment := ken.NewComment("Endpoints from the service interface.")
	mComment := ken.NewComment("The functions that creates endpoints.")

	states := make([]ken.State, 0)
	states = append(states, packageName)
	states = append(states, gComment)
	states = append(states, ken.NewNewLine())
	states = append(states, sComment)
	states = append(states, s)
	states = append(states, siComment)
	states = append(states, si...)
	states = append(states, mComment)
	states = append(states, mf...)
	states = append(states, ken.NewNewLine())
	states = append(states, nf)

	return root(states...)
}

func newEndpointGenerator(file *astparser.ASTFile, inter astparser.InterfaceType, ) Generator {
	name := inter.Name
	name = strings.Replace(name, "Service", "", -1)

	return &endpointsGenerator{
		serviceFile:      file,
		serviceInterface: inter,
		endpointsName:    name + "Endpoints",
	}
}
