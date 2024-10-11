package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i)

	i = 3
	f = 2.7
	fmt.Print("\n\n", f <= float64(i), "\n")
}
