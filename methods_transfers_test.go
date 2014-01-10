package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createTransfersClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestTransfers(t *testing.T) {
	c := createTransfersClient(t)

	transfers, err := c.Transfers(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", transfers)
}
