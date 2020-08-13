package cmd

import "github.com/spf13/cobra"

var newCmd = &cobra.Command{
	Use: "gen",
	Short: "Generate files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
