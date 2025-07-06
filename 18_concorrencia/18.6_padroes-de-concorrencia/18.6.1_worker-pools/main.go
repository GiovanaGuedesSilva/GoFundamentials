package main

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < cap(tarefas); i++ {
		tarefas <- i // Envia a posição do Fibonacci para o canal de tarefas
	}

	close(tarefas) // Fecha o canal de tarefas para indicar que não há mais tarefas a serem enviadas

	for i := 0; i < cap(resultados); i++ {
		resultado := <-resultados                  // Recebe o resultado do canal de resultados
		println("Fibonacci de", i, "é", resultado) // Imprime o resultado do Fibonacci
	}
}

// <-chan int:  // Canal de entrada, só pode receber valores
// chan<- int:  // Canal de saída, só pode enviar valores

func worker(tarefas <-chan int, resultados chan<- int) {
	for posicao := range tarefas {
		resultados <- fibonacci(posicao)
	}
}

// O padrão Worker Pool é uma técnica de concorrência onde um grupo de goroutines (workers) processa tarefas de um canal de entrada (tarefas)
// e envia os resultados para um canal de saída (resultados).

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return int(posicao)
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
