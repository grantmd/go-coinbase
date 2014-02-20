package coinbase

type Amount struct {
	Amount   float32 `json:",string"`
	Currency string
}
