package generators

import (
	"errors"
	"fmt"
	"github.com/goken/astparser"
	"github.com/goken/fs"
	"github.com/goken/ken"
	"github.com/goken/utils"
	"github.com/spf13/viper"
	"strings"
)

func GrpcServer(filepath, interfaceName string) error {
	println(filepath, interfaceName)
	//Проверить предали ли верный путь
	if !strings.Contains(filepath, ".pb.go") {
		return errors.New("filepath does not contain *.pb.go file")
	}
	//Есди да то проверить есть ли такой файл
	if b, _ := fs.Get().Exists(filepath); !b {
		return errors.New(fmt.Sprintf("file with path %s not exists", filepath))
	}
	//Если да то взять ast
	s, err := fs.Get().ReadFile(filepath)
	if err != nil {
		return err
	}
	parser := astparser.New()
	grpcFile := parser.Parse([]byte(s))

	//Если интерфейс нейм пустой
	intName := interfaceName
	if intName == "" {
		//То взять интрейс нейм из сервис файла
		servicePath := fmt.Sprintf(viper.GetString("ken_service_path"), "./"+viper.GetString("ken_package"), viper.GetString("ken_service_name"))
		s, err := fs.Get().ReadFile(servicePath)
		if err != nil {
			return err
		}
		file := parser.Parse([]byte(s))
		for _, item := range file.Interfaces {
			if strings.Contains(item.Name, "Service") {
				intName = item.Name
			}
		}
	}

	generator, err := newGrpcTransportGenerator(grpcFile, intName)
	if err != nil {
		return err
	}
	gen, err := generator.raw()
	println(gen)

	return nil
}

type grpcTransportGenerator struct {
	grpcFile      *astparser.ASTFile
	mainInterface astparser.InterfaceType
	interfaceName string
	name          string
}

func (g *grpcTransportGenerator) declareStruct() ken.State {
	st := ken.NewStruct(g.name)

	for _, f := range g.mainInterface.Functions {
		st = st.AddField(utils.ToLowerFirstCamelCase(f.Name), "grpc.Handler")
	}

	return st
}

func (g *grpcTransportGenerator) declareMethodsImpl() []ken.State {
	funcs := make([]ken.State, len(g.mainInterface.Functions))

	for i, item := range g.mainInterface.Functions {
		signature := ken.NewFuncSignature(item.Name)
		params := ""
		for _, p := range item.Parameters {
			signature = signature.AddParams(ken.NewFuncParam(p.First, g.addGrpcPackage(p.Second, true)))
			params += p.First + ", "
		}
		params = strings.TrimRight(params, ", ")
		for _, r := range item.Results {
			signature = signature.AddReturnType(g.addGrpcPackage(r.Second))
		}
		call := ken.NewRawStatement(fmt.Sprintf("_, resp, err := g.%s.ServeGRPC(%s)", utils.ToLowerFirstCamelCase(item.Name), params))
		condi := ken.NewIf("err != nil")
		if len(item.Parameters) > 1 {
			condi = condi.AddStatements(ken.NewReturnStatement(ken.NewRawStatement("nil"), ken.NewRawStatement("err")))
		} else {
			condi = condi.AddStatements(ken.NewReturnStatement(ken.NewRawStatement("err")))
		}

		ret := ken.NewReturnStatement()
		if len(item.Results) > 1 {
			ret = ret.AddReturnItems(ken.NewRawStatement(fmt.Sprintf("resp.(%s)", g.addGrpcPackage(item.Results[0].Second, true))), ken.NewRawStatement("nil"))
		} else {
			ret = ret.AddReturnItems(ken.NewRawStatement("nil"))
		}

		f := ken.NewFunction(
			ken.NewFuncReceiver("g", g.name),
			signature,
			call,
			condi,
			ret,
		)

		funcs[i] = f
	}

	return funcs
}

func (g *grpcTransportGenerator) declareServer() ken.State {
	composite := ken.NewCompositeLiteral("&" + g.name)

	for _, item := range g.mainInterface.Functions {
		composite = composite.AddField(
			utils.ToLowerFirstCamelCase(item.Name),
			ken.NewRawStatement(
				fmt.Sprintf("grpc.NewServer(\nep.%sEndpoint,\ngrpcDecode%sRequest,\ngrpcEncode%sResponse)\n",
					item.Name, item.Name, item.Name,
				)),
		)
	}

	ret := ken.NewReturnStatement(composite)

	return ken.NewFunction(
		nil,
		ken.NewFuncSignature("NewGrpcServer").
			AddParams(ken.NewFuncParam("ep", "*"+g.name)).
			AddReturnType(g.addGrpcPackage(g.mainInterface.Name)),
		ret,
	)
}

func (g *grpcTransportGenerator) addGrpcPackage(t string, pointer ...bool) string {
	typ := t
	if !strings.Contains(typ, "context") {
		typ = strings.TrimLeft(typ, "*")
		typ = fmt.Sprintf("%s.%s", g.grpcFile.PackageName, typ)
	}
	if len(pointer) > 0 && pointer[0] {
		typ = "*" + typ
	}

	return typ
}

func (g *grpcTransportGenerator) raw() (string, error) {
	states := make([]ken.State, 0)
	states = append(states, g.declareStruct())
	states = append(states, g.declareMethodsImpl()...)
	states = append(states, g.declareServer())

	return root(states...)
}

func newGrpcTransportGenerator(file *astparser.ASTFile, name string) (Generator, error) {
	var inter astparser.InterfaceType
	for _, item := range file.Interfaces {
		if item.Name == name+"Server" {
			inter = item
		}
	}

	if inter.Name == "" {
		return nil, errors.New(fmt.Sprintf("interface with name %s not found", name))
	}

	return &grpcTransportGenerator{
		grpcFile:      file,
		mainInterface: inter,
		interfaceName: name,
		name:          "grpcTransport",
	}, nil
}
