package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createAccountClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestAccountBalance(t *testing.T) {
	c := createAccountClient(t)

	balance, err := c.AccountBalance()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(balance)
}
