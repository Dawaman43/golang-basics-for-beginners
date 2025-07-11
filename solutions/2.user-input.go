package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Program to accept user input")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write your name : ")

	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.TrimSpace(userInput)

	fmt.Printf("Hello, %s", cleanedInput)
}
