package main

import "fmt"

func main() {
	dolar()
}
func dolar() {
	var num, dolar float64

	fmt.Printf("Digite o valor em real: \t")
	fmt.Scan(&num)
	dolar = num * 5.17

	fmt.Printf("O valor convertido para dolar é: \t")
	fmt.Println(dolar)

}
