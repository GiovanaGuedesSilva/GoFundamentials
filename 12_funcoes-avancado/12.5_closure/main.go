package main

import (
	"fmt"
)

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	// Closure é uma função que captura o ambiente em que foi criada, permitindo acessar variáveis externas mesmo após a função ter sido executada.
	// É útil para criar funções com estado ou para encapsular lógica.

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure() // Chama a função closure, que retorna uma função anônima
	funcaoNova()            // Chama a função anônima retornada, que imprime o texto

	contador := 0

	incrementar := func() int {
		contador++
		return contador
	}

	fmt.Println(incrementar()) // 1
	fmt.Println(incrementar()) // 2
	fmt.Println(incrementar()) // 3

	// A variável contador é capturada pela função incrementar, permitindo que seu valor seja mantido entre as chamadas.
}
