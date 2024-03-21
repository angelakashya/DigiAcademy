package main 

import "fmt"

func main() { 
	var numero int 

	fmt.Println("Digite um número: ")
	fmt.Scanln(&numero)

	if numero > 0  { 
		fmt.Print("positivo")
	} else if numero < 0  {
		fmt.Print("negativo")
	} else {
		fmt.Print("zero")
	}		
}