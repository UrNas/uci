package main

import (
	"fmt"
	"strings"
)

// Animal interface describes the methods of an animal
type Animal interface {
	Speak()
	Eat()
	Move()
}

// Cow repersent animal
type Cow struct {
	name, food, noise, locomotion string
}

// Speak repersent animal noise
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// Eat repersent animal food
func (c Cow) Eat() {
	fmt.Println(c.food)
}

// Move repersent animal locomotion
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

// Bird repersent animal
type Bird struct {
	name, food, noise, locomotion string
}

// Speak repersent animal noise
func (b Bird) Speak() {
	fmt.Println(b.noise)
}

// Eat repersent animal food
func (b Bird) Eat() {
	fmt.Println(b.food)
}

// Move represent animal locomotion
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

// Snake repersent animal
type Snake struct {
	name, food, noise, locomotion string
}

// Speak represent animal noise
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

// Eat represent animal food
func (s Snake) Eat() {
	fmt.Println(s.food)
}

// Move represent animal locomotion
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func main() {
	animals := make(map[string]Animal)
	var command, name, thirdCmd string

	for {
		fmt.Print("Command > ")
		fmt.Scan(&command, &name, &thirdCmd)
		switch command {
		case "newanimal":
			switch strings.ToLower(thirdCmd) {
			case "cow":
				animals[name] = Cow{
					food:       "grass",
					noise:      "moo",
					locomotion: "walk",
				}
				fmt.Println("Create it!")
			case "bird":
				animals[name] = Bird{
					food:       "worms",
					noise:      "peep",
					locomotion: "fly",
				}
				fmt.Println("Create it!")
			case "snake":
				animals[name] = Snake{
					food:       "mice",
					noise:      "hsss",
					locomotion: "slither",
				}
				fmt.Println("Create it!")
			default:
				fmt.Println("Animal Not exists")
			}
		case "query":
			switch strings.ToLower(thirdCmd) {
			case "eat":
				if checkQuery(name, animals) {
					animals[name].Eat()
				}
			case "speak":
				if checkQuery(name, animals) {
					animals[name].Speak()
				}
			case "move":
				if checkQuery(name, animals) {
					animals[name].Move()
				}
			default:
				fmt.Println("Not exists")
			}
		default:
			fmt.Println("command not found")
		}
	}
}

// checkQuery check if animal name is exists before fire up methods
func checkQuery(name string, an map[string]Animal) bool {
	_, ok := an[name]
	if ok {
		return true
	}
	fmt.Printf("%s can not found\n", name)
	return false
}
