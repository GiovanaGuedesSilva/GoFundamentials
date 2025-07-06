package main

import "fmt"

func generica(interf interface{}) {
	// A função generica recebe um parâmetro do tipo interface{}.
	// O tipo interface{} é um tipo genérico que pode receber qualquer valor.

	fmt.Println(interf)
}

func main() {
	// Interfaces são tipos que definem um conjunto de métodos que um tipo deve implementar.
	// Elas permitem a criação de código genérico e flexível, onde diferentes tipos podem ser tratados de forma uniforme.
	generica("Olá, mundo!") // Passa uma string
	generica(42)            // Passa um inteiro
	generica(3.14)          // Passa um float
	generica(true)          // Passa um booleano

	// equivalente a:

	fmt.Println("Olá, mundo!", 42, 3.14, true)

	// Não é recomendado usar interface{} para tudo, pois pode dificultar a legibilidade do código e a detecção de erros.
	// É melhor usar tipos específicos sempre que possível.
	mapa := map[interface{}]interface{}{
		"nome":   "João",
		"idade":  30,
		"altura": 1.75,
		"casado": false,
	}
	fmt.Println(mapa)
}
