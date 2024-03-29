// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envStructs.go
// Original timestamp: 2024/03/28 20:42

package env

import (
	"encoding/json"
	"os"
	"strings"
)

type PortainerCredsStruct struct {
	Name                 string `json:"name"`
	PortainerHostURL     string `json:"portainerhost"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PortainerEnvironment string `json:"portainerEnv,omitempty"`
	Comments             string `json:"comments,omitempty"`
}

func (e PortainerCredsStruct) SaveEnv(fullpath string) error {
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fullpath, jStream, 0600)

	return err
}

func LoadEnv(fname string) (PortainerCredsStruct, error) {
	var payload PortainerCredsStruct
	var err error

	if !strings.HasSuffix(fname, ".json") {
		fname += ".json"
	}
	jFile, err := os.ReadFile(fname)
	if err != nil {
		return PortainerCredsStruct{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return PortainerCredsStruct{}, err
	} else {
		return payload, nil
	}
}
