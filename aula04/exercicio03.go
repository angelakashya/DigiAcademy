package main

import (
	"fmt"
)

func main() {
	contacts := map[string]int{"Larissa": 93003092, "Rodrigo": 983468081, "Melissa": 991363091}

	//votos := make(map[string]int) creates an empty map

	fmt.Println(contacts["Rodrigo"])
}
