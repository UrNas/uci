package main

import (
	"fmt"
	"os"
	"time"
)

// s =½ a t2 + vot + so

func main() {
	var acceleration float64
	var velocity float64
	var displacement float64
	var period float64

	fmt.Print("Please Enter Three values > ")
	if _, err := fmt.Scanln(&acceleration, &velocity, &displacement); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fn := genDisplaceFn(acceleration, velocity, displacement)
	for {
		fmt.Print("Please Enter time > ")
		if _, err := fmt.Scan(&period); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(fn(period))
	}
}

func genDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		time.Sleep(time.Second * time.Duration(t))
		return (a/2)*(t*t) + (v * t) + s
	}
}
