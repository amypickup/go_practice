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

// check out panic vs log + exit
// tests
// err name
// type the HTML response, reader object on struct, check out "power" of structs / interfaces
// receiver methods
// Ben W

func main() {
  // https://gobyexample.com/command-line-arguments
  args := os.Args[1:]
  if len(args) != 1 {
    log.Fatalf("Incorrect number of args, please try again")
  }
  // fmt.Printf("arg provided: %v - %s, type: %T\n", os.Args[1], os.Args[1], os.Args[1])

  amount, amt_err := decimal.NewFromString(os.Args[1]) 
  // amount, amt_err := strconv.Atoi(os.Args[1])
  if amt_err != nil {
    log.Fatalf("Arg not parsed as amount, try again %s", amt_err)
  }
  fmt.Printf("arg = %v is of type %T\n", amount, amount)

  // TODO: Check if multiple use of err is ok
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

  // use a generic map for now
  // TODO: define an explicit struct
  var exchange_rates ExchangeRateResponse
  error := json.Unmarshal([]byte(body), &exchange_rates)
  if error != nil {
    log.Fatalf("Unable to marshal JSON due to %s", error)
  }
  // fmt.Println(exchange_rates.Data.Rates["ETH"])

  btc_rate, btc_err := decimal.NewFromString(exchange_rates.Data.Rates["BTC"])
  if btc_err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", btc_err)
  }
  eth_rate, eth_err := decimal.NewFromString(exchange_rates.Data.Rates["ETH"])
  if eth_err != nil {
    log.Fatalf("Unable to parse exchange rates, %s", eth_err)
  }
  // fmt.Printf("btc_rate = %T\n", btc_rate)
  // fmt.Printf("eth_rate = %T\n", eth_rate)

	// TODO: check if type inference works as expected
	// ex. var amount float32 = 100

	//dummy_amount := decimal.NewFromInt(100)
  btc_multiplier := decimal.NewFromFloat(.7)
  eth_multiplier := decimal.NewFromFloat(.3)
	btc_amount := amount.Mul(btc_multiplier).Mul(btc_rate)
	eth_amount := amount.Mul(eth_multiplier).Mul(eth_rate)

  fmt.Printf("Btc purchase: $%v USD * %v Rate = %v BTC \n", amount.Mul(btc_multiplier), btc_rate, btc_amount)
  fmt.Printf("Eth purchase: $%v USD * %v Rate = %v ETH \n", amount.Mul(eth_multiplier), eth_rate, eth_amount)

}