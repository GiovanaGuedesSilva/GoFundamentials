package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(
		escrever("Olá, Mundo!"),
		escrever("Olá, Go!"),
		escrever("Programando em Go!"),
	)

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

// O padrão Multiplexador é uma técnica de concorrência onde várias goroutines enviam mensagens para um único canal de saída.
// Ele é útil para combinar mensagens de diferentes fontes em um único canal, permitindo que o consumidor
// receba mensagens de forma assíncrona e sem se preocupar com a origem das mensagens.
func multiplexar(canaisDeEntrada ...<-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canaisDeEntrada[0]:
				canalDeSaida <- fmt.Sprintf("Canal 1: %s", msg)
			case msg := <-canaisDeEntrada[1]:
				canalDeSaida <- fmt.Sprintf("Canal 2: %s", msg)
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500) // Pausa de 500 milissegundos entre as mensagens
		}
	}()

	return canal
}
