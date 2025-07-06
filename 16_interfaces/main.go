package main

import "fmt"

type forma interface {
	area() float64
}

type retangulo struct {
	base   float64
	altura float64
}

func (r retangulo) area() float64 {
	// O método area calcula a área do retângulo multiplicando a base pela altura.
	return r.base * r.altura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// O método area calcula a área do círculo usando a fórmula πr².
	return 3.14 * c.raio * c.raio
}

func escreverArea(f forma) {
	// A função escreverArea recebe uma interface forma como parâmetro
	// e chama o método area da forma passada.
	// Isso permite que diferentes formas implementem o método area de maneiras diferentes.
	fmt.Println("A área da forma é: ", f.area())
}

func main() {
	r := retangulo{base: 5, altura: 10}
	c := circulo{raio: 7}

	escreverArea(r) // Chama a função escreverArea passando um retângulo
	escreverArea(c) // Chama a função escreverArea passando um círculo
}
