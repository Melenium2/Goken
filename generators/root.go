package generators

import "github.com/goken/ken"

func root(states ...ken.State) (string, error) {
	return ken.
		NewRoot(states...).
		Gofmt("-s").
		Goimports().
		EnableSyntaxChecking().
		Generate(0)
}
