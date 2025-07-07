package enderecos

// TESTE DE UNIDADE

// Em go, os testes de unidade são escritos em arquivos separados com o sufixo _test.go.
// para rodar os testes, você pode usar o comando `go test` no terminal.

import (
	"testing" // Importa o pacote de testes do Go
)

func TestTipoDeEndereco(t *testing.T) { // Começa com "Test" para ser reconhecido como um teste
	// Testa se a função TipoDeEndereco retorna o tipo correto de endereço
	resultado := TipoDeEndereco("Rua das Flores, 123")
	esperado := "Rua"

	if resultado != esperado {
		t.Errorf("TipoDeEndereco recebeu %q e retornou %q, mas o esperado era %q", "Rua das Flores, 123", resultado, esperado)
		// .error é usado para registrar um erro de teste, mas não interrompe a execução do teste.
		// .Errorf formata a mensagem de erro com os valores fornecidos.
	}
}
