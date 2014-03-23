package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createOrdersClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestOrders(t *testing.T) {
	c := createOrdersClient(t)

	orders, err := c.GetOrders()
	
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", orders)
}
