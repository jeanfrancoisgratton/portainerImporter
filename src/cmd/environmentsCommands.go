// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/cmd/environmentsCommands.go
// Original time: 2023/07/16 00:02

package cmd

import (
	"github.com/spf13/cobra"
	"portainerImporter/environments"
)

var envGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get the requested environment",
	Run: func(cmd *cobra.Command, args []string) {
		environments.GetEnv()
	},
}

var envShowCmd = &cobra.Command{
	Use:   "show",
	Short: "lists all environments",
	Run: func(cmd *cobra.Command, args []string) {
		environments.ListEnv()
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	//	tokenCmd.AddCommand(tknAddCmd)
	tokenCmd.AddCommand(tknGetCmd)
	tokenCmd.AddCommand(tknShowCmd)

	//rootCmd.PersistentFlags().StringVarP(&config.PortainerHostConfigFile, "environment", "e", "portainerImporter.json", "Environment file.")
}
