package main

import "fmt"

func main() {
	hobbies := map[string]string{"Larissa": "read", "Rodrigo": "fix cars", "Melissa": "play video games"}

	for name, hobby := range hobbies {
		fmt.Println("O hobby de: " + name + " é " + hobby)
	}
}
