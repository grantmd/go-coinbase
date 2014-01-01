package coinbase

// Contains code for making requests
// You don't want to call these, probably. Look in methods.go for functions to call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	COINBASE_API_ENDPOINT = "https://coinbase.com/api/v1/"
)

// The client holds the necessary keys and our HTTP client for making requests
type Client struct {
	APIKey     string
	httpClient *http.Client
}

// Call an API method with auth, return the raw, unprocessed body
func (c *Client) Call(http_method string, api_method string, params map[string]interface{}) ([]byte, error) {
	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	apiURL := COINBASE_API_ENDPOINT + api_method

	var req *http.Request
	var err error
	if http_method == "POST" {
		params["api_key"] = c.APIKey
		postBody, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest("POST", apiURL, bytes.NewReader(postBody))
		if err != nil {
			return nil, err
		}
	} else if http_method == "GET" {
		apiURL = apiURL + "/?api_key=" + c.APIKey
		req, err = http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Unknown HTTP method: %s", http_method)
	}

	// Make the request
	req.Header.Set("Content-type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Make sure we close the body stream no matter what
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// TODO: Check status code

	// Return
	return body, nil
}
