// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/cmd/tokensCommands.go
// Original time: 2023/07/06 18:30

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"portainerImporter/helpers"
	"portainerImporter/token"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token management subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select one of the following verbs: {list | add}")
	},
}

//	var tknAddCmd = &cobra.Command{
//		Use:   "add",
//		Short: "Add a token to specified configs file",
//		Run: func(cmd *cobra.Command, args []string) {
//			executor.TokenAdd()
//		},
//	}
var tknGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get the auth token",
	Run: func(cmd *cobra.Command, args []string) {
		token.GetToken()
	},
}

var tknShowCmd = &cobra.Command{
	Use:   "show",
	Short: "prints the auth token",
	Run: func(cmd *cobra.Command, args []string) {
		var t = ""
		var e error
		t, e = token.TokenShow()
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Printf("Token is: %s\n", helpers.Green(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	//	tokenCmd.AddCommand(tknAddCmd)
	tokenCmd.AddCommand(tknGetCmd)
	tokenCmd.AddCommand(tknShowCmd)

	//rootCmd.PersistentFlags().StringVarP(&configs.PortainerHostConfigFile, "environment", "e", "portainerImporter.json", "Environment file.")
}
