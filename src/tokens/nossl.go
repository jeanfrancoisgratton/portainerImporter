// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// nossl.go, jfgratton : 2023-07-14

package tokens

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getAuthTokenHTTPS(username, password, portainerHost string) (string, error) {
	requestBody := fmt.Sprintf(`{"Username": "%s", "Password": "%s"}`, username, password)

	client := &http.Client{}
	resp, err := client.Post(portainerHost+"/api/auth", "application/json", bytes.NewBufferString(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to obtain authentication token: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResp struct {
		JWT string `json:"jwt"`
	}
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.JWT, nil
}

func getAuthTokenHTTP(username, password, portainerHost string) (string, error) {
	requestBody := fmt.Sprintf(`{"Username": "%s", "Password": "%s"}`, username, password)

	// Disable TLS verification for insecure connections
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	resp, err := http.Post(portainerHost+"/api/auth", "application/json", bytes.NewBufferString(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to obtain authentication token: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResp struct {
		JWT string `json:"jwt"`
	}
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.JWT, nil
}

//func noSSL() {
//	// Set the variables
//	username := os.Args[1]
//	password := os.Args[2]
//	portainerHost := os.Args[3]
//	environmentName := os.Args[4]
//	filePath := os.Args[5]
//
//	// Attempt a secure HTTPS connection first
//	token, err := getAuthTokenHTTPS(username, password, portainerHost)
//	if err != nil {
//		log.Println("Failed to establish a secure HTTPS connection. Attempting an insecure connection.")
//		token, err = getAuthTokenHTTP(username, password, portainerHost)
//		if err != nil {
//			log.Fatalf("Failed to obtain authentication token: %v", err)
//		}
//	}
//
//	// Get the environment ID
//	environmentID, err := getEnvironmentID(token, portainerHost, environmentName)
//	if err != nil {
//		log.Fatalf("Failed to get environment ID: %v", err)
//	}
//
//	// Import the Docker image
//	err = importDockerImage(token, portainerHost, environmentID, filePath)
//	if err != nil {
//		log.Fatalf("Failed to import Docker image: %v", err)
//	}
//
//	fmt.Println("Image imported successfully!")
//}

// Remaining functions...
