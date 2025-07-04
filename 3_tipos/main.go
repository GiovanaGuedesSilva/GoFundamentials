package main

import (
	"errors"
	"fmt"
)

func main() {

	// var pode ser substituído por := (inferência de tipo)

	//		INT	(valor inicial é 0)

	// int8, int16, int32, int64 := 10, 20, 30, 40 // Declaração de variáveis do tipo inteiro com diferentes tamanhos
	// int sozinho depende do sistema operacional, no caso de 64 bits, é int64

	var numero int64 = -10000
	fmt.Println("Número:", numero)

	// é do tipo int
	char := 'A'                     // Declaração de variável do tipo rune (representa um único caractere Unicode)
	fmt.Println("Caractere:", char) // 'A' é o código Unicode do caractere A, que é 65 em decimal
	// Considera a tabela ASCII, onde cada caractere tem um valor numérico associado

	//		UINT (valor inicial é 0)

	var numero2 uint8 = 255 // Declaração de variável do tipo inteiro sem sinal (não aceita números negativos)
	fmt.Println("Número 2:", numero2)

	//		RUNE (valor inicial é 0)

	// alias (é uma forma de renomear um tipo, mas não é recomendado)

	// INT32 = RUNE
	var numero3 rune = 123456 // Declaração de variável do tipo inteiro com sinal
	// (aceita números negativos e positivos, mas é usado para representar caracteres Unicode)
	fmt.Println("Número 3:", numero3)

	//		BYTE (valor inicial é 0)

	// BYTE = uint8
	var numero4 byte = 123
	fmt.Println("Número 4:", numero4)

	//		FLOAT (valor inicial é 0.0)

	// float32, float64 (são usados para representar números reais, com ponto flutuante)

	var numeroReal = 123.456 // Declaração de variável do tipo float64 (padrão do Go)
	fmt.Println("Número Real:", numeroReal)

	var numeroReal2 float32 = 123.456 // Declaração de variável do tipo float32 (menos preciso que o float64)
	fmt.Println("Número Real 2:", numeroReal2)

	numeroReal3 := 123.456 // Declaração de variável do tipo float64 (padrão do Go, mas depende do sistema operacional)
	fmt.Println("Número Real 3:", numeroReal3)

	//		STRING (valor inicial é uma string vazia "")

	var texto string = "Texto" // Declaração de variável do tipo string (sequência de caracteres)
	fmt.Println("Texto:", texto)

	text2 := "Texto 2" // Declaração de variável do tipo string (sequência de caracteres, inferência de tipo)
	fmt.Println("Texto 2:", text2)

	//		BOOLEAN (valor inicial é false)

	var booleano bool = true // Declaração de variável do tipo booleano (verdadeiro ou falso)
	fmt.Println("Booleano:", booleano)

	//		ERROS (valor inicial é nil)

	var erro error = errors.New("Erro de exemplo") // Declaração de variável do tipo erro (interface que representa um erro)
	if erro != nil {
		fmt.Println("Erro:", erro)
	} else {
		fmt.Println("Sem erros")
	}

	// erro não é um tipo primitivo, mas é uma interface que pode ser implementada por outros tipos
}
