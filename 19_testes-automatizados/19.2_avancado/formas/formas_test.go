package formas_test

import (
	"math"
	. "testes-avancados/formas"
	"testing"
)

func TestArea(t *testing.T) {
	// TDD - Test Driven Development

	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{Base: 5, Altura: 10}
		areaEsperada := float64(50)
		areaRecebida := retangulo.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("Área do retângulo esperada %f, mas foi %f", areaEsperada, areaRecebida)
			// O método Fatalf falha o teste e imprime a mensagem formatada.
			// Ele é usado quando a condição de teste não é atendida, interrompendo a execução do teste.
			// É diferente do método Errorf, que apenas registra o erro mas continua a execução do teste.
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{Raio: 7}
		areaEsperada := float64(math.Pi * 7 * 7)
		areaRecebida := circulo.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("Área do círculo esperada %f, mas foi %f", areaEsperada, areaRecebida)
		}
	})
}
