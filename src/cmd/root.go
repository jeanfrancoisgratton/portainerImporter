// portainerImporter
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"portainerImporter/helpers"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:     "portainerImporter",
	Short:   "Import docker images in Portainer",
	Version: helpers.White(fmt.Sprintf("1.00.00-0-%s (2024.03.29)", runtime.GOARCH)),
	//	Long: `This tools allows you to a software directory structure.
	//This follows my template and allows you with minimal effort to package your software once built`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(clCmd, cfgCmd)

}
