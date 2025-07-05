package main

import "fmt"

func main() {
	// Funções podem receber parâmetros e retornar valores.

	fmt.Println(somar(3, 5)) // Chamada da função somar com os parâmetros 3 e 5
	// A função somar recebe dois parâmetros do tipo int e retorna um valor do tipo int

	var f = func(txt string) string {
		fmt.Println("Função anônima chamada com texto:", txt)
		return "Texto retornado"
		// Esta é uma função anônima que recebe um parâmetro do tipo string e
	}

	resultado := f("Teste") // Chamada da função anônima
	fmt.Println("Resultado da função anônima:", resultado)

	r1, r2, r3, r4 := calculos(10, 5)
	r1_, _, r3_, _ := calculos(10, 5) // Quando uma função retorna mais de um valor, você pode escolher quais valores deseja capturar
	// Aqui, r1_ e r3_ capturam os primeiros e terceiros valores retornados, respectivamente,
	// enquanto os segundos e quartos valores são descartados com o uso de _ (underscore).
	fmt.Println("Resultados dos cálculos:", r1, r2, r3, r4)
	fmt.Println("Resultados dos cálculos novamente:", r1_, r3_)

	
}

func somar(a, b int) int {
	return a + b
}

func calculos(a, b int) (int, int, int, string) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	final := "fim"
	return soma, subtracao, multiplicacao, final
}
