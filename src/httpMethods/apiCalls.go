// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/httpMethods/apiCalls.go
// Original timestamp: 2024/04/05 16:05

package httpMethods

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PortainerAPI struct {
	BaseURL string
	Client  *http.Client
}

// This is where we handle payloads for GET/POST methods,
// once the PortainerAPI struct has been set up
//type APIInterface interface {
//	Get(endpoint string, result interface{}) error
//	Post(endpoint string, data, result interface{}) error
//}

// This is the base struct, where we'll include all members
// Common across all the endpoints
type BasePayload struct {
	Username         string `json:"username"`
	ServerStatusCode int    `json:"serverstatus,omitempty"`
	JWT              string `json:"jwt,omitempty"`
}

// First (un-named, so far) endpoint
type Endpoint1 struct {
	BasePayload
	MyParam1    string `json:"myParam1"`
	MyParam2    int    `json:"myParam2,omitempty"`
	ExtraParam1 bool   `json:"ExtraParam1,omitempty"`
	ExtraParam2 string `json:"ExtraParam2,omitempty"`
}

// Second (un-named, so far) endpoint
type Endpoint2 struct {
	BasePayload
	MyParam1 string `json:"MyParam1"`
	MyParam2 bool   `json:"MyParam2,omitempty"`
}

// This is the generic GET call
func (p *PortainerAPI) Get(endpoint string, result interface{}) error {
	url := p.BaseURL + endpoint
	resp, err := p.Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

// This is the generic POST call
func (p *PortainerAPI) Post(endpoint string, data, result interface{}) error {
	url := p.BaseURL + endpoint
	requestData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := p.Client.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
