package main

import (
	"github.com/goken/ken"
	"log"
)

func main() {
	g := ken.NewAnonFunc(
		true,
		ken.NewAnonFuncSignature().AddParams(
			ken.NewFuncParam("hi", "string")),
	).
		Invocation(ken.NewFuncInvocation("hello"))

	k := ken.NewRoot().AddState(
		g,
	).Gofmt("-s").Goimports().EnableSyntaxChecking()


	code, err := k.Generate(0)
	if err != nil {
		log.Fatal(err)
	}

	println(code)
}
