package main

import "fmt"

func main() {
	media()
}
func media() {
	var array1 [99]float64
	// var array:=[3] float64{}
	// array1[...] = float64 {10,20,3}

	for i := 0; i < len(array1); i++ {
		array1 = array1 + array1/99
		fmt.Printf("A media  é: %f \t", array1)
	}
}
