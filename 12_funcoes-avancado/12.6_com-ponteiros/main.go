package main

import "fmt"

func inverterSinal(numero int) int { // parâmetro por valor cópia
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { // O * indica que estamos usando um ponteiro
	// O ponteiro permite modificar o valor original da variável passada como argumento
	*numero = *numero * -1 // desferenciação do ponteiro
	// O operador * é usado para acessar o valor apontado pelo ponteiro
}

func main() {

	numero := 10
	invertido := inverterSinal(numero)
	fmt.Println("Número original:", numero)
	fmt.Println("Número invertido:", invertido)
	fmt.Println("Número original pós inversão:", numero)

	numero = 20
	invertido = inverterSinal(numero)
	fmt.Println("____________________________________")
	fmt.Println("Número original:", numero)
	fmt.Println("Número invertido:", invertido)
	fmt.Println("Número original pós inversão:", numero)

	numeroPonteiro := 30
	inverterSinalComPonteiro(&numeroPonteiro) // Passando o endereço
	// O operador & é usado para obter o endereço de memória da variável
	fmt.Println("____________________________________")
	fmt.Println("Número original:", 30)
	fmt.Println("Número invertido:", numeroPonteiro)
	fmt.Println("Número original pós inversão:", numeroPonteiro)

	// A variável orifinal não é modificada quando passamos por valor, mas é modificada quando passamos por referência (ponteiro).

}
