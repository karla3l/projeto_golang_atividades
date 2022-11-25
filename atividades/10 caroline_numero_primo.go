package main

import "fmt"

func main() {
	resultado()
}
func resultado() {
	resultado := 0
	num := 0
	fmt.Printf("Digite um numero: \t")
	fmt.Scan(&num)

	//for i := 0; i < 2; i++ {

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			resultado++
			break
		}
	}

	if resultado == 0 {
		fmt.Printf("%d é um número primo\n", num)
	} else if resultado != 0 {
		fmt.Printf("%d não é um número primo\n", num)

	}
}
