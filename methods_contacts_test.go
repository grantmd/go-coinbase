package coinbase

import (
	"fmt"
	"os"
	"testing"
)

func createContactsClient(t *testing.T) (c *Client) {
	c = &Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

	if c.APIKey == "" {
		t.Skip("Coinbase api key is missing (should be in the COINBASE_API_KEY environment variable)")
	}

	return c
}

func TestContacts(t *testing.T) {
	c := createContactsClient(t)

	contacts, err := c.GetContacts()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", contacts)
}
