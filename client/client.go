package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client interface that holds http operations for communicationg with airport db
type Client interface {
	ProcessRequest(*http.Request) ([]byte, error)
}

// AirportClient type that implements Client interface
type AirportClient struct {
}

// ProcessRequest get a simple response from a REST call
func (a *AirportClient) ProcessRequest(request *http.Request) ([]byte, error) {
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return bodyBytes, err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

// GetAirPortDbSearchUri gat airport db search url, use a default one if does not exist
func GetAirPortDbSearchUri(airportDbUrl string) string {
	if len(airportDbUrl) > 0 {
		return airportDbUrl
	}
	return "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo"
}

// CreateGetRequest creates an AirportDB GET request
func CreateGetRequest(params url.Values, airportDbUrl string) (*http.Request, error) {
	uri := GetAirPortDbSearchUri(airportDbUrl)
	if len(params) > 0 {
		uri = fmt.Sprintf("%v?%v", uri, params.Encode())
	}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

// GetHttpClient create HTTP client instance for Airport DB
func GetHttpClient() *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			MaxIdleConns:          100,
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}
	return httpClient
}
