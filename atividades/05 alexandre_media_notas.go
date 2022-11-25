package main

import "fmt"

func main() {
	soma()
}
func soma() {
	var soma, media float64
	for i := 0; i < 5; i++ {
		var notas float64
		fmt.Printf("Digite uma nota: \t")
		fmt.Scan(&notas)
		soma = notas + soma
		media = soma / 5

	}

	fmt.Printf("A média das notas é: \t")
	fmt.Println(media)

}
