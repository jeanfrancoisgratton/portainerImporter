// push2registry
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/cmd/configsCommands.go
// Original time: 2023/07/05 15:23

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"push2registry/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config management",
	Long:  `This is where you configure your endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { add | delete | list } must be passed to the db command")
	},
}

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a config file",
	Run: func(cmd *cobra.Command, args []string) {
		config.ConfigCreate()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configCreateCmd)
}
