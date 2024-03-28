package main

import "fmt"

func main() {

	const oi = "Hello"

	s := make([]string, 0, 4)
	s = append(s, oi, "Maria", "Ana", "Julia", "Rafaela")

	v := s[2]

	fmt.Println("esse é meu slice", s)
	//imprimindo a frase "Hello, Fulano"
	fmt.Println(s[0], v)
}
