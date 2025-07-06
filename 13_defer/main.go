package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func funcao3() {
	fmt.Println("Executando a função 3")
}

func alunoEstaAprovado(nota1, nota2 float64) bool {
	fmt.Println("Calculando a média do aluno")
	media := (nota1 + nota2) / 2
	defer fmt.Println("Média calculada:", media)

	// Para não repetir o código, podemos usar o DEFER
	// if media >= 7 {
	// 	fmt.Println("Média calculada:", media)
	// 	return true
	// }else{
	//  fmt.Println("Média calculada:", media)
	//  return false
	// }

	if media >= 7 {
		return true
	} else {
		return false
	}
}

func main() {

	// Defer é usado para adiar a execução de uma função até que a função que a contém retorne.
	// É útil para garantir que certas ações sejam executadas, como liberar recursos ou fechar conexões.
	defer fmt.Println("Iniciando a função main")

	defer funcao1() // defer executa a função após o término da função main
	funcao2()
	defer funcao3()

	fmt.Println(alunoEstaAprovado(7, 8))

	// A saída será:
	// Executando a função 2
	// Calculando a média do aluno
	// Média calculada: 7.5
	// true
	// Executando a função 3
	// Executando a função 1
	// Iniciando a função main
}
