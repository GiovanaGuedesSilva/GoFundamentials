package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 500)
		canal1 <- "Mensagem do canal 1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		canal2 <- "Mensagem do canal 2"
	}()

	for {
		select {
		case msg := <-canal1:
			fmt.Println("Recebido:", msg)
		case msg := <-canal2:
			fmt.Println("Recebido:", msg)
		}

		// Assim é mais lento, pois espera 1 segundo para receber a mensagem do canal 2
		// Assim é mais lento, pois espera 1 segundo para receber a mensagem do canal 2
		// mensagemcanal1 := <-canal1
		// fmt.Println("Recebido:", mensagemcanal1)
		// mensagemcanal2 := <-canal2
		// fmt.Println("Recebido:", mensagemcanal2)
	}

}
