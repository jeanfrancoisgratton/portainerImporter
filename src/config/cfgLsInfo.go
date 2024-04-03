// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/config/cfgLsInfo.go
// Original timestamp: 2024/03/28 21:18

package config

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"path/filepath"
	"portainerImporter/helpers"
	"strings"
)

func ListConfigs(envdir string) error {
	var err error
	var dirFH *os.File
	var finfo, fileInfos []os.FileInfo

	// list environment files
	if envdir == "" {
		envdir = filepath.Join(os.Getenv("HOME"), ".config", "JFG", "portainerImporter")
	}
	if dirFH, err = os.Open(envdir); err != nil {
		return err
	}

	if fileInfos, err = dirFH.Readdir(0); err != nil {
		return err
	}

	for _, info := range fileInfos {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			finfo = append(finfo, info)
		}
	}
	if err != nil {
		return err
	}

	fmt.Printf("Number of environment files: %s\n", helpers.Green(fmt.Sprintf("%d", len(finfo))))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Environment file", "File size", "Modification time"})

	for _, fi := range finfo {
		t.AppendRow([]interface{}{helpers.Green(fi.Name()), helpers.Green(helpers.SI(uint64(fi.Size()))),
			helpers.Green(fmt.Sprintf("%v", fi.ModTime().Format("2006/01/02 15:04:05")))})
	}
	t.SortBy([]table.SortBy{
		{Name: "Environment file", Mode: table.Asc},
		{Name: "File size", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatDefault
	t.Render()

	return nil
}

func ExplainCfgFile(cfgfiles []string) error {
	//oldEnvFile := EnvConfigFile

	fmt.Println("For obvious reasons, the password is not shown here")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Config file", "Config name", "Portainer host", "Username", "Portainer environment", "Comments"})
	for _, cfgfile := range cfgfiles {

		if c, err := LoadCfg(cfgfile); err != nil {
			//EnvConfigFile = oldEnvFile
			return err
		} else {
			t.AppendRow([]interface{}{helpers.Green(cfgfile), helpers.Green(c.Name), helpers.Green(c.PortainerHostURL),
				helpers.Green(c.Username), helpers.Green(c.PortainerEnvironment), helpers.Green(c.Comments)})
		}
	}
	t.SortBy([]table.SortBy{
		{Name: "Config file", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatDefault
	t.Render()

	return nil
}
