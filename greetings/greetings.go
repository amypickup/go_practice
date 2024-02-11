// https://go.dev/doc/tutorial/create-module

// collect related functions
package greetings

import "fmt"

// Hello returns a greeting for the named person.
// Capitalized function names are "exported names"
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}