// ©2023 J.F.Gratton (jean-francois@famillegratton.net)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"portainerImporter/cmd"
)

func main() {

	// We ensure that the user has a config directory
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".config", "portainerImporter"), os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//// Second, we create a configuration file if none exists
	//_, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".config", "portainerImporter", "portainerImporter.json"))
	//if err != nil {
	//	//config.TemplateConfigCreate("")
	//	config.TemplatedConfigCreate()
	//}

	cmd.Execute()
}
