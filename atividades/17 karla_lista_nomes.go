package main

import (
	"fmt"
	"sort"
)

func main() {
	nome()
}
func nome() {
	nomes := []string{"karla", "brendo", "joao", "maria"}
	sort.Strings(nomes)
	fmt.Printf("\n aqui estão os nomes na ordem digitados: \n")
	fmt.Printf("%s\n", nomes)

	/* var N, j, i int
	//var aux string

	fmt.Printf("digite a quantidade de nomes ")
	fmt.Scan("%d", &N)
	//getchar();

	//var nomes [N]string
	//nomes[N] = make([]int, N)
	nomes := make([][]int, N)
	for i = 0; i < N; i++ {
		fmt.Printf("digite o %d nome ", i+1)
		fmt.Scan("%[^\n]s", nomes[i])
		//getchar()
	}

	fmt.Printf("\n aqui estão os nomes na ordem digitados \n")
	for i = 0; i < N; i++ {
		fmt.Printf("%s\n", nomes[i])
	}
	for i = 0; i < N; i++ {
		for j = i + 1; j < N; j++ {
			if nomes[i][0] > nomes[j][0] {
				//copy(aux, nomes[i])
				copy(nomes[i], nomes[j])
				//copy(nomes[j], aux)
			}
		}
	}
	fmt.Printf("\n aqui os nomes em ordem crescente \n")
	for i = 0; i < N; i++ {
		fmt.Printf("%s\n", nomes[i])
	} */
}
