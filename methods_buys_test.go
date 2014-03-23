package coinbase

import (
	"os"
	"testing"
)

func createBuysClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestBuys(t *testing.T) {
	/*
		c := createBuysClient(t)

		users, err := c.GetBuys(0.0003, false)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%+v\n", users)
	*/
}
