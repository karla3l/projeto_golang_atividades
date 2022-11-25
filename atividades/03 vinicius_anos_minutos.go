package main

import "fmt"

func main() {
	ano()
}
func ano() {
	var anos, minutos float64

	fmt.Printf("Digite quantos anos: \t")
	fmt.Scan(&anos)
	minutos = anos * 525960

	fmt.Printf("Os anos convertidos em minutos são: %fm/s \t", minutos)

}
