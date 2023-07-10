// push2registry
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/config/cfgCreation.go
// Original time: 2023/07/05 15:03

package config

import (
	"bufio"
	"fmt"
	"os"
	"push2registry/helpers"
	"strings"
)

// Simple function, we create a templated config file
func TemplatedConfigCreate() error {
	var portainerhostconfig = PortainerHostConfigStruct{Token: "", Environment: "local", PortainerHost: "https://localhost:19943"}
	return portainerhostconfig.Json2ConfigFile(PortainerHostConfigFile)
}

func queryConfigValues() PortainerHostConfigStruct {
	var env PortainerHostConfigStruct
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter the auth token (ENTER to skip): ")
	inputScanner.Scan()
	env.Token = inputScanner.Text()

	fmt.Printf("Please enter the selected environment (ENTER to select 'local'): ")
	inputScanner.Scan()
	env.Environment = inputScanner.Text()
	if env.Environment == "" {
		env.Environment = "local"
	}

	fmt.Printf("Please enter the Portainer host's URL: ")
	inputScanner.Scan()
	env.PortainerHost = inputScanner.Text()

	fmt.Printf("Please enter the Portainer username (ENTER to skip): ")
	inputScanner.Scan()
	env.Username = inputScanner.Text()

	if env.Username != "" {
		env.Password = helpers.Encrypt(helpers.GetPassword("Please enter the Portainer user's password: "))
	}
	return env
}

func ConfigCreate() error {
	configstruct := queryConfigValues()
	if !strings.HasSuffix(PortainerHostConfigFile, ".json") {
		PortainerHostConfigFile += ".json"
	}
	return configstruct.Json2ConfigFile(PortainerHostConfigFile)
}
