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
