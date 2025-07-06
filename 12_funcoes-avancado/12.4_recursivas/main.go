package main

import "fmt"

func fibonacci(posicao uint) int {
	if posicao <= 1 {
		return int(posicao)
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

// Estouro de pilha (stack overflow) ocorre quando uma função recursiva chama a si mesma muitas vezes
// Sem atingir uma condição de parada, resultando em um loop infinito.
// Isso pode acontecer se a condição de parada não for bem definida ou se o número de chamadas recursivas for muito grande.

func main() {
	// Funções recursivas são funções que se chamam a si mesmas durante sua execução.
	// Elas são úteis para resolver problemas que podem ser divididos em subproblemas menores, como o cálculo de fatorial.

	// Definindo fatorial como uma função nomeada para permitir recursão
	var fatorial func(n int) int
	fatorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fatorial(n-1)
	}

	fmt.Println(fatorial(5)) // Exibe 120

	// FIBONACCCI
	// 1 1 2 3 5 8 13 21 34 55 89

	posicao := uint(10)
	fmt.Printf("Fibonacci na posição %d: %d\n", posicao, fibonacci(posicao))
	// Fibonacci na posição 10: 55

	for i := uint(0); i <= posicao; i++ {
		fmt.Printf("Fibonacci na posição %d: %d\n", i, fibonacci(i))
	}
}
