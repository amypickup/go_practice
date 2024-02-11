// https://go.dev/doc/tutorial/getting-started
// https://go.dev/doc/tutorial/call-module-code

package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}