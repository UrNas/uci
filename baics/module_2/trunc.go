package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var fnumber float64
	fmt.Print("Please Enter Float Number: ")
	if _, err := fmt.Scan(&fnumber); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	truncNumber := math.Trunc(fnumber)
	fmt.Println(truncNumber)
}
