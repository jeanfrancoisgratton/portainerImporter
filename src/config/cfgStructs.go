// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/config/cfgStructs.go
// Original timestamp: 2024/03/28 20:42

package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"portainerImporter/helpers"
	"strings"
)

var CurrentConfigFile = "defaultCfg.json"

type PortainerCredsStruct struct {
	Name                 string `json:"name"`
	PortainerHostURL     string `json:"portainerhost"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PortainerEnvironment string `json:"portainerEnv,omitempty"`
	Comments             string `json:"comments,omitempty"`
}

// Save the type struct into a JSON
func (e PortainerCredsStruct) SaveCfg() error {
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	if !strings.HasSuffix(CurrentConfigFile, ".json") {
		CurrentConfigFile += ".json"
	}
	err = os.WriteFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "portainerImporter", CurrentConfigFile), jStream, 0600)

	return err
}

// Map a JSON file to a PortainerCredsStruct
func LoadCfg(fname string) (PortainerCredsStruct, error) {
	var payload PortainerCredsStruct
	var err error

	if !strings.HasSuffix(fname, ".json") {
		fname += ".json"
	}
	jFile, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "portainerImporter", fname))
	if err != nil {
		return PortainerCredsStruct{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return PortainerCredsStruct{}, err
	} else {
		// we need to decode the password before using it elsewhere
		payload.Password = helpers.DecodeString(payload.Password)
		return payload, nil
	}
}
