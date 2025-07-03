package auxiliar

import "fmt"

// Escrever é uma função pública que pode ser chamada de outros pacotes
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() // Chama a função privada do mesmo pacote
}

/*
	EscreverMensagem é uma função pública que pode ser chamada de outros pacotes
	Como Go não é um linguagem orientada a objetos, não há encapsulamento
	Portanto, todas as funções e variáveis que começam com letra maiúscula são públicas

	Por padrão, funções públicas tem um comentário explicativo
	que deve ser escrito logo acima da função, começando com o nome da função
*/
