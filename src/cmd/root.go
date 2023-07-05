// push2registry : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"push2registry/helpers"
)

var version = "0.100-0 (2023.06.25)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "p2r",
	Short:   "Push a binary package to specified NxRM repo",
	Version: version,
	Long:    `This tools allows you to push your binary package to your own NxRM repo.`,
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
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clCmd)
}
