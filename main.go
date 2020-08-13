package main

import (
	"github.com/goken/cmd"
	"github.com/spf13/viper"
	"path"
)

func main() {
	defaults()

	// Проверка что бинарник лежит в го бин

	cmd.Execute()
	//p := astparser.New()
	//
	//st := ken.NewStruct("service").
	//	AddField("field", "string", `json:"field"`)
	//
	//inter := ken.NewInterface(
	//	"Service",
	//	ken.NewFuncSignature("foo").
	//		AddParams(ken.NewFuncParam("a", "string")).
	//		AddReturnTypeStatement(ken.NewFuncReturnType("string")),
	//)
	///*
	//NEW COMMENT
	//123
	// */
	//fo := ken.NewFor("i := 1; i < 50; i++", ken.NewRawStatement("println(i)"))
	//
	//fo1 := ken.NewFor("",  ken.NewRawStatement("println(\"infinity\")"))
	//
	//i := ken.NewIf("a > b").
	//	AddStatements(ken.NewRawStatement("var s string = \"privet\"\nprintln(s)")).
	//	ElseIfBlocks(ken.NewElseIf("a == c", ken.NewRawStatement("var s string = \"poka\"\nprintln(s)"))).
	//	Else(ken.NewElse(ken.NewRawStatement("var s string = \"ky\"\nprintln(s)")))
	//
	//s := ken.NewSwitch("os := runtime.GOOS; os").
	//	AddCase(ken.NewCase("\"linux\"")).
	//	Default(ken.NewDefaultState())
	//
	//f := ken.NewFunction(
	//	nil,
	//	ken.NewFuncSignature("test").
	//		AddParams(ken.NewFuncParam("t", "*testing.T")),
	//	).
	//	AddStatements(s).
	//	AddStatements(i).
	//	AddStatements(fo).
	//	AddStatements(fo1)
	//
	//imp := ken.NewPackage("main")
	//
	//
	//k := ken.NewRoot(imp, inter, st, f).
	//	Gofmt("-s").
	//	Goimports().
	//	EnableSyntaxChecking()
	//
	//code, err := k.Generate(0)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////println(code)
	//
	//file := p.Parse([]byte(code))
	//fmt.Fprintf(os.Stdout, "Anon: %v\n", file.AnonFunctions)
	//fmt.Fprintf(os.Stdout, "Structures: %v\n", file.Structures)
	//fmt.Fprintf(os.Stdout, "Interfaces: %v\n", file.Interfaces)
	//fmt.Fprintf(os.Stdout, "Vars: %v\n", file.Vars)
	//fmt.Fprintf(os.Stdout, "Consts: %v\n", file.Constants)
	//fmt.Fprintf(os.Stdout, "Imports: %v\n", file.Imports)
	//fmt.Fprintf(os.Stdout, "Functions: %v\n", file.Functions)
	//fmt.Fprintf(os.Stdout, "Comments: %v\n", file.Comments)
	//fmt.Fprintf(os.Stdout, "PackageName: %v\n", file.PackageName)
}

func defaults() {
	viper.SetDefault("ken_service_path", path.Join("%s", "%s"))
	viper.SetDefault("ken_endpoint_path", path.Join("%s", "%s"))
	viper.SetDefault("ken_http_transport_path", path.Join("%s", "transport"))
	viper.SetDefault("ken_grpc_transport_path", path.Join("%s", "transport"))

	viper.SetDefault("ken_service_name", "service.go")
	viper.SetDefault("ken_endpoints_name", "endpoints.go")
	viper.SetDefault("ken_http_transport_name", "http_transport.go")
	viper.SetDefault("ken_grpc_transport_name", "grpc_transport.go")
	viper.SetDefault("ken_grpc_client_name", "grpc_client_transport.go")
}

//https://github.com/moznion/gowrtr
