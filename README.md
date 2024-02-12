#### Prompt

This endpoint provides up-to-the-minute crypto exchange rates relative to US dollars:

[https://api.coinbase.com/v2/exchange-rates?currency=USD](https://api.coinbase.com/v2/exchange-rates?currency=USD)

Crypto is all the rage these days and I don't want to miss out! I want to keep 70% of my crypto holdings in BTC and 30% in ETH. Write a function that takes the amount I have to spend in USD as a parameter and returns the number of Bitcoin and Ethereum to buy.

I have $X I want to keep in BTC and ETH, 70/30 split. How many of each should I buy?

I'd say you should take the input amount as a command line parameter and print a json response with the allocations back out of your program.

#### Task Breakdown

Create a go program that, via CLI, requires an input arg and prints json containing the 70% BTC and 30% ETC allocations

Initial sketch of steps to explore go:

- Accept input
- Accept input as a command line parameter (args library w/ flag, help)
- Calculate + store allocations
- Print json response
- Write test file

Optional things to explore:

- Check out if docstrings
- Classes?
- Print to a file

#### Actual Steps

1. Used [Go Getting Started](https://go.dev/doc/tutorial/getting-started) to setup directory and complete hello world exapmle.
   Use `go mod init example/hello`, `go mod init example.com/greetings` to enable dependency tracking
   Use `go mod tidy` to add imported modules
   Use `go mod edit -replace example.com/greetings=../greetings` to use local vs. published package
   Go has built in testing, add \_test.go to have `go test` include it in check
   `go build` - compiles pkgs + deps, no install
   `go install` - compiles pkgs + install
   `go list -f '{{.Target}}'` - discover the go install path, where binaries are installed
