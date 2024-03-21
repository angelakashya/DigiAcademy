package main 

import "fmt"

func main() {

	fmt.Println("Digite um número:")
	var number1, number2 int 
	number1 = 10

	fmt.Scanln(&number2) 

	var result = number1 + number2 

	fmt.Println("O resultado da soma é:")
	fmt.Println(result)
}