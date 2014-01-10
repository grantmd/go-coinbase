package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createPricesClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestPricesBuy(t *testing.T) {
	c := createPricesClient(t)

	prices, err := c.PricesBuy()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", prices)
}

func TestPricesSell(t *testing.T) {
	c := createPricesClient(t)

	prices, err := c.PricesSell()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", prices)
}

func TestPricesSpotRate(t *testing.T) {
	c := createPricesClient(t)

	rate, err := c.PricesSpotRate()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", rate)
}

func TestPricesHistorical(t *testing.T) {
	c := createPricesClient(t)

	historical, err := c.PricesHistorical(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(historical)
}
