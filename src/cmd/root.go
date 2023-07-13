// portainerImporter : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"portainerImporter/configs"
	"portainerImporter/helpers"
)

var version = "0.200-0 (2023.07.12)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "portainerImporter",
	Short:   "Push a binary package to a specified portainer host",
	Version: version,
	Long:    `This tools allows you to push your docker image (tarball format) to your portainer host.`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.Changelog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clCmd)
	rootCmd.PersistentFlags().StringVarP(&configs.PortainerHostConfigFile, "configs", "c", "", "Environment file")
	rootCmd.PersistentFlags().StringVarP(&configs.PortainerHost, "portainerhost", "p", "", "Portainer url")
	rootCmd.PersistentFlags().StringVarP(&configs.PortainerUsername, "user", "u", "", "Portainer user")
	rootCmd.PersistentFlags().StringVarP(&configs.PortainerEnv, "environment", "e", "", "Portainer environment")

}
