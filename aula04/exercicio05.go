package main

import "fmt"

//exercicio de exemplo para tratamento de erros

var lista = map[string]int{"Batata": 3, "Cenoura": 3}

func main() {
	qtdRemovida, err := removerDaListaCompras()

	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	} else {
		fmt.Println(qtdRemovida, "produtos foram removidos da lista")
	}
}

func removerDaListaCompras() (int, error) {
	produto := ""
	var removidos int
	if qtd, ok := lista[produto]; ok {
		removidos = qtd
		delete(lista, produto)
	} else {
		return 0, fmt.Errorf("produto não encontrado na lista")
	}
	return removidos, nil
}
