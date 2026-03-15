package coinbase

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

type Client struct {
	http.Client

	APIKey    string
	APISecret string
}

func (this *Client) Get(api_method string, params url.Values) ([]byte, error) {
	api_url := COINBASE_API_ENDPOINT + api_method

	if params != nil {
		api_url = "/?" + params.Encode()
	}

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		return nil, err
	}

	this.setAuth(api_url, req)

	req.Header.Add("Accept", "application/json")

	return this.makeRequest(req)
}

func (this *Client) GetJSON(api_method string, params url.Values) (string, error) {
	buffer, err := this.Get(api_method, params)

	if err != nil {
		return "", err
	}

	var jsonIndent bytes.Buffer
	json.Indent(&jsonIndent, buffer, "", "   ")

	return jsonIndent.String(), nil
}

func (this *Client) PostJSON(api_method string, params map[string]interface{}) ([]byte, error) {
	api_url := COINBASE_API_ENDPOINT + api_method

	if params == nil {
		params = make(map[string]interface{}, 1)
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", api_url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	this.setAuth(api_url, req)

	req.Header.Set("Content-type", "application/json")

	return this.makeRequest(req)
}

func (this *Client) PostForm(api_method string, params url.Values) ([]byte, error) {
	api_url := COINBASE_API_ENDPOINT + api_method

	if params == nil {
		params = url.Values{}
	}

	req, err := http.NewRequest("POST", api_url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	this.setAuth(api_url, req)

	req.Header.Set("Content-type", "application/json")

	return this.makeRequest(req)
}

func (this *Client) setAuth(url string, req *http.Request) {
	api_nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	api_msg := api_nonce + url
	api_sign := this.getHMAC(api_msg)

	req.Header.Add("ACCESS_KEY", this.APIKey)
	req.Header.Add("ACCESS_NONCE", api_nonce)
	req.Header.Add("ACCESS_SIGNATURE", api_sign)
}

func (this *Client) makeRequest(req *http.Request) ([]byte, error) {
	resp, err := this.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP response code: %d", resp.StatusCode)
	}

	return body, nil
}

func (this *Client) getHMAC(msg string) string {
	key_bytes := []byte(this.APISecret)
	msg_bytes := []byte(msg)

	mac := hmac.New(sha256.New, key_bytes)
	mac.Write(msg_bytes)

	return hex.EncodeToString(mac.Sum(nil))
}
