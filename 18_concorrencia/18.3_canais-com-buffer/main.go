package main

import "fmt"

func main() {
	// canal := make(chan string)

	// Aqui estou enviando um valor mas não tem ninguém recebendo, então o programa vai dar deadlock
	// Para iss, se deve usar um canal com buffer, que permite enviar valores sem precisar de um receptor imediato

	canal := make(chan string, 2) // Cria um canal com buffer de tamanho 1

	canal <- "Olá, Mundo!"        // Envia uma mensagem para o canal
	canal <- "Programando em Go!" // Envia outra mensagem para o canal

	// Se eu adicionar mais mensagens do que o buffer permite, o programa vai dar deadlock
	// canal <- "Mais uma mensagem" // Isso causaria um deadlock, pois o canal já está cheio

	mensagem := <-canal  // Recebe a mensagem do canal
	mensagem2 := <-canal // Recebe a segunda mensagem do canal

	fmt.Println(mensagem)  // Imprime a mensagem recebida do canal
	fmt.Println(mensagem2) // Imprime a segunda mensagem recebida do canal
}
