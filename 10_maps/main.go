package main

import "fmt"

func main() {
	// Mapas são estruturas de dados que armazenam pares chave-valor, permitindo acesso rápido aos valores por meio de suas chaves.
	// Eles são semelhantes a dicionários ou tabelas hash em outras linguagens de programação.

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	// Acessando valores do mapa
	fmt.Println("Nome:", usuario["nome"]) // Diferente do struct, não é necessário usar o ponto para acessar os valores
	fmt.Println("Sobrenome:", usuario["sobrenome"])
	fmt.Println("Idade:", usuario["idade"])

	usuario2 := map[string]map[string]string{
		"usuario1": {
			"nome":      "Maria",
			"sobrenome": "Oliveira",
			"idade":     "25",
		},
		"curso1": {
			"nome":      "Matemática",
			"professor": "Dr. Silva",
			"duracao":   "60 horas",
		},
	}

	// Acessando valores do mapa aninhado
	fmt.Println("Nome do usuário 1:", usuario2["usuario1"]["nome"])
	fmt.Println("Sobrenome do usuário 1:", usuario2["usuario1"]["sobrenome"])
	fmt.Println("Idade do usuário 1:", usuario2["usuario1"]["idade"])
	fmt.Println("Nome do curso 1:", usuario2["curso1"]["nome"])
	fmt.Println("Professor do curso 1:", usuario2["curso1"]["professor"])
	fmt.Println("Duração do curso 1:", usuario2["curso1"]["duracao"])

	delete(usuario2, "idade") // Remove a chave "idade" do mapa usuario
	fmt.Println("Usuário após remoção da idade:", usuario2)

	usuario2["signo"] = map[string]string{
		"nome":     "Áries",
		"elemento": "Fogo",
		"planeta":  "Marte",
	}

	fmt.Println("Usuário após adição do signo:", usuario2)
}
