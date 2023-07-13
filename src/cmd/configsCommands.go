// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/cmd/configsCommands.go
// Original time: 2023/07/05 15:23

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"portainerImporter/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config management",
	Long:  `This is where you configure your endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { create | delete | list } must be passed to the db command")
	},
}

var configCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Creates a config file",
	Run: func(cmd *cobra.Command, args []string) {
		if e := config.CreateConfig(); e != nil {
			fmt.Println(e.Error())
		}
	},
}

var configRmCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "del", "rm"},
	Short:   "Removes a config file as per the -c flag",
	Long:    "This means that the -c flag must be passed to this command.",
	Run: func(cmd *cobra.Command, args []string) {
		if e := config.RemoveConfig(); e != nil {
			fmt.Println(e.Error())
		}
	},
}
var configListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all config files in your $HOME/.config/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		if e := config.ListConfigs(); e != nil {
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
