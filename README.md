[![Build Status](https://travis-ci.org/grantmd/go-coinbase.png?branch=master)](https://travis-ci.org/grantmd/go-coinbase)

go-coinbase
=======

A Coinbase API client in Go (golang). Focus is currently on buying and selling BTC, but eventually all methods will be supported.

Usage
-----

* First, get yourself a Coinbase api key: https://coinbase.com/account/api

* Become familiar with the rate limits that apply to your account: https://coinbase.com/verifications

* Install the library:

        go get github.com/grantmd/go-coinbase

* Include it in your project:

        import "github.com/grantmd/go-coinbase"

* Setup your Coinbase client:

        c := &coinbase.Client{
		APIKey: os.Getenv("COINBASE_API_KEY"),
	}

* Make a call with your API key:

	balance, err := c.AccountBalance()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", balance)
