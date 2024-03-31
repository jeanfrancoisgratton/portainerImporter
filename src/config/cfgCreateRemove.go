// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/config/cfgCreateRemove.go
// Original timestamp: 2024/03/29 12:15

package config

import (
	"portainerImporter/helpers"
)

func CreateEnvFile(cfgFilename string) error {
	//var cfgFile PortainerCredsStruct
	cfgFile := promptForEnvValues()

	cfgFile.SaveCfg(cfgFilename)

	return nil
}

func promptForEnvValues() PortainerCredsStruct {
	var pcs PortainerCredsStruct

	pcs.Name = helpers.GetStringValFromPrompt("Please name this connection: ")
	pcs.PortainerHostURL = helpers.GetStringValFromPrompt("What is the Portainer host URL ? ")
	pcs.Username = helpers.GetStringValFromPrompt("What is the username ? ")
	pcs.Password = helpers.EncodeString(helpers.GetPassword("Enter that username password: "))
	pcs.PortainerEnvironment = helpers.GetStringValFromPrompt("What is the Portainer environment ? ")
	pcs.Comments = helpers.GetStringValFromPrompt("Enter optional comments, ENTER to skip: ")

	return pcs
}

func RemoveCfgFile(args []string) error {
	var ce CustomErrors
}
