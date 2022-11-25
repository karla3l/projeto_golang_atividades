package main

import "fmt"

func main() {
	calculadora()
}
func calculadora() {

	var num1, num2, soma, subtracao, divisao, multiplicacao float64
	var operacao string

	fmt.Printf("Digite um numero \t")
	fmt.Scan(&num1)
	fmt.Printf("Digite mais um numero \t")
	fmt.Scan(&num2)
	fmt.Printf("Digite a operacao: \n somar, subtrair, multiplicar ou dividir \t")
	fmt.Scan(&operacao)

	switch operacao {
	case "somar":
		soma = num1 + num2
		fmt.Printf("a soma dos numeros é: \t")
		fmt.Println(soma)
	case "subtrair":
		subtracao = num1 - num2
		fmt.Printf("a subtracao dos numeros é: \t")
		fmt.Println(subtracao)
	case "multiplicar":
		multiplicacao = num1 * num2
		fmt.Printf("a multiplicacao dos numeros é: \t")
		fmt.Println(multiplicacao)
	case "dividir":
		divisao = num1 / num2
		fmt.Printf("a divisao dos numeros é: \t")
		fmt.Println(divisao)
	default:
		fmt.Println("No information available.")
	}
}
