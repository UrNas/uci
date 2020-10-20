package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intsChan := make(chan []int, 4)
	var numbers []int
	var finalSortList []int
	counter := 0
	prompt(&numbers)
	if len(numbers) < 4 {
		fmt.Println("Please Enter At least four numbers")
		os.Exit(1)
	}
	groups := setGroups(numbers, 4)
	// create four goroutine and pass goroutine ID, array of ints and recieve chanel as values
	for i, v := range groups {
		go func(id int, arr []int, c chan<- []int) {
			sort.Ints(arr)
			fmt.Printf("[#%v] %v\n", id, arr)
			c <- arr
		}(i+1, v, intsChan)
	}
	for counter < 4 {
		dataOfInts := <-intsChan
		finalSortList = append(finalSortList, dataOfInts...)
		counter++
	}
	sort.Ints(finalSortList)
	fmt.Println(finalSortList)

}

// prompt user to input list of numbers
func prompt(arr *[]int) {
	fmt.Print("Enter List Of Numbers > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strList := strings.Split(scanner.Text(), " ")
		for _, v := range strList {
			if v == "" {
				continue
			}
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			*arr = append(*arr, n)
		}
		break
	}

}

// setGroups is generic function to group array based on groups para.
func setGroups(arr []int, groups int) [][]int {
	var gnumbers [][]int
	lenArr := len(arr)
	for lenArr > 0 {
		for _, v := range arr {
			if len(gnumbers) < groups {
				gnumbers = append(gnumbers, []int{v})
			} else {
				firstEle := gnumbers[0]
				firstEle = append(firstEle, v)
				gnumbers = append(gnumbers[1:], firstEle)
			}

			lenArr--
		}
	}
	return gnumbers
}
