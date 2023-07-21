// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/environments/getEnv.go
// Original time: 2023/07/06 18:30

package environments

import (
	"portainerImporter/config"
	"portainerImporter/token"
)

func GetEnv() (string, error) {
	var cfg config.PortainerHostConfigStruct
	var tkn, env string
	var err error

	if cfg, err = token.PopulateConfigStruct(); err != nil {
		return "", err
	}

	tkn, err = token.GetToken()
	if err != nil {
		return "", nil
	}

	env, err = getEnvironmentID(tkn, cfg.PortainerHost, cfg.Environment)
	if err != nil {
		return "", err
	}

	return env, nil
}

func getEnvironmentID(tkn string, host string, env string) (string, error) {
	return "", nil
}
