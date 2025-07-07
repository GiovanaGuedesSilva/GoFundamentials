package enderecos_test

// TESTE DE UNIDADE

import (
	. "introducao-testes/enderecos" // Importa o pacote que contém a função TipoDeEndereco
	"testing"                       // Importa o pacote de testes do Go
)

type cenarioDeTeste struct {
	entrada  string // Entrada do teste
	esperado string // Resultado esperado
}

func TestTipoDeEndereco(t *testing.T) { // Começa com "Test" para ser reconhecido como um teste
	// Testa se a função TipoDeEndereco retorna o tipo correto de endereço

	t.Parallel() // Marca o teste como paralelo, permitindo que ele seja executado em paralelo com outros testes

	cenarios := []cenarioDeTeste{
		{"Rua das Flores, 123", "Rua"},
		{"Avenida Paulista, 456", "Avenida"},
		{"Travessa da Alegria, 789", "Travessa"},
		{"Largo do Machado, 101", "Largo"},
		{"Praça da Sé, 202", "Praça"},
		{"Alameda Santos, 303", "Alameda"},
		{"Beco do Batman, 404", "Tipo de endereço desconhecido"}, // Cenário com tipo desconhecido
	}

	for _, cenario := range cenarios {
		resultado := TipoDeEndereco(cenario.entrada)
		if resultado != cenario.esperado {
			t.Errorf("TipoDeEndereco recebeu %q e retornou %q, mas o esperado era %q", cenario.entrada, resultado, cenario.esperado)
		}
	}
	// .Error é usado para registrar um erro de teste, mas não interrompe a execução do teste.
	// .Errorf formata a mensagem de erro com os valores fornecidos.
}

func TestQualquer(t *testing.T) {
	// Testa se a função TipoDeEndereco retorna o tipo correto de endereço
	// Esse teste é apenas um exemplo e não faz nada, mas é necessário para mostrar como escrever testes em Go.

	t.Parallel() // Marca o teste como paralelo, permitindo que ele seja executado em paralelo com outros testes

	t.Error("Esse teste falhou intencionalmente") // Força uma falha no teste
}
