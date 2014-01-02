package coinbase

import (
	"encoding/json"
	"testing"
)

func TestUnauthed(t *testing.T) {
	c := &Client{}

	body, err := c.Get("prices/spot_rate", nil)
	if err != nil {
		t.Fatal(err)
	}

	type Response struct {
		Amount   string
		Currency string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Currency != "USD" {
		t.Fatalf("Expected currency USD, got %s", response.Currency)
	}
}
