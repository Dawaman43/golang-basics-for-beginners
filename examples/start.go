// Hello everyone this is easy explanation in Golang.
// I provided pure code with explaining comments so don’t get frustrated.
// Please try to write your own code after understanding this.

// In Go, first we have to package the .go files or code files to make containerize them
package main

// Every built-in function comes from a library (also called package).
// In this lesson we cover 4 libraries: "bufio", "fmt", "os" and "strings"

// First we have to declare the libraries at the top inside import()
import (
	"bufio"   // Used for reading input from users
	"fmt"     // (short for "format") The basic library used in Go for formatting and printing text
	"os"      // Provides functionality for interacting with the operating system
	"strings" // Provides functionality to manipulate and work with strings
)

// When we run a Go program, the entry point is the main function
// To define a function, we use the func keyword
func main() {
	// When declaring a new variable and assigning a value, we use := operator
	textToPrintInTheTerminal := "Hello world"

	// Print the text to the terminal using fmt.Println
	fmt.Println(textToPrintInTheTerminal)

	// Create a reader to take user input from the terminal
	reader := bufio.NewReader(os.Stdin)

	// Ask the user to type their name
	fmt.Print("Enter your name: ")

	// Read user input until they press Enter (\n)
	userInput, _ := reader.ReadString('\n')

	// Remove the newline character from the input using strings.TrimSpace
	cleanedInput := strings.TrimSpace(userInput)

	// Greet the user with their name
	fmt.Printf("Hello, %s! Welcome to Golang.\n", cleanedInput)
}
