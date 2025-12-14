// Declare a greetings package to collect related functions
// Implement a Hello function to return the greeting

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return this error in the message
	if name == "" {
		return name, errors.New("no name")
	}
	// Return a random message using randomFormat()
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat function
//? Goal: Returns one of a set of greeting messages, the returned message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string {
		"Hi %v. Welcome!",
		"Hola %v. Je ne suis pas mexicano",
		"Hoi %v. Ben je meer van de bitterballen of de frikandellen?",
	}
	// Return a random selected message from the slice of formats
	return formats[rand.Intn(len(formats))]
}