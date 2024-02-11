// https://go.dev/doc/tutorial/create-module

// collect related functions
package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Capitalized function names are "exported names"
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}