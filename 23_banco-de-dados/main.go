package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import implícito do driver MySQL para Go (a conexão vai precisar dele)
)

func main() {

	urlConexao := "root:1984@/devbook?charset=utf8&parseTime=True&loc=Local" // URL de conexão com o banco de dados
	// O primeiro parâmetro é o usuário, o segundo é a senha e o terceiro é o nome do banco de dados
	// Charset=utf8 define o conjunto de caracteres a ser usado
	// ParseTime=True permite que o driver converta os tipos de data e hora do banco de dados para tipos Go
	// Loc=Local define o fuso horário local a ser usado

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		log.Fatal(erro) // Se houver um erro ao abrir a conexão, o programa será encerrado com uma mensagem de erro
	}
	defer db.Close() // Garante que a conexão com o banco de dados será fechada ao final do programa

	if erro := db.Ping(); erro != nil {
		log.Fatal(erro) // Se houver um erro ao pingar o banco de dados, o programa será encerrado com uma mensagem de erro
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!") // Mensagem

	linhas, erros := db.Query("SELECT * FROM usuarios")
	if erros != nil {
		log.Fatal(erros) // Se houver um erro ao executar a consulta, o programa será encerrado com uma mensagem de erro
	}
	defer linhas.Close() // Garante que as linhas retornadas pela consulta serão fechadas ao final do programa

	fmt.Println(linhas) // Imprime as linhas retornadas pela consulta
}
