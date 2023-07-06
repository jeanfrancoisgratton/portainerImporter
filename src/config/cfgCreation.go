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
)

// Simple function, we create a config file
func ConfigCreate(configfile string, isTemplate bool) error {
	var portainerhostconfig = PortainerHostConfigStruct{Token: "", Environment: "local", PortainerHost: "https://localhost:19943"}
	if configfile == "" {
		configfile = PortainerHostConfigFile
	}

	if !isTemplate {
		portainerhostconfig = queryConfigValues()
	}
	return portainerhostconfig.Json2ConfigFile(configfile)
}

func queryConfigValues() PortainerHostConfigStruct {
	var env PortainerHostConfigStruct
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter the auth token (ENTER to leave empty): ")
	inputScanner.Scan()
	env.Token = inputScanner.Text()

	fmt.Printf("Please enter the selected environment (ENTER to select 'local'): ")
	inputScanner.Scan()
	env.Environment = inputScanner.Text()

	fmt.Printf("Please enter the Portainer host's URL: ")
	inputScanner.Scan()
	env.PortainerHost = inputScanner.Text()

	fmt.Printf("Please enter the Portainer username: ")
	inputScanner.Scan()
	env.Username = inputScanner.Text()

	env.Password = helpers.Encrypt(helpers.GetPassword("Please enter the Portainer user's password: "))
	return env
}
