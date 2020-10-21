/*
Write a program which prompts the user to enter integers and stores
the integers in a sorted slice.
The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3.
During each pass through the loop,
the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice,
sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers
which the user decides to enter.
The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer
*/
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	var userIntput string
	count := 0
	for {
		fmt.Print("> ")
		if _, err := fmt.Scan(&userIntput); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if userIntput == "X" {
			break
		}
		intInput, err := strconv.Atoi(userIntput)
		if err != nil {
			fmt.Println(err)
		}
		if count < 3 {
			for i, v := range sli {
				if v == 0 {
					sli[i] = intInput
					break
				}
			}
			count++
		} else {
			sli = append(sli, intInput)
		}
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
