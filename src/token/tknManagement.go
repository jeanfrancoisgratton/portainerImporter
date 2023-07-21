// portainerImporter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/token/tknManagement.go
// Original time: 2023/07/13 07:39

package token

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"portainerImporter/config"
	"portainerImporter/helpers"
)

// GetToken (): fetches Portainer's auth token to-from an http post-get
// 1. We populate the HostConfig structure
// 2. We set the token
func GetToken() (string, error) {
	var cfg config.PortainerHostConfigStruct
	var err error

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if cfg, err = PopulateConfigStruct(); err != nil {
		return "", err
	}

	if cfg.Environment == "" || cfg.PortainerHost == "" {
		fmt.Println("The PortainerHost and Environment variables must not be empty")
		os.Exit(1)
	}
	cfg.Token, err = getSetToken(cfg.Username, cfg.Password, cfg.PortainerHost)
	if err != nil {
		return "", err
	}

	cfg.Json2ConfigFile(config.PortainerHostConfigFile)
	return cfg.Token, nil
}

func TokenShow() (string, error) {
	var tkn = ""
	var err error
	if tkn, err = GetToken(); err != nil {
		return "", err
	}
	return tkn, err
}

func getSetToken(uname string, passwd string, host string) (string, error) {

	requestBody := fmt.Sprintf(`{"Username": "%s", "Password": "%s"}`, uname, helpers.Decrypt(passwd))
	resp, err := http.Post(host+"/api/auth", "application/json", bytes.NewBufferString(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to obtain authentication token: %s", resp.Status)
	}

	var tokenResp struct {
		JWT string `json:"jwt"`
	}
	err = decodeJSON(resp.Body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.JWT, nil
}
