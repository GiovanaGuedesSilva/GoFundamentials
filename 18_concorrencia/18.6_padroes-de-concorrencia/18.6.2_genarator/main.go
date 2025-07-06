package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá, Mundo!") // Chama a função escrever que retorna um canal

	for i := 0; i < 5; i++ { // Limita a 5 mensagens para evitar um loop infinito
		fmt.Println(<-canal) // Recebe mensagens do canal e imprime
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	// O padrão Generator é uma técnica de concorrência onde uma função produz valores que podem ser consumidos por outras partes do programa.
	// Ele é útil para gerar sequências de valores de forma assíncrona, permitindo que o programa continue executando enquanto os valores são gerados.
	// Aqui, a função escrever cria um canal e inicia uma goroutine que envia mensagens para o canal indefinidamente.
	// Isso permite que outras partes do programa possam receber essas mensagens de forma assíncrona

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500) // Pausa de 500 milissegundos entre as mensagens
		}
	}()

	return canal
}
