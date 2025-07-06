package main

import "fmt"

func main() {
	// Funções anônimas são funções que não possuem um nome associado a elas.
	// Elas podem ser atribuídas a variáveis, passadas como argumentos para outras funções ou retornadas de funções.
	// Elas são úteis quando você precisa de uma função temporária ou de uma função de callback.

	func() {
		fmt.Println("Essa é uma função anônima!")
	}() // Chamada imediata da função anônima

	func(nome string) {
		fmt.Printf("Olá, %s\n", nome) // printf permite formatação de strings (diferente do Println, que não formata) (concatenação)
	}("João") // Chamada da função anônima com um argumento

	func(nome string) string {
		return "Olá, " + nome
	}("Maria") // Chamada da função anônima com um argumento e retorno

	retorno := func(nome string) string {
		return "Olá, " + nome
	}("Maria")

	fmt.Println(retorno) // Exibe o retorno da função anônima
}
