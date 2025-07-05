package main

import "fmt"

func main() {
	// Ponteiros são variáveis que armazenam o endereço de memória de outra variável.
	// Eles permitem manipular diretamente o valor de uma variável sem precisar copiá-lo.

	// Declaração de uma variável do tipo inteiro
	var x int = 10
	var y int = x // Atribuição do valor de x a y (copia o valor, não o endereço)

	// Declaração de um ponteiro para um inteiro
	var p *int = &x // O operador & obtém o endereço de memória da variável x

	// Imprime o valor da variável x e o endereço de memória armazenado no ponteiro p
	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)
	fmt.Println("Endereço de memória de x:", p)

	// Modifica o valor da variável x através do ponteiro p
	*p = 20 // O operador * desreferencia o ponteiro, permitindo acessar e modificar o valor apontado

	fmt.Println("Novo valor de x:", x) // Imprime o novo valor da variável x
	fmt.Println("Valor de y:", y)      // y ainda mantém o valor original, pois foi copiado antes da modificação

	// Ponteiro é uma referência para o endereço de memória de uma variável.

	var variavel1 int
	var ponteiro1 *int // Declaração de um ponteiro para um inteiro

	fmt.Println("Valor de variavel1:", variavel1) // Imprime o valor da variável variavel1 -> 0
	fmt.Println("Valor do ponteiro1:", ponteiro1) // Imprime o valor do ponteiro ponteiro1 -> <nil>

	fmt.Println("Endereço de memória de variavel1:", &variavel1) // Imprime o endereço de memória da variável variavel1 ->  0xc00000a100
	fmt.Println("Endereço de memória do ponteiro1:", &ponteiro1) // Imprime o endereço de memória do ponteiro ponteiro1 -> 0xc00000a108

	variavel1 = 10
	ponteiro1 = &variavel1 // Atribui o endereço de memória da variável variavel1 ao ponteiro ponteiro1

	fmt.Println("Valor de variavel1 após atribuição:", variavel1)  // Imprime o valor da variável variavel1 -> 10
	fmt.Println("Valor do ponteiro1 após atribuição:", *ponteiro1) // Imprime o valor do ponteiro ponteiro1 -> 10
	// Se eu usar o operador * antes do ponteiro, eu consigo acessar o valor da variável que ele aponta -> DESREFERENCIAÇÃO

	variavel1 = 20

	fmt.Println("Valor do ponteiro1 após modificação da variavel1:", *ponteiro1) // Imprime o valor do ponteiro ponteiro1 -> 20
	// O ponteiro ainda aponta para o endereço de memória da variável variavel1, então se eu modificar o valor da variável, o ponteiro também reflete essa mudança

}
