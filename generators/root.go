package generators

import "github.com/Melenium2/goken/ken"

type Generator interface {
	raw() (string, error)
}

func root(states ...ken.State) (string, error) {
	return ken.
		NewRoot(states...).
		Gofmt("-s").
		Goimports().
		EnableSyntaxChecking().
		Generate(0)
}
