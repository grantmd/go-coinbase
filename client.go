package coinbase

// Contains code for making requests
// You don't want to call these, probably. Look in methods.go for functions to call

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	COINBASE_API_ENDPOINT = "https://api.coinbase.com/v1/"
)

// The client holds the necessary keys and our HTTP client for making requests
type Client struct {
	APIKey    string
	APISecret string

	httpClient *http.Client
}

func (c *Client) Get(api_method string, params url.Values) ([]byte, error) {
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	api_url := COINBASE_API_ENDPOINT + api_method

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		return nil, err
	}

	api_nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	api_msg := api_nonce + api_url
	api_sign := getHMAC(api_msg, c.APISecret)

	req.Header.Add("ACCESS_KEY", c.APIKey)
	req.Header.Add("ACCESS_NONCE", api_nonce)
	req.Header.Add("ACCESS_SIGNATURE", api_sign)

	req.Header.Add("Accept", "application/json")

	if params == nil {
		params = url.Values{}
	}

	// Make the request
	return c.makeRequest(req)
}

func (c *Client) PostJSON(api_method string, params map[string]interface{}) ([]byte, error) {
	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	apiURL := COINBASE_API_ENDPOINT + api_method

	if params == nil {
		params = make(map[string]interface{}, 1)
	}

	if c.APIKey != "" {
		params["api_key"] = c.APIKey
	}

	var req *http.Request
	var err error
	postBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", apiURL, bytes.NewReader(postBody))
	if err != nil {
		return nil, err
	}

	// Make the request
	req.Header.Set("Content-type", "application/json")
	return c.makeRequest(req)
}

func (c *Client) PostForm(api_method string, params url.Values) ([]byte, error) {
	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	apiURL := COINBASE_API_ENDPOINT + api_method

	if params == nil {
		params = url.Values{}
	}

	if c.APIKey != "" {
		params.Set("api_key", c.APIKey)
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	// Make the request
	return c.makeRequest(req)
}

func (c *Client) makeRequest(req *http.Request) ([]byte, error) {
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

	//fmt.Println(string(body))

	// Check status code
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return nil, fmt.Errorf("Invalid HTTP response code: %d", resp.StatusCode)
	}

	// Return
	return body, nil
}

func getHMAC(msg, key string) string {
	key_bytes := []byte(key)
	msg_bytes := []byte(msg)

	mac := hmac.New(sha256.New, key_bytes)
	mac.Write(msg_bytes)

	return hex.EncodeToString(mac.Sum(nil))
}
