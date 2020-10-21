package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
	fmt.Print("Please Enter Word > ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	// convert all letters to lower case
	lowerStr := strings.ToLower(userInput)

	if strings.HasPrefix(lowerStr, "i") && strings.HasSuffix(lowerStr, "n") && strings.Contains(lowerStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")

	}

}
