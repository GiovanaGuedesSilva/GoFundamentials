package main

import (
	"time"
)

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func diaDaSemana2(dia int) string {

	var diaDaSemana string

	switch {
	case dia == 1:
		diaDaSemana = "Domingo"
		fallthrough // O fallthrough permite que o controle continue para o próximo case, mesmo que a condição não seja atendida
		// Isso é útil quando você deseja que várias condições compartilhem o mesmo bloco de código
		// se cair aqui, o controle continuará para o próximo case
	case dia == 2:
		diaDaSemana = "Segunda-feira"
	case dia == 3:
		diaDaSemana = "Terça-feira"
	case dia == 4:
		diaDaSemana = "Quarta-feira"
	case dia == 5:
		diaDaSemana = "Quinta-feira"
	case dia == 6:
		diaDaSemana = "Sexta-feira"
	case dia == 7:
		diaDaSemana = "Sábado"
	default:
		return "Dia inválido"
	}

	return diaDaSemana
}

func main() {
	// Estruturas de controle são usadas para controlar o fluxo de execução do programa.
	// Elas permitem que você execute diferentes blocos de código com base em condições específicas.

	// 		CONDICIONAIS

	numero := 10
	if numero > 0 { // As chaves são obrigatórias em estruturas de controle em Go
		println("O número é positivo")
	} else if numero < 0 {
		println("O número é negativo")
	} else {
		println("O número é zero")
	}

	if numero2 := numero + 5; numero2 > 10 { // Declaração de variável dentro do escopo do if
		println("Número2 é maior que 10:", numero2) // Não é possível acessar a variável numero2 fora do escopo do if
	}

	println("Dia da semana:", diaDaSemana(3))    // Chamada da função diaDaSemana
	println("Dia da semana 2:", diaDaSemana2(8)) // Chamada da função diaDaSemana2

	// 		LAÇOS DE REPETIÇÃO

	i := 0

	for i < 5 { // Laço de repetição for com condição -> semelhante ao while
		i++
		println("Valor de i:", i)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	}

	for j := 0; j < 5; j += 2 { // Laço de repetição for com inicialização, condição e incremento
		println("Valor de j:", j)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	} // j é invisível fora do escopo do for

	nomes := []string{"João", "Maria", "José"}

	// index representa o índice do elemento na coleção, e nome representa o valor do elemento
	for index, nome := range nomes { // Laço de repetição for com range -> percorre elementos de uma coleção
		println("Índice:", index, "Nome:", nome)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	}

	for _, nome := range nomes { // Laço de repetição for com range, mas ignorando o índice
		println("Nome:", nome)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	}

	for indice, letra := range "Go" { // Laço de repetição for com range em uma string
		println("Índice:", indice, "Letra:", string(letra)) // Convertendo o rune para string
		// Se eu passar apenas a letra, ela será exibida como um número (código Unicode)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	for chave, valor := range usuario {
		println("Chave:", chave, "Valor:", valor)
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
	}

	// Não tem como fazer um laço de repetição com range em um struct, pois ele não é uma coleção iterável.
	// Para iterar sobre os campos de um struct, você pode usar a reflexão (reflection) em Go, mas isso é mais avançado e não é comum em situações simples.

	for {
		println("Este laço de repetição é infinito")
		time.Sleep(time.Second) // Pausa a execução por 1 segundo
		// Para sair do laço, você pode usar uma condição de parada ou um comando break.
		// Exemplo: if i > 10 { break }
	}
}
