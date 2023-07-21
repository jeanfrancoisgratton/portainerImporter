package token

import (
	"encoding/json"
	"io"
	"portainerImporter/config"
	"portainerImporter/helpers"
)

func PopulateConfigStruct() (config.PortainerHostConfigStruct, error) {
	var cfg config.PortainerHostConfigStruct
	if config.PortainerHostConfigFile != "" {
		var err error
		cfg, err = config.ConfigFile2Json()
		if err != nil {
			return config.PortainerHostConfigStruct{}, err
		}
	} else {
		cfg = config.PortainerHostConfigStruct{Environment: config.PortainerEnv,
			PortainerHost: config.PortainerHost, Username: config.PortainerUsername}
	}
	if cfg.Password == "" && cfg.Username != "" {
		cfg.Password = helpers.Decrypt(helpers.GetPassword("Please enter Portainer user's password: "))
	}
	return cfg, nil
}

func decodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
