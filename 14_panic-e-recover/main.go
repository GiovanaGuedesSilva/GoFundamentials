package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução após um pânico")
	if r := recover(); r != nil { // A função recover captura o pânico
		// Se recover retornar um valor diferente de nil, significa que um pânico ocorreu
		fmt.Println("Recuperado do pânico:", r)
	} else {
		fmt.Println("Nenhum pânico ocorreu")
	}

	// A função recover é usada para recuperar a execução de um pânico.
	// Ela deve ser chamada dentro de uma função defer para capturar o pânico e evitar a interrupção do programa.
	// Se recover for chamado fora de uma função defer, ele não terá efeito e o pânico continuará.
}

func alunoEstaAprovado(nota1, nota2 float64) bool {
	defer recuperarExecucao() // Chama a função de recuperação após um pânico
	media := (nota1 + nota2) / 2

	// A média nunca pode ser 6

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// A diferença entre error e panic é que o panic interrompe a execução do programa, enquanto o error pode ser tratado de forma controlada.

	panic("A média não pode ser 6") // Gera um pânico se a média for exatamente 6
	// panic mata a execução do programa e imprime a mensagem de erro
	// O programa não continua após o pânico, a menos que seja tratado com recover
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6)) // panic: A média não pode ser 6
	fmt.Println("Pós execução da função alunoEstaAprovado")
}
