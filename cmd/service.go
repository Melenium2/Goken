package cmd

import (
	"github.com/Melenium2/goken/generators"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var packageName = "logic"

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate new service with provided name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You need to provide service name")
			return
		}
		if err := generators.Service(args[0], packageName); err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	serviceCmd.Flags().StringVarP(&packageName, "package", "p", packageName, "Provide package name for service")
	viper.BindPFlag("ken_package", serviceCmd.Flags().Lookup("package"))
	newCmd.AddCommand(serviceCmd)
}
