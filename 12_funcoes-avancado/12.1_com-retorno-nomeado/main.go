package main

import "fmt"

func calculos(a int, b int) (soma int, subtracao int) { // Vai retornar dois valores nomeados
	soma = a + b
	subtracao = a - b
	return
	// returns soma, subtracao // Também poderia ser usado assim, mas não é necessário e fica mais verboso
}

func main() {
	// Funções com retorno nomeado permitem que você retorne múltiplos valores de uma função
	// e nomeie esses valores diretamente na assinatura da função.

	fmt.Println(calculos(10, 5))        // Chama a função calculos e imprime os valores retornados
	soma, subtracao := calculos(20, 10) // Atribui os valores retornados a variáveis
	fmt.Println(soma, subtracao)
}
