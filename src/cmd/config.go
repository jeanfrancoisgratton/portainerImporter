// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/config.go
// Original timestamp: 2024/03/29 16:10

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"portainerImporter/config"
)

var cfgCmd = &cobra.Command{
	Use:     "cfg",
	Aliases: []string{"config"},
	Short:   "Config sub-command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Valid subcommands are: { list | add | remove }")
	},
}

//	var envListCmd = &cobra.Command{
//		Use:     "list",
//		Aliases: []string{"ls"},
//		Example: "cm env list [directory]",
//		Short:   "Lists all environment files",
//		Run: func(cmd *cobra.Command, args []string) {
//			argument := ""
//			if len(args) > 0 {
//				argument = args[0]
//			}
//			if err := environment.ListEnvironments(argument); err != nil {
//				fmt.Println(err)
//				os.Exit(2)
//			}
//		},
//	}

var cfgRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Example: "cfg remove { FILE[.json] | defaultCfg.json }",
	Short:   "Removes the configuration FILE",
	Run: func(cmd *cobra.Command, args []string) {
		fname := []string{}
		if len(args) == 0 {
			fname = []string{"defaultCfg.json"}
		} else {
			fname = args
		}
		if err := config.RemoveCfgFile(fname); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	},
}
var cfgAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Example: "cfg add [FILE[.json]]",
	Short:   "Adds the configuration FILE",
	Long: `The extension (.json) is implied and will be added if missing. Moreover, not specifying a filename
Will create a defaultCfg.json file, which is the application's default file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fname := ""
		if len(args) == 0 {
			fname = "defaultCfg.json"
		} else {
			fname = args[0]
		}
		if err := config.CreateEnvFile(fname); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	},
}

//var envInfoCmd = &cobra.Command{
//	Use:     "info",
//	Aliases: []string{"explain"},
//	Example: "cm env info FILE1[.json] FILE2[.json]... FILEn[.json]",
//	Short:   "Prints the environment FILE[12n] information",
//	Long:    `You can list as many environment files as you wish, here`,
//	Run: func(cmd *cobra.Command, args []string) {
//		envfiles := []string{"defaultEnv.json"}
//		if len(args) != 0 {
//			envfiles = args
//		}
//		if err := environment.ExplainEnvFile(envfiles); err != nil {
//			fmt.Println(err)
//			os.Exit(2)
//		}
//	},
//}

func init() {
	cfgCmd.AddCommand(cfgAddCmd)

}
