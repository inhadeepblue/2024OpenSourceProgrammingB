package main

import (
	"fmt"
	"log"
	"week11/keyboard"
)

func main() {
	//n, _ := keyboard.GetFloat()
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
