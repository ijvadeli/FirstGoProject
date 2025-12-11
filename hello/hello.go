// Declare main package (code executed as application should be a main package)
package main

// Import 2 packages, gives code access to those functions
import (
	"fmt" // Functions for handling in and output text (Printing text on console etc.)

	"example.com/greetings" // Access to the Hello function from module
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Papum") // Get a greeting by calling the package's Hello function
	fmt.Println(message)
}