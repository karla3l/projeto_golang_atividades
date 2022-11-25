package main

import "fmt"

func main() {
	arvore()
}
func arvore() {

	for i := 0; i < 1; i++ {
		num := 0
		fmt.Printf("Digite 0 para construir a arvore de natal  \t")
		fmt.Scan(&num)
		fmt.Printf("\t\t\t             * \n")
		fmt.Printf("\t\t\t            **** \n")
		fmt.Printf("\t\t\t           **o**** \n")
		fmt.Printf("\t\t\t          ****o**** \n")
		fmt.Printf("\t\t\t         ***o****o** \n")
		fmt.Printf("\t\t\t       ******o******o* \n")
		fmt.Printf("\t\t\t      **o*********o***** \n")
		fmt.Printf("\t\t\t     *******o************* \n")
		fmt.Printf("\t\t\t    *o***********o*****o*** \n")
		fmt.Printf("\t\t\t    **o***********o*****o*** \n")
		fmt.Printf("\t\t\t  *o******o*******o***o******* \n")
		fmt.Printf("\t\t\t****o******o*******o***o********* \n")
		fmt.Printf("\t\t\t          ########### \n")
		fmt.Printf("\t\t\t          ########### \n\n")
		fmt.Printf("\t\t   _____A R V O R E____D E____ N A T A L____")

	}
}
