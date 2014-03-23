package coinbase

type Amount struct {
	Amount   float32 `json:",string"`
	Currency string
}

type CentsAmount struct {
	Cents       int
	CurrencyISO string `json:"currency_iso"`
}

type Transfer struct {
	ID            string
	Type          string
	Code          string
	CreatedAt     string `json:"created_at"`
	Fees          map[string]CentsAmount
	Status        string
	PayoutDate    string `json:"payout_date"`
	TransactionID string `json:"transaction_id"`
	BTC           Amount
	Subtotal      Amount
	Total         Amount
	Description   string
}

//////////////////////////////////////////////////////////////////////////////////////////////
// Account
//////////////////////////////////////////////////////////////////////////////////////////////

type AccountReceiveAddress struct {
	Success     bool
	Address     string
	CallbackUrl string `json:"callback_url"`
}

//////////////////////////////////////////////////////////////////////////////////////////////
// Addresses
//////////////////////////////////////////////////////////////////////////////////////////////

type Address struct {
	Address struct {
		Address     string
		CallbackUrl string `json:"callback_url"`
		Label       string
		CreatedAt   string `json:"created_at"`
	}
}

type Addresses struct {
	Addresses   []Address
	TotalCount  int `json:"total_count"`
	NumPages    int `json:"num_pages"`
	CurrentPage int `json:"current_page"`
}

//////////////////////////////////////////////////////////////////////////////////////////////
// Contacts
//////////////////////////////////////////////////////////////////////////////////////////////

type Contact struct {
	Email string
}

type Contacts struct {
	Contacts    []Contact
	TotalCount  int
	NumPages    int
	CurrentPage int
}

//////////////////////////////////////////////////////////////////////////////////////////////
// Currencies
//////////////////////////////////////////////////////////////////////////////////////////////

type ExchangeRates map[string]string
type Currencies [][]string

//////////////////////////////////////////////////////////////////////////////////////////////
// Orders
//////////////////////////////////////////////////////////////////////////////////////////////

type TotalCurrency struct {
	Cents       int
	CurrencyISO string
}

type Button struct {
	Type        string
	Name        string
	Description string
	ID          string
}

type Transaction struct {
	ID            string
	Hash          string
	Confirmations int
}

type Order struct {
	ID        string
	CreatedAt string
	Status    string

	TotalBTC    TotalCurrency
	TotalNative TotalCurrency

	Custom      string
	Button      Button
	Transaction Transaction
}

type Orders struct {
	Orders      []Order
	TotalCount  int
	NumPages    int
	CurrentPage int
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Prices
///////////////////////////////////////////////////////////////////////////////////////////////////////

type PricesBuy struct {
	SubTotal Amount
	Fees     []map[string]Amount
	Total    Amount
}

type PricesSell struct {
	SubTotal Amount
	Fees     []map[string]Amount
	Total    Amount
	Amount   string
	Currency string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Sells
///////////////////////////////////////////////////////////////////////////////////////////////////////

type Sells struct {
	Success  bool
	Errors   []string
	Transfer Transfer
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buys
///////////////////////////////////////////////////////////////////////////////////////////////////////

type Buys struct {
	Success  bool
	Errors   []string
	Transfer Transfer
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Transfers
///////////////////////////////////////////////////////////////////////////////////////////////////////

type Transfers struct {
	Transfers   []map[string]Transfer
	TotalCount  int `json:"total_count"`
	NumPages    int `json:"num_pages"`
	CurrentPage int `json:"current_page"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Users
///////////////////////////////////////////////////////////////////////////////////////////////////////

type User struct {
	User struct {
		ID             string
		Name           string
		Email          string
		TimeZone       string `json:"time_zone"`
		NativeCurrency string `json:"native_currency"`
		Balance        Amount
		BuyLevel       int    `json:"buy_level"`
		SellLevel      int    `json:"sell_level"`
		BuyLimit       Amount `json:"buy_limit"`
		SellLimit      Amount `json:"sell_limit"`
	}
}

type Users struct {
	Users []User
}
