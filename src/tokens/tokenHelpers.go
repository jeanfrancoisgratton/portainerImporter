package tokens

import (
	"encoding/json"
	"io"
	"portainerImporter/configs"
	"portainerImporter/helpers"
)

func decodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func populateConfigStruct() (configs.PortainerHostConfigStruct, error) {
	var cfg configs.PortainerHostConfigStruct
	if configs.PortainerHostConfigFile != "" {
		var err error
		cfg, err = configs.ConfigFile2Json()
		if err != nil {
			return configs.PortainerHostConfigStruct{}, err
		}
	} else {
		cfg = configs.PortainerHostConfigStruct{Environment: configs.PortainerEnv,
			PortainerHost: configs.PortainerHost, Username: configs.PortainerUsername}
	}
	if cfg.Password == "" && cfg.Username != "" {
		cfg.Password = helpers.Decrypt(helpers.GetPassword("Please enter Portainer user's password: "))
	}
	return cfg, nil
}
