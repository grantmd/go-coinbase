package coinbase

// Actual client functions to call

import (
	"encoding/json"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account changes
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) AccountBalance() (interface{}, error) {
	body, err := c.Call("GET", "account/balance", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Amount   string
		Currency string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AccountReceiveAddress() (interface{}, error) {
	body, err := c.Call("GET", "account/receive_address", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Success     bool
		Address     string
		CallbackUrl string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Addresses
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Addresses() (interface{}, error) {
	body, err := c.Call("GET", "addresses", nil)
	if err != nil {
		return nil, err
	}

	type Address struct {
		Address struct {
			Address     string
			CallbackUrl string
			Label       string
			CreatedAt   string
		}
	}

	type Response struct {
		Addresses   []Address
		TotalCount  int
		NumPages    int
		CurrentPage int
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buttons
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buys
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Contacts
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Contacts() (interface{}, error) {
	body, err := c.Call("GET", "contacts", nil)
	if err != nil {
		return nil, err
	}

	type Contact struct {
		Email string
	}

	type Response struct {
		Contacts    []Contact
		TotalCount  int
		NumPages    int
		CurrentPage int
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Currencies
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Currencies() (interface{}, error) {
	body, err := c.Call("GET", "currencies", nil)
	if err != nil {
		return nil, err
	}

	// parse into json
	var response [][]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CurrenciesExchangeRates() (interface{}, error) {
	body, err := c.Call("GET", "currencies/exchange_rates", nil)
	if err != nil {
		return nil, err
	}

	type Rates map[string]string

	// parse into json
	var response Rates
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Orders
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Orders() (interface{}, error) {
	body, err := c.Call("GET", "orders", nil)
	if err != nil {
		return nil, err
	}

	type TotalCurrency struct {
		Cents       int
		CurrencyISO string
	}

	type Button struct {
		Type        string
		Name        string
		Description string
		ID          string
	}

	type Transaction struct {
		ID            string
		Hash          string
		Confirmations int
	}

	type Order struct {
		ID        string
		CreatedAt string
		Status    string

		TotalBTC    TotalCurrency
		TotalNative TotalCurrency

		Custom      string
		Button      Button
		Transaction Transaction
	}

	type Response struct {
		Orders      []Order
		TotalCount  int
		NumPages    int
		CurrentPage int
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Prices
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) PricesBuy() (interface{}, error) {
	body, err := c.Call("GET", "prices/buy", nil)
	if err != nil {
		return nil, err
	}

	type Amount struct {
		Amount   string
		Currency string
	}

	type Response struct {
		SubTotal Amount
		Fees     []map[string]Amount
		Total    Amount
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) PricesSell() (interface{}, error) {
	body, err := c.Call("GET", "prices/sell", nil)
	if err != nil {
		return nil, err
	}

	type Amount struct {
		Amount   string
		Currency string
	}

	type Response struct {
		SubTotal Amount
		Fees     []map[string]Amount
		Total    Amount
		Amount   string
		Currency string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SpotRate() (interface{}, error) {
	body, err := c.Call("GET", "prices/spot_rate", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Amount   string
		Currency string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Recurring Payments
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Sells
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Subscribers
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Tokens
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Transactions
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Transfers
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Users
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Users() (interface{}, error) {
	body, err := c.Call("GET", "users", nil)
	if err != nil {
		return nil, err
	}

	type Amount struct {
		Amount   string
		Currency string
	}

	type User struct {
		ID             string
		Name           string
		Email          string
		TimeZone       string
		NativeCurrency string
		Balance        Amount
		BuyLevel       int
		SellLevel      int
		BuyLimit       Amount
		SellLimit      Amount
	}

	type Response struct {
		Users []User
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
