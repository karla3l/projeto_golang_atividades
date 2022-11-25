package main

import "fmt"

func main() {
	teste()
}
func teste() {
	var array1 := [3] float64                        // var array:=[3] float64{}
	array1[0], array1[1], array1[2] = 10, 20, 30 // array1[...] = float64 {10,20,3}

	for i := 0; i < len(array1); i++ {
		fmt.Printf("O array  é: %f \t", array1)
	}
}
// Questão 01/10
// 01 - Escreva um algoritmo em Go usando array, que calcule a média simples de 99 valores pontos flutuantes
// 02 - Escreva um algoritmo em Go usando array, contendo 5 valores inteiros de 8bits e calcule um array de saída de ponto flutuante, com valores normalizados de [0:1] e escreva o resultado
// normalizado é mudar os range de dominio
//append / make