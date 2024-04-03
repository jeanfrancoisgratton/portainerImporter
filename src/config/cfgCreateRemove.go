// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/config/cfgCreateRemove.go
// Original timestamp: 2024/03/29 12:15

package config

import (
	"fmt"
	"github.com/jeanfrancoisgratton/customError"
	"os"
	"path/filepath"
	"portainerImporter/helpers"
	"strings"
)

// Creates a config file. If no filename is provided, it will be defaultCfg.json
// All files are hosted in $HOME/.config/JFG/portainerImporter/
func CreateEnvFile(cfgFilename string) error {
	var pcs PortainerCredsStruct

	pcs.Name = helpers.GetStringValFromPrompt("Please name this connection: ")
	pcs.PortainerHostURL = helpers.GetStringValFromPrompt("What is the Portainer host URL ? ")
	pcs.Username = helpers.GetStringValFromPrompt("What is the username ? ")
	pcs.Password = helpers.EncodeString(helpers.GetPassword("Enter that username password: "))
	pcs.PortainerEnvironment = helpers.GetStringValFromPrompt("What is the Portainer environment ? ")
	pcs.Comments = helpers.GetStringValFromPrompt("Enter optional comments, ENTER to skip: ")

	return pcs.SaveCfg(cfgFilename)
}

// Removes a config file. If no filename is provided, it will be defaultCfg.json
// All files are hosted in $HOME/.config/JFG/portainerImporter/
// If multiple filenames are provided, they will all be erased.
func RemoveCfgFile(args []string) error {
	for _, filename := range args {
		if !strings.HasSuffix(filename, ".json") {
			filename += ".json"
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "portainerImporter", filename)); err != nil {
			ce := customError.CustomError{Fatality: customError.Warning, Message: err.Error()}
			fmt.Println(ce.Error())
		} else {
			fmt.Printf("%s %s\n", filename, helpers.Green("removed"))
		}
	}
	return nil
}
