package main 

import "fmt"

func main () { 
	var idade int 
	var permissaoDosPais bool

	fmt.Println("Quantos anos você tem?")
	fmt.Scan(&idade)

	fmt.Println("Você tem permissão dos seus pais?")
	fmt.Scan(&permissaoDosPais)

	if idade >= 18 || (idade >= 13 && permissaoDosPais) { 
		podeAssistir := true 
	}

	podeAssistir := idade >= 18 || (idade >= 13 && permissaoDosPais)
	fmt.Println("Você pode assistir a este filme", podeAssistir)
}

