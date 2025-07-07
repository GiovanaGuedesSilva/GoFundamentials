## Estrutura de Testes em Go

- Os arquivos de teste devem ter o sufixo `_test.go`.
- Funções de teste devem começar com `Test` e receber um parâmetro `*testing.T`.
- Exemplo básico de teste:

```go
func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    esperado := 5

    if resultado != esperado {
        t.Errorf("esperado %d, mas obteve %d", esperado, resultado)
    }
}
```

## Comandos Úteis para Testes

- `go test`  
  Executa todos os testes do pacote atual.

- `go test ./...`  
  Executa todos os testes do pacote atual e de todos os subpacotes.

- `go test -v`  
  Executa os testes mostrando detalhes (verbose) de cada teste.

- `go test -run TestNomeDoTeste`  
  Executa apenas o teste especificado pelo nome.

- `go test -run TestNomeDoTeste -v`  
  Executa o teste especificado e mostra detalhes.

- `go test -run TestNomeDoTeste -v -count=1`  
  Executa o teste especificado, mostra detalhes e desativa o cache de testes (útil para garantir execução do zero).

- `go test --cover`  
  Executa os testes e mostra um relatório de cobertura de código.

- `go test --coverprofile cobertura.txt`  
  Executa os testes e gera um arquivo de perfil de cobertura chamado `cobertura.txt`.

- `go tool cover --func=cobertura.txt`  
  Analisa o arquivo de cobertura e mostra um resumo das funções cobertas.

- `go tool cover --html=cobertura.txt`  
  Gera um relatório de cobertura em HTML para visualização gráfica.

> **Dica:** O Go utiliza cache para acelerar a execução dos testes. Use a flag `-count=1` para forçar a execução dos testes do zero.

## Boas Práticas

- Escreva testes pequenos e focados em um único comportamento.
- Use tabelas de teste para cobrir múltiplos cenários.
- Utilize funções auxiliares para evitar repetição de código.
- Sempre rode os testes antes de commitar alterações.

## Recursos Úteis

- [Documentação oficial de testes em Go](https://golang.org/pkg/testing/)
- [Go by Example: Testing](https://gobyexample.com/testing)
- [Go Tools: cover](https://golang.org/cmd/cover/)
