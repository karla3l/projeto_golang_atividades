package main

import "fmt"

func main() {
	mult()
}
func mult() {
	mult := 1
	for i := 0; i < 3; i++ {
		num := 0
		fmt.Printf("Digite um numero: \t")
		fmt.Scan(&num)
		mult = mult * num

	}

	fmt.Printf("Multiplicao dos 3 números digiados é: \t")
	fmt.Println(mult)

}
