package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Usuario struct {
	Nome  string `json:"nome"` // Se o valor for "-", o campo não será serializado para JSON
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

func main() {
	usuario := Usuario{
		Nome:  "João",
		Email: "joao@example.com",
		Idade: 30,
	}

	// Convertendo struct para JSON
	usuarioJSON, err := json.Marshal(usuario) // usuarioJSON é um slice de bytes
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usuarioJSON))
	fmt.Println(bytes.NewBuffer(usuarioJSON).String()) // Convertendo slice de bytes para string

	// Convertendo JSON para struct
	var usuario2 Usuario
	usuarioJSON2 := `{"nome":"Maria","email":"maria@example.com","idade":25}`
	err = json.Unmarshal([]byte(usuarioJSON2), &usuario2) // p1: dados que serão deserializados, p2: endereço de memória onde os dados serão armazenados
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usuario2)

	// Convertendo JSON para map[string]string
	// Isso é útil quando você não tem uma struct definida ou quando os dados são dinâmicos.
	// Aqui, o JSON é convertido para um mapa onde as chaves são strings e os valores também são strings.
	usuario3emJSON := `{"nome":"Pedro","email":"pedro@example.com","idade":28}`
	usuario3 := make(map[string]string)

	if err := json.Unmarshal([]byte(usuario3emJSON), &usuario3); err != nil {
		log.Fatal(err)
	}
	fmt.Println(usuario3)
}
