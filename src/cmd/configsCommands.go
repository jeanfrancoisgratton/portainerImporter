// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/cmd/configsCommands.go
// Original time: 2023/07/05 15:23

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"portainerImporter/configs"
)

var configCmd = &cobra.Command{
	Use:   "configs",
	Short: "Config management",
	Long:  `This is where you configure your endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { create | delete | list } must be passed to the db command")
	},
}

var configCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Creates a configs file",
	Run: func(cmd *cobra.Command, args []string) {
		if e := configs.CreateConfig(); e != nil {
			fmt.Println(e.Error())
		}
	},
}

var configRmCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "del", "rm"},
	Short:   "Removes a configs file as per the -c flag",
	Long:    "This means that the -c flag must be passed to this command.",
	Run: func(cmd *cobra.Command, args []string) {
		if e := configs.RemoveConfig(); e != nil {
			fmt.Println(e.Error())
		}
	},
}
var configListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all configs files in your $HOME/.configs/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		if e := configs.ListConfigs(); e != nil {
			fmt.Println(e.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configCreateCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configRmCmd)
}
