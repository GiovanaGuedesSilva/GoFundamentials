package main

import "fmt"

// Exemplo de definição de uma estrutura
type pessoa struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	cidade string
	estado string
	cep    string
}

func main() {

	// Structs são tipos de dados compostos que permitem agrupar diferentes tipos de dados em uma única entidade.
	// Elas são úteis para organizar dados relacionados e criar modelos mais complexos.

	// Criando uma instância da estrutura Pessoa
	p := pessoa{
		nome:  "João",
		idade: 30,
		endereco: endereco{
			cidade: "São Paulo",
			estado: "SP",
			cep:    "12345-678",
		},
	}
	fmt.Println("Nome:", p.nome)
	fmt.Println("Idade:", p.idade)
	fmt.Println("Endereço:", p.endereco)

	p2 := pessoa{}  // Inicialização de uma struct sem valores -> Valores padrão
	fmt.Println(p2) //Vai imprimir os valores padrão dos campos da struct, que são: nome="", idade=0, endereco=""

	p3 := pessoa{"Maria", 25, endereco{cidade: "Bragança", cep: "10183-23"}} // Inicialização direta com valores -> Inferencia de tipo
	fmt.Println(p3)

	p4 := pessoa{nome: "Maria"} // Inicialização direta com valores nulos
	fmt.Println(p4)             // Vai imprimir os valores padrão dos campos da struct, exceto o nome que será "Maria"

}
