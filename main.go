package main

import (
	"fmt"
	ent "internship/golang/entity"
)

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	var names = []string{"a", "b", "c", "d", "e"}
	person := ent.Person{Name: "Prince", Gender: "Male"}
	fmt.Println(person.IsMale())
	for i, name := range names {
		fmt.Println(name, "At Postition", i)
	}
}

func ReturnTypes() (name string, age int) {
	return "name", 23
}
