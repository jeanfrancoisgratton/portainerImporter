// ©2023 J.F.Gratton (jean-francois@famillegratton.net)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"push2registry/cmd"
	"push2registry/config"
)

func main() {

	// First we ensure that the user has a config directory
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".config", "push2registry"), os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// Second, we create a configuration file if none exists
	_, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".config", "push2registry", "push2registry.json"))
	if err != nil {
		config.TemplateConfigCreate("")
	}

	cmd.Execute()
}
