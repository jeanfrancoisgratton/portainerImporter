// ©2023 J.F.Gratton (jean-francois@famillegratton.net)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"portainerImporter/cmd"
)

func main() {

	// We ensure that the user has a configs directory
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".config", "portainerImporter"), os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	cmd.Execute()
}
