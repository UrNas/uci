package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := prompt()
	userInput := setUpSlice(data)
	bubbleSort(userInput)
	fmt.Println(userInput)
}

// prompt user input
func prompt() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please Enter a sequence of up to 10 integers > ")
	for scanner.Scan() {
		data = scanner.Text()
		break
	}
	return data
}

// set up slice of int's
func setUpSlice(data string) []int {
	var numbers []int
	for _, v := range strings.Split(data, " ") {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

// Bubble Sort is the simplest sorting algorithm
func bubbleSort(numbers []int) {
	sliLen := len(numbers)
	loops := sliLen - 1
	for loops > 0 {
		for i := range numbers {
			if i != sliLen-1 {
				swap(numbers, i)
			}
		}
		loops--
	}
}

// swap numbers in slice
func swap(sli []int, idx int) {
	currentNumber := sli[idx]
	nextNumber := sli[idx+1]
	if currentNumber > nextNumber {
		sli[idx] = nextNumber
		sli[idx+1] = currentNumber
	}
}
