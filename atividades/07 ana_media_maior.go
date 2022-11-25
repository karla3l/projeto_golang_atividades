package main

import "fmt"

func main() {
	soma()
}
func soma() {
	soma := 0
	media := 0
	maior := 0
	menor := 0
	for i := 0; i < 10; i++ {
		num := 0
		fmt.Printf("Digite um numero: \t")
		fmt.Scan(&num)
		soma = num + soma
		media = soma / 10

		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}
	}
	fmt.Printf("Média dos números digiados é: \t")
	fmt.Println(media)
	fmt.Printf("Maior numero digitado é: \t")
	fmt.Println(maior)
	//fmt.Printf("Menor numero \t")
	//fmt.Println(menor)

}
