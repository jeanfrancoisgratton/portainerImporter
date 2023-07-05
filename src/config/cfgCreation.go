// push2registry
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/config/cfgCreation.go
// Original time: 2023/07/05 15:03

package config

// Simple function, we create a config file
func TemplateConfigCreate(configfile string) error {
	if configfile == "" {
		configfile = PortainerHostConfigFile
	}
	var portainerhostconfig = PortainerHostConfigStruct{Token: "", Environment: "local", PortainerHost: "https://nas:19943"}
	return portainerhostconfig.Json2ConfigFile(configfile)
}
