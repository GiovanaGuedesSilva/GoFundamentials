package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Altura float32
}

type Esudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {

	// Em go, não existe herança como em outras linguagens orientadas a objetos.
	// Em vez disso, Go utiliza composição, onde uma struct pode conter outras structs como campos

	p1 := Pessoa{
		Nome:   "João",
		Idade:  30,
		Altura: 1.75,
	}
	fmt.Println(p1)

	e1 := Esudante{
		Pessoa:    p1,
		curso:     "Engenharia",
		faculdade: "Universidade XYZ",
	}
	fmt.Println(e1)

	e2 := Esudante{
		Pessoa: Pessoa{
			Nome:   "Maria",
			Idade:  25,
			Altura: 1.65,
		},
		curso:     "Medicina",
		faculdade: "Universidade ABC",
	}
	fmt.Println(e2)
}
