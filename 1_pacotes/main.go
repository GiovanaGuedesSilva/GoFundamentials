package main

import (
	"fmt"
	"modulo/auxiliar" // Importa o pacote auxiliar

	"github.com/badoux/checkmail"
	// Importa o pacote chackmail do GitHub
)

func main() {

	/*
		Este é o ponto de entrada do programa em que é possível chamar funções de outros pacotes ou executar a lógica principal do seu aplicativo
	*/

	fmt.Println("Bem-vindo ao GoFundamentals!")
	auxiliar.Escrever() // Chama a função Escrever do pacote auxiliar

	// Chama a função ValidateFormat do pacote checkmail para validar o formato de um e-mail
	erro := checkmail.ValidateFormat("giovana@gmail.com")
	fmt.Println(erro)
}

/*
	go run main.go

	ou

	go build (gera um executável)	Não atualiza em tempo real, então é necessário executar o comando de novo para pegar as alterações
	./modulo (executa o programa)

	ou

	go install (instala o pacote no diretório $GOPATH/bin, tornando-o disponível globalmente)
*/
