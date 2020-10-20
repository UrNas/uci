package main

import (
	"fmt"
	"os"
	"time"
)

// s =½ a t2 + vot + so

func main() {
	var acceraion float64
	var velocity float64
	var displacemnt float64
	var peroid float64

	fmt.Print("Please Enter Three values > ")
	if _, err := fmt.Scanln(&acceraion, &velocity, &displacemnt); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fn := genDisplaceFn(acceraion, velocity, displacemnt)
	for {
		fmt.Print("Please Enter time > ")
		if _, err := fmt.Scan(&peroid); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(fn(peroid))
	}
}

func genDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		time.Sleep(time.Second * time.Duration(t))
		return (a/2)*(t*t) + (v * t) + s
	}
}
