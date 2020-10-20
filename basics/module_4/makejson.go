package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		person := make(map[string]string)
		fmt.Print("Please Enter Your Name > ")
		for scanner.Scan() {
			person["name"] += scanner.Text()
			break
		}
		fmt.Print("Please Enter Your Adress > ")
		for scanner.Scan() {
			person["address"] += scanner.Text()
			break
		}
		jsonData, err := json.Marshal(&person)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(jsonData))
		}
	}
}
