package utils

import (
	"io/ioutil"
	"lagoucoursepull/internal/model"
	"net/http"
)

// LagouHTTPClient LagouHTTPClient
type LagouHTTPClient struct {
	hc *http.Client
}

// NewLagouHTTPClient NewLagouHTTPClient
func NewLagouHTTPClient() *LagouHTTPClient {
	return &LagouHTTPClient{
		hc: &http.Client{},
	}
}

// Do Do
func (client *LagouHTTPClient) Do(request *model.LagouRequest, cookie, auth string) (body []byte, err error) {
	lagouHTTPRequest, err := NewLagouHTTPRequest(request)
	if err != nil {
		return nil, err
	}

	lagouHTTPRequest.SetHeader()
	lagouHTTPRequest.SetAuth(cookie, auth)
	resp, err := client.hc.Do(lagouHTTPRequest.r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
