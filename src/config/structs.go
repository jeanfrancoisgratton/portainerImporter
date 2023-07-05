// push2registry
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/structs.go
// Original time: 2023/07/05 14:23

package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

var PortainerHostConfigFile = "c.json"

type PortainerHostConfigStruct struct {
	Token         string `json:"Token, omitempty"`
	Environment   string `json:"Environment"`
	PortainerHost string `json:"PortainerHost"`
}

func (p PortainerHostConfigStruct) ConfigFile2Json() (PortainerHostConfigStruct, error) {
	var payload PortainerHostConfigStruct
	var err error

	if !strings.HasSuffix(PortainerHostConfigFile, ".json") {
		PortainerHostConfigFile += ".json"
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "push2registry", PortainerHostConfigFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return PortainerHostConfigStruct{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return PortainerHostConfigStruct{}, err
	} else {
		return payload, nil
	}
}

func (p PortainerHostConfigStruct) Json2ConfigFile(outputfile string) error {
	if outputfile == "" {
		outputfile = PortainerHostConfigFile
	}
	jStream, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "push2registry", outputfile)
	err = os.WriteFile(rcFile, jStream, 0600)

	return err
}
