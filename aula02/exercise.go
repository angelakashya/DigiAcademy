package main 

import "fmt"

func main() {
	fmt.Println("Digite seu nome:")
	var name string 

	fmt.Scanln(&name) 

	fmt.Print ("Bem vindo, " + name)

}
