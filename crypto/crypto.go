package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func main() {

  // TODO: Check if multiple use of err is ok
  // suggested usage of Get with ReadAll: https://pkg.go.dev/net/http
  resp, err := http.Get(COINBASE_URL)
  if err != nil {
      log.Fatalf("Unable to fetch coinbase url json due to %s", err)
      os.Exit(1)
  }
	defer resp.Body.Close()
  // fmt.Println(resp)

  // TODO: Not sure if necessary, can you skip directly to unmarshalling?
	body, err := io.ReadAll(resp.Body)
  if err != nil {
		log.Fatalf("Unable to readall json due to %s", err)
    os.Exit(1)
  }
	// fmt.Printf("%s", body)

  // use a generic map for now
  // TODO: define an explicit struct
  var exchange_rates ExchangeRateResponse
  error := json.Unmarshal([]byte(body), &exchange_rates)
  if error != nil {
    log.Fatalf("Unable to marshal JSON due to %s", error)
    os.Exit(1)
  }
  // fmt.Println(exchange_rates.Data.Rates["ETH"])

  btc_rate, btc_err := decimal.NewFromString(exchange_rates.Data.Rates["BTC"])
  if btc_err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", btc_err)
    os.Exit(1)
  }
  eth_rate, eth_err := decimal.NewFromString(exchange_rates.Data.Rates["ETH"])
  if eth_err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", eth_err)
    os.Exit(1)
  }
  // fmt.Printf("btc_rate = %T\n", btc_rate)
  // fmt.Printf("eth_rate = %T\n", eth_rate)

	// TODO: check if type inference works as expected
	// ex. var amount float32 = 100

	amount := decimal.NewFromInt(100)
  btc_multiplier := decimal.NewFromFloat(.7)
  eth_multiplier := decimal.NewFromFloat(.3)
	btc_amount := amount.Mul(btc_multiplier).Mul(btc_rate)
	eth_amount := amount.Mul(eth_multiplier).Mul(eth_rate)

  fmt.Printf("Btc purchase: %v * %v = %v \n", amount.Mul(btc_multiplier), btc_rate, btc_amount)
  fmt.Printf("Eth purchase: %v\n", eth_amount)

}