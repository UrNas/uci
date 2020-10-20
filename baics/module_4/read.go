package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type person struct {
	fname string
	lname string
}

func main() {
	var persons []person
	var filename string
	fmt.Print("Please Enter fileName: ")
	if _, err := fmt.Scan(&filename); err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		d := bytes.Split(data, []byte("\n"))
		for i, v := range d {
			fullname := bytes.Split(v, []byte(" "))
			// check if line empty or if it have just first name or last name
			if len(fullname) != 2 {
				fmt.Println("#", i, " len of fullname is not 2")
				continue
			}
			fname := fullname[0]
			lname := fullname[1]
			// check length of fname and lname, it should be not more than 20 char's
			if len(fname) < 20 && len(lname) < 20 {
				persons = append(persons, person{
					fname: string(fname),
					lname: string(lname),
				})
			} else {
				if len(fname) >= 20 {
					fmt.Printf("We can not add coz len of first name %s is more then 20 chars.\n", fname)
				}
				if len(lname) >= 20 {
					fmt.Printf("We can not add coz len of last name %s is more then 20 chars.\n", lname)
				}
			}
		}
		for i, v := range persons {
			fmt.Printf("[%d] %s %s\n", i, v.fname, v.lname)
		}
	}
}
