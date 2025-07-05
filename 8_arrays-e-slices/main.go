package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays são estruturas de dados que armazenam uma coleção de elementos do mesmo tipo.
	// Slices são abstrações mais flexíveis sobre arrays, permitindo redimensionamento dinâmico.

	//		ARRAY

	var array1 [5]string // Declaração de um array de inteiros com tamanho fixo de 5
	// Se eu não colocar um tipo, o Go assume que é um slice de strings
	fmt.Println("Array 1:", array1)
	// Imprime o array com valores padrão ("" para string)

	array1[0] = "Posição 1"                                                             // Atribui um valor à primeira posição do array
	array1 = [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"} // Atribui valores ao array
	fmt.Println("Array 1 após atribuição:", array1)

	array2 := [...]int{1, 2, 3, 4, 5} // Declaração de um array com tamanho inferido pelo compilador
	// [...] indica que o tamanho do array será determinado automaticamente com base nos valores fornecidos
	fmt.Println("Array 2:", array2)

	// SLICES

	// Um slice é uma abstração sobre um array que permite redimensionamento dinâmico e manipulação mais fácil de coleções de dados.
	// É como uma fatia de um array, mas sem tamanho fixo.

	slice1 := []int{1, 2, 3, 4, 5} // Não é necessário especificar o tamanho do slice, pois ele pode crescer dinamicamente
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1)) // Imprime o tipo do slice	(o reflect é um pacote que permite inspecionar tipos em tempo de execução)
	fmt.Println(reflect.TypeOf(array1)) // Imprime o tipo do array

	slice1 = append(slice1, 6, 7, 8) // Adiciona elementos ao slice usando a função append

	slice2 := array1[1:3] // Essa sintaxe : indica que estamos criando um slice a partir de um array, incluindo os elementos do índice 1 ao 2 (exclusivo)
	// Nesse caso, o slice2 conterá os elementos "Posição 2" e "Posição 3" do array1
	fmt.Println("Slice 2:", slice2)

	array1[1] = "Posição 2 Modificada" // Modifica o valor do array original
	fmt.Println("Slice 2 após modificação:", slice2)

	// ARRAYS INTERNOS

	slice3 := make([]float32, 5, 10) // Cria um slice com 5 elementos do tipo float32, com capacidade inicial de 10
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Comprimento de Slice 3:", len(slice3)) // Imprime o comprimento
	fmt.Println("Capacidade de Slice 3:", cap(slice3))  // Imprime a capacidade

	slice3 = append(slice3, 1.1, 2.2, 3.3) // Adiciona elementos ao slice
	fmt.Println("Slice 3 após append:", slice3)
	fmt.Println("Comprimento de Slice 3 após append:", len(slice3)) // Imprime o novo comprimento
	fmt.Println("Capacidade de Slice 3 após append:", cap(slice3))  // Imprime

	// Se eu estourar a capacidade do slice, o Go automaticamente aloca um novo array com o dobro da capacidade original
	// Se meu slice tinha capacidade 11 e eu adiciono mais um elemento, o Go aloca um novo array com capacidade 24 (12*2)

	slice4 := make([]int, 5) // Cria um slice com 5 elementos do tipo int, com capacidade inicial de 5
	// Quando eu crio um slice com make, o Go aloca um array interno para armazenar os elementos do slice
	// Se eu não especificar o tamanho, o Go assume que o tamanho é a capacidade inicial
	fmt.Println("Slice 4:", slice4)
	fmt.Println("Comprimento de Slice 4:", len(slice4)) // Imprime o comprimento "5"
	fmt.Println("Capacidade de Slice 4:", cap(slice4))  // Imprime a capacidade "5"
	slice4 = append(slice4, 6, 7, 8)                    // Adiciona elementos ao slice
	fmt.Println("Slice 4 após append:", slice4)
	fmt.Println("Comprimento de Slice 4 após append:", len(slice4)) // Imprime o comprimento "8"
	fmt.Println("Capacidade de Slice 4 após append:", cap(slice4))  // Imprime a capacidade "10"

	// Array é uma lista com tamanho fixo, enquanto slice é uma lista dinâmica sem tamanho fixo.
}
