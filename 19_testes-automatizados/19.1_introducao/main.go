package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua das Flores, 123")
	fmt.Println("Tipo de Endereço:", tipoEndereco)
}
