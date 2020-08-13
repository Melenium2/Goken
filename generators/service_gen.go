package generators

import (
	"fmt"
	"github.com/goken/astparser"
	"github.com/goken/fs"
	"github.com/goken/ken"
	"github.com/goken/utils"
	"github.com/spf13/viper"
)

func Service(name, packageName string) error {
	s, err := rawService(name, packageName)
	if err != nil {
		return err
	}
	println(s)

	path := fmt.Sprintf(viper.GetString("ken_service_path"), "./"+packageName, viper.GetString("ken_service_name"))

	return fs.Get().WriteFile(
		path,
		s,
		false,
	)
}

func rawService(name, packageName string) (string, error) {
	serviceName := name + "Service"
	lowerServiceName := utils.ToLowerFirstCamelCase(serviceName)
	blankLine := ken.NewNewLine()
	pack := ken.NewPackage(packageName)

	rootComments := ken.NewComment(">> Things represented below is auto-generated by KEN Generator <<")
	comments := ken.NewComment(`Add your methods below in interface`)
	serviceInterface := ken.NewInterface(serviceName)
	commentStruct := ken.NewComment("Struct which should implements this interface")
	serviceStruct := ken.NewStruct(lowerServiceName)
	newFuncComment := ken.NewComment("New function for initialize service")
	newServiceFunc := ken.NewFunction(
		nil,
		ken.NewFuncSignature(fmt.Sprintf("New%s", serviceName)).
			ReturnType(serviceName),
		ken.NewReturnStatement(ken.NewRawStatement(fmt.Sprintf("&%s{}", lowerServiceName))),
	)

	return root(
		pack,
		rootComments,
		blankLine,
		blankLine,
		comments,
		serviceInterface,
		commentStruct,
		serviceStruct,
		newFuncComment,
		newServiceFunc,
	)
}

func parseService(path string) (*astparser.ASTFile, error) {
	content, err := fs.Get().ReadFile(path)
	if err != nil {
		return nil, err
	}

	parser := astparser.New()

	return parser.Parse([]byte(content)), nil
}