package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo
	// Concorrência é quando várias tarefas estão progredindo ao mesmo tempo, mas não necessariamente executando ao mesmo tempo.
	// Paralelismo é quando várias tarefas estão sendo executadas ao mesmo tempo, em diferentes núcleos de CPU.
	// Goroutines são funções que podem ser executadas concorrentemente com outras funções.

	go escrever("Olá, Mundo!")
	escrever("Programando em Go!") // A palavra-chave 'go' inicia uma nova goroutine, permitindo que a função escrever execute de forma concorrente.
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second) // Pausa de 1 segundo entre as mensagens
		// A função escrever continua executando indefinidamente, imprimindo o texto a cada segundo.
		// Isso demonstra como uma goroutine pode continuar a executar enquanto outras tarefas estão em andamento
	}
}
