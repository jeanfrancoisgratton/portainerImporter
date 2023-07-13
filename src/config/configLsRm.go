// portainerImporter
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// configLsRm.go, jfgratton : 2023-07-12

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListConfigs() error {
	cfgdir := filepath.Join(os.Getenv("HOME"), ".config", "portainerImporter")
	// Open the directory
	dirEntries, err := os.ReadDir(cfgdir)
	if err != nil {
		return err
	}

	fmt.Printf("Current configuration files in %s:\n", cfgdir)
	// Iterate over the directory entries
	for _, entry := range dirEntries {
		// Check if the entry is a file and has a .json extension
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			fmt.Printf("• %s\n", entry.Name())
		}
	}
	return nil
}

func RemoveConfig() error {
	if PortainerHostConfigFile == "" {
		fmt.Println("You must pass a -c flag to this command.")
		return nil
	}
	if !strings.HasSuffix(PortainerHostConfigFile, ".json") {
		PortainerHostConfigFile += ".json"
	}
	return os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "portainerImporter", PortainerHostConfigFile))
}
