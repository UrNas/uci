package main

import (
	"fmt"
)

// Animal represents animal info
type Animal struct {
	noise      string
	food       string
	locomotion string
}

// Speak return back noise of animal
func (an Animal) Speak() {
	fmt.Println(an.noise)
}

// Eat return back food of animal
func (an Animal) Eat() {
	fmt.Println(an.food)
}

// Move return back locomotion of animal
func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func main() {
	var name, action string
	animals := map[string]Animal{
		"cow": Animal{
			noise:      "moo",
			food:       "grass",
			locomotion: "walk",
		},
		"bird": Animal{
			noise:      "peep",
			food:       "worms",
			locomotion: "fly",
		},
		"snake": Animal{
			noise:      "hsss",
			food:       "mice",
			locomotion: "slither",
		},
	}
	for {
		fmt.Print("Please Enter Request > ")
		resetVars(&name, &action)
		_, err := fmt.Scanln(&name, &action)
		// check if user input more than two values
		if err != nil {
			fmt.Println()
			continue
		}

		requestAnimal, ok := animals[name]
		// check if user input unexists animal
		if !ok {
			continue
		}
		switch action {
		case "speak":
			requestAnimal.Speak()
		case "eat":
			requestAnimal.Eat()
		case "move":
			requestAnimal.Move()
		default:
			if action == "" {
				fmt.Println("No Action")
			} else {
				fmt.Printf("%s Not Exists\n", action)
			}
		}
	}
}
func resetVars(name, action *string) {
	*name = ""
	*action = ""
}
