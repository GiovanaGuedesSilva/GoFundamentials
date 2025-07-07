package enderecos

import "strings"

// TipoDeEndereco verifica se o tipo de endereço é válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Travessa", "Largo", "Praça", "Alameda"}
	pimeiraPalavraDoEndereco := strings.Split(endereco, " ")[0] // Divide a string em um slice de strings usando o espaço como delimitador

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if strings.EqualFold(tipo, pimeiraPalavraDoEndereco) { // Compara as strings ignorando diferenças de maiúsculas e minúsculas
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return pimeiraPalavraDoEndereco
	}

	return "Tipo de endereço desconhecido"
}
