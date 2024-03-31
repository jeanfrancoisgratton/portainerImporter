package main

import (
	"os"
	"path"
	"portainerImporter/cmd"
)

func main() {
	// Ensure that the configuration directory exists
	os.MkdirAll(path.Join(os.Getenv("HOME"), ".config", "JFG", "portainerImporter"), os.ModePerm)

	// Now go to Cobras's command loop
	cmd.Execute()
}
