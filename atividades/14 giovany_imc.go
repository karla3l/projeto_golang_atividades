package main

import "fmt"

func main() {
	imc()
}
func imc() {
	var peso, altura, imc float64

	fmt.Printf("Digite seu peso: \t")
	fmt.Scan(&peso)
	fmt.Printf("Digite sua altura: \t")
	fmt.Scan(&altura)
	imc = peso / (altura * altura)

	fmt.Printf("seu imc é: \t")
	fmt.Println(imc)

}
