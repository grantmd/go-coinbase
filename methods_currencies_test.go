package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createCurrenciesClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestCurrencies(t *testing.T) {
	c := createCurrenciesClient(t)

	currencies, err := c.GetCurrencies()
	
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", currencies)
}

func TestCurrenciesExchangeRates(t *testing.T) {
	c := createCurrenciesClient(t)

	rates, err := c.GetExchangeRates()
	
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", rates)
}
