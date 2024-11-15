package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a slice by make function
	var gpaSlice []float64
	gpaSlice = make([]float64, 3)
	gpaSlice[0] = 4.1
	gpaSlice[1] = 4.5
	gpaSlice[2] = 3.9
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slice literal
	// gpaSlice := []float64{4.1, 4.5, 3.9} // slice literal
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slicing an existing array
	// gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	// fmt.Println(gpas, reflect.TypeOf(gpas))
	// gpaSlice := gpas[1:4] // slice := slicing array
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))
}
