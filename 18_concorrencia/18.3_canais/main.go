package main

import (
	"fmt"
	"time"
)

func main() {
	// Canais são estruturas que permitem a comunicação entre goroutines.
	// Eles permitem que uma goroutine envie dados para outra, sincronizando a execução entre elas.

	canal := make(chan string)
	go escrever("Olá, Mundo!", canal)

	fmt.Println("Depois da função escrever") // Esta linha será executada imediatamente após a criação da goroutine, sem esperar que ela termine.

	//<-canal // Espera até que a primeira mensagem seja recebida no canal
	// A partir daqui, o programa continuará recebendo mensagens do canal até que todas as goroutines tenham terminado de enviar mensagens.

	//mensagem := <-canal // Recebe a primeira mensagem do canal -> Indica que o canal espera receber um valor, que será passado para mensagem
	// A linha acima bloqueia a execução até que uma mensagem seja recebida no canal.
	// Se não houver mensagens disponíveis, o programa ficará esperando indefinidamente.
	//fmt.Println(mensagem)

	for {
		// Assim vai dar deadlock, pois a goroutine nunca vai terminar de enviar mensagens
		// Se você quiser receber mensagens indefinidamente, pode usar um loop para receber mensagens
		mensagem, aberto := <-canal // ele espera uma mensagem que nunca vai receber, pois a goroutine nunca vai terminar de enviar mensagens
		if !aberto {
			break
		}
		fmt.Println(mensagem)

	}

	// Esse for pode substituir o loop anterior

	for mensagem := range canal { // Recebe mensagens do canal até que ele seja fechado
		fmt.Println(mensagem) // Imprime cada mensagem recebida do canal
		// A linha acima imprime cada mensagem recebida do canal até que o canal seja fechado.
	}

	// Canais podem enviar ou receber um dado

	fmt.Println("Canal fechado, programa terminado") // Esta linha será executada após o fechamento do canal, indicando que não haverá mais mensagens a serem recebidas.
	// O programa termina aqui, pois todas as mensagens foram recebidas e o canal foi fechado
}

func escrever(texto string, canal chan string) {
	time.Sleep(5 * time.Second) // Significa que a goroutine vai esperar 5 segundos antes de enviar a mensagem para o canal
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal após enviar todas as mensagens
	// Fechar o canal é importante para indicar que não haverá mais mensagens enviadas/recebidas.
}
