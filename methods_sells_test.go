package coinbase

import (
	"os"
	"testing"
)

func createSellsClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestSells(t *testing.T) {
	/*
		c := createSellsClient(t)

		sells, err := c.GetSells(0.0003)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%+v\n", sells)
	*/
}
