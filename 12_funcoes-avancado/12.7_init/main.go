package main

import "fmt"

var n int

func init() {
	// A função init é executada antes da função main, mesmo que esteja definida após ela.
	// É útil para inicializar variáveis, configurar o ambiente ou executar código de configuração.
	// Pode ter uma função init por arquivo, e não pode receber parâmetros nem retornar valores.

	fmt.Println("Função init sendo executada")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println("Valor de n:", n)
}
