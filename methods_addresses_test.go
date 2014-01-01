package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createAddressesClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestAddresses(t *testing.T) {
	c := createAddressesClient(t)

	addresses, err := c.Addresses()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(addresses)
}
