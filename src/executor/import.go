package executor

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func importDockerImage(token, portainerHost, environmentID, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, portainerHost+"/api/endpoints/"+environmentID+"/docker/images/load", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to import Docker image: %s", resp.Status)
	}

	return nil
}
