package main

import "fmt"

func main() {
	min()
}
func min() {
	var segundos, minutos float64

	fmt.Printf("Digite quantos segundos: \t")
	fmt.Scan(&segundos)
	minutos = segundos / 60

	fmt.Printf("Os segundos convertidos em minutos são: %fm/s \t", minutos)

}
