package cmd

import (
	"github.com/goken/generators"
	"github.com/spf13/cobra"
	"log"
)

var endpointsCmd = &cobra.Command{
	Use: "endpoints",
	Aliases: []string{"ep"},
	Short: "Generate new service with provided name",
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = ""
		} else {
			path = args[0]
		}
		if err := generators.Endpoints(path); err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	// Добавить новый флаг except
	// В котороый юзер передает аргументы названий методов которые он не хочет видеть в ендпоинтах
	newCmd.AddCommand(endpointsCmd)
}
