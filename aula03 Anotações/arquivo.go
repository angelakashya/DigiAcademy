package main 

//Import de Libs/Bibliotecas
import "fmt"

var variavelGlobal string 

func FuncaoPublica() {
	variavel := "SóDessaFunçãoPublica"
	funçãoPrivada()
	fmt.Print(variavel)

	variavelGlobal = "variavelNaPublica"
}

func FuncaoPrivada() {
	variavel := "SóDessaFunçãoPrivada"
	fmt.Print(variavel)
	//Escopo da função 
	variavelGlobal= "variavelNaPrivada"
}
