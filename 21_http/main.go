package main

import (
	"log"
	"net/http"
)

func main() {
	// Http é um protocolo de comunicação usado para transferir dados na web.
	// Ele é baseado no modelo cliente-servidor, onde o cliente faz requisições e o servidor responde.
	// O pacote "net/http" do Go fornece uma implementação do protocolo HTTP, permitindo criar servidores e clientes HTTP.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página raíz!"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// HandleFunc é uma função que associa um caminho de URL a uma função manipuladora
		// O primeiro parâmetro é o caminho da URL
		// O segundo parâmetro é a função manipuladora que será chamada quando a URL for acessada
		w.Write([]byte("Bem-vindo à página inicial!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil)) // Inicia um servidor HTTP na porta 5000
	// http.ListenAndServe é uma função que inicia um servidor HTTP na porta especificada.
	// O primeiro parâmetro é a porta onde o servidor irá escutar as requisições.
	// O segundo parâmetro é um manipulador de requisições (handler) que define como as requisições serão tratadas.
	// nil significa que o servidor usará o manipulador padrão, que é capaz de lidar com requisições HTTP básicas.
	// log.Fatal é usado para registrar um erro fatal e encerrar o programa se ListenAndServe retornar um erro.
	// Se o servidor não puder ser iniciado, log.Fatal irá registrar o erro e encerrar o programa.
	// Isso é útil para garantir que o servidor seja iniciado corretamente e que qualquer erro seja tratado.

}
