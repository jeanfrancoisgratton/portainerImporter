// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/configs/configCreate.go
// Original time: 2023/07/05 15:03

package configs

import (
	"bufio"
	"fmt"
	"os"
	"portainerImporter/helpers"
	"strings"
)

func CreateConfig() error {
	if PortainerHostConfigFile == "" {
		fmt.Println("You must pass a -c flag to this command.")
		return nil
	}

	configstruct := queryConfigValues()
	if !strings.HasSuffix(PortainerHostConfigFile, ".json") {
		PortainerHostConfigFile += ".json"
	}
	return configstruct.Json2ConfigFile(PortainerHostConfigFile)
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
