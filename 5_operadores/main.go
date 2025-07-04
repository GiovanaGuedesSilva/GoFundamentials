package main

import "fmt"

func main() {

	//		ARITMÉTICOS

	soma := 10 + 5
	subtracao := 10 - 5
	multiplicacao := 10 * 5
	divisao := 10 / 5
	modulo := 10 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, modulo)

	var numero1 int16 = 10
	var numero2 int32 = 30
	// soma2 := numero1 + numero2 // Soma de dois números inteiros de diferentes tamanhos, dará erro de compilação
	// Para somar números de diferentes tamanhos, é necessário fazer a conversão explícita
	soma2 := numero1 + int16(numero2) // Converte numero2 para int
	fmt.Println("Soma:", soma2)

	//		ATRIBUIÇÃO

	fmt.Println("_______________________")
	var variavel1 string = "Texto"
	variavel2 := "Outro Texto" // Atribuição com inferência de tipo
	println("Variável 1:", variavel1, "Variável 2:", variavel2)

	//		RELACIONAIS

	fmt.Println("_______________________")
	fmt.Println(1 > 2) // Operadores relacionais retornam valores booleanos e usados para comparar valores e verificar condições
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2) // Igualdade
	fmt.Println(1 != 2) // Diferente

	// 		LOGICOS

	fmt.Println("_______________________")
	fmt.Println(true && false) // E lógico
	fmt.Println(true || false) // Ou lógico
	fmt.Println(!true)         // Não lógico

	//		UNÁRIOS

	fmt.Println("_______________________")
	numero := 10
	numero++ // Incremento
	fmt.Println("Número após incremento:", numero)

	numero-- // Decremento
	fmt.Println("Número após decremento:", numero)

	numero += 5 // Atribuição com adição
	fmt.Println("Número após adição:", numero)

	numero -= 3 // Atribuição com subtração
	fmt.Println("Número após subtração:", numero)

	numero *= 2 // Atribuição com multiplicação
	fmt.Println("Número após multiplicação:", numero)

	numero /= 4 // Atribuição com divisão
	fmt.Println("Número após divisão:", numero)

	numero %= 3 // Atribuição com módulo
	fmt.Println("Número após módulo:", numero)

	//		TERNÁRIOS

	fmt.Println("_______________________")

	// Go não possui operador ternário, mas podemos simular seu comportamento com uma estrutura condicional

	numero3 := 10
	numero4 := 20
	resultado := "Menor que 15"
	if numero3+numero4 < 15 {
		resultado = "Maior que 15"
	}
	fmt.Println("Resultado da verificação:", resultado)

	// Isso poderia ser represetado com operador ternário assim:
	// resultado := if numero3+numero4 < 15 ? "Menor que 15" : "Maior que 15"
}
