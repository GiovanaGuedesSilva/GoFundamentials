package main

import "fmt"

func main() {

	var v1 string = "v1" // Declaração de variável com tipo explícito
	v2 := "v2"           // Declaração de variável com inferência de tipo

	// O Go não deixa declarar variáveis sem uso

	fmt.Println("Variável 1:", v1, "\nVariável 2:", v2)

	var ( // Declaração de múltiplas variáveis com tipo explícito
		v3 string = "v3"
		v4 string = "v4"
	)

	fmt.Println("Variável 3:", v3, "\nVariável 4:", v4)

	v5, v6 := "v5", "v6" // Declaração de múltiplas variáveis com inferência de tipo
	fmt.Println("Variável 5:", v5, "\nVariável 6:", v6)

	const c1 string = "c1" // Declaração de constante com tipo explícito
	fmt.Println("Constante 1:", c1)

	v5, v5 = v6, v5 // Troca os valores de v5 e v6
	fmt.Println("Variável 5:", v5, "\nVariável 6:", v6)
}
