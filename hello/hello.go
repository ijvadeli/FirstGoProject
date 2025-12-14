// Declare main package (code executed as application should be a main package)
package main

// Import 2 packages, gives code access to those functions
import (
	"fmt" // Functions for handling in and output text (Printing text on console etc.)
	"log"

	"example.com/greetings" // Access to the Hello function from module
)

func main() {
	// Set properties of predefined logger including
	// - Log entry prefix and a flag to disable printing
	// - The time, source file and file number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Papum") // Get a greeting by calling the package's Hello function
	// If an error was returned print into console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	fmt.Println(message)
}