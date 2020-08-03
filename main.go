package main

import (
	"github.com/goken/ken"
	"log"
)

func main() {
	s := ken.NewSwitch("os := runtime.GOOS; os").
		AddCase(ken.NewCase("\"linux\"")).
		Default(ken.NewDefaultState())

	f := ken.NewFunction(
		nil,
		ken.NewFuncSignature("test").
			AddParams(ken.NewFuncParam("t", "*testing.T")),
		).AddStatements(s)


	k := ken.NewRoot().AddState(
		f,
	).Gofmt("-s").Goimports().EnableSyntaxChecking()

	code, err := k.Generate(0)
	if err != nil {
		log.Fatal(err)
	}

	println(code)
}

//https://github.com/moznion/gowrtr
