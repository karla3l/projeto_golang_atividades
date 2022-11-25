package main

import "fmt"

func main() {
	soma()
}
func soma() {
	var soma float64
	for i := 0; i < 2; i++ {
		var notas float64
		fmt.Printf("Digite uma nota: \t")
		fmt.Scan(&notas)
		soma = notas + soma

	}

	fmt.Printf("A soma das notas é: \t")
	fmt.Println(soma)

}
