package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Print("Enter name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter address: ")
	fmt.Scanf("%s", &address)

	myMap := map[string]string{"name": name, "address": address}
	jsonData, _ := json.Marshal(myMap)
	fmt.Println(string(jsonData))
}
