package coinbase

// Actual client functions to call

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account changes
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetAccountBalance() (Amount, error) {
	body, err := c.Get("account/balance", nil)

	if err != nil {
		return Amount{}, err
	}

	// parse into json
	var response Amount
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Amount{}, err
	}

	return response, nil
}

func (c *Client) GetAccountReceiveAddress() (AccountReceiveAddress, error) {
	body, err := c.Get("account/receive_address", nil)

	if err != nil {
		return AccountReceiveAddress{}, err
	}

	// parse into json
	var response AccountReceiveAddress
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AccountReceiveAddress{}, err
	}

	return response, nil
}

func (c *Client) GenerateAccountReceiveAddress(callbackURL string) (AccountReceiveAddress, error) {
	params := make(map[string]interface{})

	if callbackURL != "" {
		params["address"] = map[string]string{
			"callback_url": callbackURL,
		}
	}

	body, err := c.PostJSON("account/generate_receive_address", params)
	if err != nil {
		return AccountReceiveAddress{}, err
	}

	// parse into json
	var response AccountReceiveAddress
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AccountReceiveAddress{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Addresses
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetAddresses(page int, limit int, query string) (Addresses, error) {

	params := url.Values{}
	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if limit != 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	if query != "" {
		params.Set("query", query)
	}

	body, err := c.Get("addresses", params)
	if err != nil {
		return Addresses{}, err
	}

	// parse into json
	var response Addresses
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Addresses{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buttons
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buys
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBuys(quantity float32, agree_btc_amount_varies bool) (Buys, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))

	if agree_btc_amount_varies {
		params.Set("agree_btc_amount_varies", "true")
	}

	body, err := c.PostForm("buys", params)
	if err != nil {
		return Buys{}, err
	}

	// parse into json
	var response Buys
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Buys{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Contacts
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetContacts() (Contacts, error) {
	body, err := c.Get("contacts", nil)

	if err != nil {
		return Contacts{}, err
	}

	// parse into json
	var response Contacts
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Contacts{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Currencies
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetCurrencies() (Currencies, error) {
	body, err := c.Get("currencies", nil)

	if err != nil {
		return Currencies{}, err
	}

	// parse into json
	var response Currencies
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Currencies{}, err
	}

	return response, nil
}

func (c *Client) GetExchangeRates() (ExchangeRates, error) {
	body, err := c.Get("currencies/exchange_rates", nil)
	
	if err != nil {
		return ExchangeRates{}, err
	}

	// parse into json
	var response ExchangeRates
	err = json.Unmarshal(body, &response)

	if err != nil {
		return ExchangeRates{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Orders
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetOrders() (Orders, error) {
	body, err := c.Get("orders", nil)

	if err != nil {
		return Orders{}, err
	}

	// parse into json
	var response Orders
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Orders{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Prices
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetPricesBuy() (PricesBuy, error) {
	body, err := c.Get("prices/buy", nil)

	if err != nil {
		return PricesBuy{}, err
	}

	// parse into json
	var response PricesBuy
	err = json.Unmarshal(body, &response)

	if err != nil {
		return PricesBuy{}, err
	}

	return response, nil
}

func (c *Client) GetPricesSell() (PricesSell, error) {
	body, err := c.Get("prices/sell", nil)

	if err != nil {
		return PricesSell{}, err
	}

	// parse into json
	var response PricesSell
	err = json.Unmarshal(body, &response)

	if err != nil {
		return PricesSell{}, err
	}

	return response, nil
}

func (c *Client) GetSpotRate() (Amount, error) {
	body, err := c.Get("prices/spot_rate", nil)

	if err != nil {
		return Amount{}, err
	}

	// parse into json
	var response Amount
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Amount{}, err
	}

	return response, nil
}

func (c *Client) GetHistoricalPrices(page int) (string, error) {
	params := url.Values{}

	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	body, err := c.Get("prices/historical", params)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Recurring Payments
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Sells
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetSells(quantity float32) (Sells, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))

	body, err := c.PostForm("sells", params)

	if err != nil {
		return Sells{}, err
	}

	// parse into json
	var response Sells
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Sells{}, err
	}

	return response, nil
}

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

func (c *Client) GetTransfers(page int, limit int) (Transfers, error) {
	params := url.Values{}

	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if limit != 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	body, err := c.Get("transfers", params)

	if err != nil {
		return Transfers{}, err
	}

	// parse into json
	var response Transfers
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Transfers{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Users
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) GetUsers() (Users, error) {
	body, err := c.Get("users", nil)

	if err != nil {
		return Users{}, err
	}

	// parse into json
	var response Users
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Users{}, err
	}

	return response, nil
}
