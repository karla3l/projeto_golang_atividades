package main

import "fmt"

func main() {
	temperatura()
}
func temperatura() {
	var fahrenheit, celsius float64

	fmt.Printf("Digite a temperatura em celsius: \t")
	fmt.Scan(&celsius)
	fahrenheit = celsius*1.8 + 32

	fmt.Printf("a temperatura de celsius para fahrenheit é: \t")
	fmt.Println(fahrenheit)

}
