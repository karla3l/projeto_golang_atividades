package main

import "fmt"

func main() {
	nps()
}
func nps() {

	num := 0
	fmt.Printf("Net Promoter Score, o indicador de Lealdade de Clientes \n\n")
	fmt.Printf("Numa escala de 0 a 10 \n qual é a probabilidade de você recomendar a empresa X \n para um familiar ou amigo: \t\n")
	fmt.Scan(&num)

	if num == 0 || num == 1 || num == 2 || num == 3 || num == 4 || num == 5 || num == 6 {
		fmt.Printf("%d  clientes Detratores\n", num)
	}
	if num == 7 || num == 8 {
		fmt.Printf("%d clientes Neutros\n", num)
	}
	if num == 9 || num == 10 {
		fmt.Printf("%d clientes Promotores\n", num)

	}
}
