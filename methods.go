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

func (c *Client) AccountBalance() (Amount, error) {
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

func (c *Client) AccountReceiveAddress() (AccountReceiveAddressResponse, error) {
	body, err := c.Get("account/receive_address", nil)

	if err != nil {
		return AccountReceiveAddressResponse{}, err
	}

	// parse into json
	var response AccountReceiveAddressResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AccountReceiveAddressResponse{}, err
	}

	return response, nil
}

func (c *Client) AccountGenerateReceiveAddress(callbackURL string) (AccountGenerateReceiveAddressResponse, error) {
	params := make(map[string]interface{})

	if callbackURL != "" {
		params["address"] = map[string]string{
			"callback_url": callbackURL,
		}
	}

	body, err := c.PostJSON("account/generate_receive_address", params)
	if err != nil {
		return AccountGenerateReceiveAddressResponse{}, err
	}

	// parse into json
	var response AccountGenerateReceiveAddressResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AccountGenerateReceiveAddressResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Addresses
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Addresses(page int, limit int, query string) (AddressesResponse, error) {

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
		return AddressesResponse{}, err
	}

	// parse into json
	var response AddressesResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AddressesResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buttons
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buys
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Buys(quantity float32, agree_btc_amount_varies bool) (BuysResponse, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))

	if agree_btc_amount_varies {
		params.Set("agree_btc_amount_varies", "true")
	}

	body, err := c.PostForm("buys", params)
	if err != nil {
		return BuysResponse{}, err
	}

	// parse into json
	var response BuysResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return BuysResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Contacts
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Contacts() (ContactsResponse, error) {
	body, err := c.Get("contacts", nil)

	if err != nil {
		return ContactsResponse{}, err
	}

	// parse into json
	var response ContactsResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return ContactsResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Currencies
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Currencies() (CurrenciesResponse, error) {
	body, err := c.Get("currencies", nil)

	if err != nil {
		return CurrenciesResponse{}, err
	}

	// parse into json
	var response CurrenciesResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return CurrenciesResponse{}, err
	}

	return response, nil
}

func (c *Client) CurrenciesExchangeRates() (Rates, error) {
	body, err := c.Get("currencies/exchange_rates", nil)
	if err != nil {
		return Rates{}, err
	}

	// parse into json
	var response Rates
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Rates{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Orders
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Orders() (OrdersResponse, error) {
	body, err := c.Get("orders", nil)

	if err != nil {
		return OrdersResponse{}, err
	}

	// parse into json
	var response OrdersResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return OrdersResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Prices
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) PricesBuy() (PricesBuyResponse, error) {
	body, err := c.Get("prices/buy", nil)

	if err != nil {
		return PricesBuyResponse{}, err
	}

	// parse into json
	var response PricesBuyResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return PricesBuyResponse{}, err
	}

	return response, nil
}

func (c *Client) PricesSell() (PricesSellResponse, error) {
	body, err := c.Get("prices/sell", nil)

	if err != nil {
		return PricesSellResponse{}, err
	}

	// parse into json
	var response PricesSellResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return PricesSellResponse{}, err
	}

	return response, nil
}

func (c *Client) PricesSpotRate() (Amount, error) {
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

func (c *Client) PricesHistorical(page int) (string, error) {
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

func (c *Client) Sells(quantity float32) (SellsResponse, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))

	body, err := c.PostForm("sells", params)

	if err != nil {
		return SellsResponse{}, err
	}

	// parse into json
	var response SellsResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return SellsResponse{}, err
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

func (c *Client) Transfers(page int, limit int) (TransfersResponse, error) {
	params := url.Values{}

	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if limit != 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	body, err := c.Get("transfers", params)

	if err != nil {
		return TransfersResponse{}, err
	}

	// parse into json
	var response TransfersResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return TransfersResponse{}, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Users
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Users() (UsersResponse, error) {
	body, err := c.Get("users", nil)

	if err != nil {
		return UsersResponse{}, err
	}

	// parse into json
	var response UsersResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return UsersResponse{}, err
	}

	return response, nil
}
