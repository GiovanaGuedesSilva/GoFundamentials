package main

import "fmt"

func somaVariadica(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever(texto string, numeros ...int) { // Eu não posso passar dois parâmetros variádicos, apenas um, e também tem que ser o último parâmetro
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Funções variádicas permitem que você passe um número variável de argumentos para uma função.
	// Elas são úteis quando você não sabe quantos argumentos serão passados.

	fmt.Println(somaVariadica(1, 2, 3, 4, 5)) // Soma de vários números
	fmt.Println(somaVariadica(10, 20, 30))    // Soma de outros números
	// Você pode passar qualquer número de argumentos para a função variádica, inclusive nenhum argumento, considerado um slice vazio.

	// Você também pode passar slices como argumentos para funções variádicas
	numeros := []int{100, 200, 300}
	fmt.Println(somaVariadica(numeros...)) // Usando o operador ... para expandir o slice (função variádica)

	escrever("Número:", 1, 2, 3, 4, 5) // Escreve um texto seguido de vários números
	escrever("Valor:", 10, 20, 30)     // Escreve outro
}
