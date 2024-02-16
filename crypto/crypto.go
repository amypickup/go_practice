package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/shopspring/decimal"
)

const COINBASE_URL string = "https://api.coinbase.com/v2/exchange-rates?currency=USD"

type ExchangeRates struct {
  ToCurrency string
  Rate string
}

type ExchangeRateResponse struct {
	Data struct {
		FromCurrency string `json:"currency"`
		Rates map[string]string `json:"rates"`
	}
}

// check out panic vs log + exit
// tests
// err name
// type the HTML response, reader object on struct, check out "power" of structs / interfaces
// receiver methods
// Ben W

func main() {
  // Used https://gobyexample.com/command-line-flags and https://pkg.go.dev/flag
  a_flag := flag.String("a", "", "the amount you plan to spend in USD")
  flag.Parse()
  
  amount, err := decimal.NewFromString(*a_flag)
  if err != nil {
    log.Fatalf("Arg not parsed as amount, try again %s", err)
  }
  // fmt.Printf("arg = %v is of type %T\n", amount, amount)

  // suggested usage of Get with ReadAll: https://pkg.go.dev/net/http
  resp, err := http.Get(COINBASE_URL)
  if err != nil {
    log.Fatalf("Unable to fetch coinbase url json due to %s", err)
  }
  defer resp.Body.Close()
  // fmt.Println(resp)

  // TODO: Not sure if necessary, can you skip directly to unmarshalling?
	body, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Fatalf("Unable to readall json due to %s", err)
  }
  // fmt.Printf("%s", body)

  // TODO: Confirm struct definition is decent practice
  // Ex. should use explicit struct w/ all currency defs like json to go
  // or is generic ok? vs. ok to use map vs. named strings

  // Used https://pkg.go.dev/encoding/json
  var exchange_rates ExchangeRateResponse
  if err := json.Unmarshal([]byte(body), &exchange_rates); err != nil {
    log.Fatalf("Unable to marshal JSON due to %s", err)
  }
  // fmt.Printf("Found eth value in json: %v", exchange_rates.Data.Rates["ETH"])

  btc_rate, err := decimal.NewFromString(exchange_rates.Data.Rates["BTC"])
  if err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", err)
  }
  eth_rate, err := decimal.NewFromString(exchange_rates.Data.Rates["ETH"])
  if err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", err)
  }
  // fmt.Printf("btc_rate = %T\n", btc_rate)
  // fmt.Printf("eth_rate = %T\n", eth_rate)

  // TODO: check if type inference / money calcs works as expected

  //dummy_amount := decimal.NewFromInt(100)
  btc_multiplier := decimal.NewFromFloat(.7)
  eth_multiplier := decimal.NewFromFloat(.3)
  btc_amount := amount.Mul(btc_multiplier).Mul(btc_rate)
  eth_amount := amount.Mul(eth_multiplier).Mul(eth_rate)

  fmt.Printf("Btc purchase: $%v USD * %v Rate = %v BTC \n", amount.Mul(btc_multiplier), btc_rate, btc_amount)
  fmt.Printf("Eth purchase: $%v USD * %v Rate = %v ETH \n", amount.Mul(eth_multiplier), eth_rate, eth_amount)

}