package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { // Todos usuários podem salvar
	fmt.Printf("Usuário %s salvo com sucesso no banco de dados!\n", u.nome)
	// %s é um verbo de formatação para strings
	// %d é um verbo de formatação para inteiros
	// %f é um verbo de formatação para floats
	// %v é um verbo de formatação para valores genéricos
	// %T é um verbo de formatação para o tipo do valor
	// %t é um verbo de formatação para booleanos
	// %p é um verbo de formatação para ponteiros
	// %q é um verbo de formatação para strings com aspas
}

func (u usuario) maiorIdade() bool {
	// Método para verificar se o usuário é maior de idade
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	// Método para incrementar a idade do usuário
	u.idade++
	fmt.Printf("Feliz aniversário, %s! Agora você tem %d anos.\n", u.nome, u.idade)
}

func main() {

	// Métodos são funções associadas a um tipo específico.
	// Eles permitem que você defina comportamentos para tipos personalizados, como structs.
	// A diferença entre funções e métodos é que métodos têm um receptor (o tipo ao qual estão associados).

	var u usuario
	u.nome = "João"
	u.idade = 30
	fmt.Println(u)
	u.salvar()

	maiorDeIdade := u.maiorIdade()
	fmt.Println("Usuário é maior de idade:", maiorDeIdade)

	u.fazerAniversario()
	fmt.Println("Usuário após o aniversário:", u)
}
