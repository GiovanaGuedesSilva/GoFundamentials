package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Se você quiser que o programa principal espere as goroutines terminarem, você pode usar um WaitGroup.

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Adiciona duas goroutines ao WaitGroup

	go func() {
		escrever("Olá, Mundo!")
		waitGroup.Done() // Indica que a goroutine terminou	 	-1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // Indica que a goroutine terminou	 	-1
	}()

	waitGroup.Wait() // Espera todas as goroutines terminarem	//0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
