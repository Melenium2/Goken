package main

import (
	"github.com/goken/ken"
	"log"
)

func main() {
	st := ken.NewStruct("service").
		AddField("field", "string", `json:"field"`)

	inter := ken.NewInterface(
		"Service",
		ken.NewFuncSignature("foo").
			AddParams(ken.NewFuncParam("a", "string")).
			AddReturnTypeStatement(ken.NewFuncReturnType("string")),
	)

	fo := ken.NewFor("i := 1; i < 50; i++", ken.NewRawStatement("println(i)"))

	fo1 := ken.NewFor("",  ken.NewRawStatement("println(\"infinity\")"))

	i := ken.NewIf("a > b").
		AddStatements(ken.NewRawStatement("var s string = \"privet\"\nprintln(s)")).
		ElseIfBlocks(ken.NewElseIf("a == c", ken.NewRawStatement("var s string = \"poka\"\nprintln(s)"))).
		Else(ken.NewElse(ken.NewRawStatement("var s string = \"ky\"\nprintln(s)")))

	s := ken.NewSwitch("os := runtime.GOOS; os").
		AddCase(ken.NewCase("\"linux\"")).
		Default(ken.NewDefaultState())

	f := ken.NewFunction(
		nil,
		ken.NewFuncSignature("test").
			AddParams(ken.NewFuncParam("t", "*testing.T")),
		).
		AddStatements(s).
		AddStatements(i).
		AddStatements(fo).
		AddStatements(fo1)



	k := ken.NewRoot().AddState(inter).AddState(st).AddState(
		f,
	).Gofmt("-s").Goimports().EnableSyntaxChecking()

	code, err := k.Generate(0)
	if err != nil {
		log.Fatal(err)
	}

	println(code)
}

//https://github.com/moznion/gowrtr
