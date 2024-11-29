package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int
	for {
		fmt.Print("Name? ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("Age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}
	for k, v := range ages {
		fmt.Printf("%s is %d years old~\n", k, v)
	}
}
