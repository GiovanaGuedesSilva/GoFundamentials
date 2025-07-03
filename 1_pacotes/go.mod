module modulo

go 1.24.4

// Centraliza todos os pacotes utilizados no projeto
// para facilitar a atualização e manutenção

// go mod init modulo (cria o arquivo go.mod)
// go mod tidy (remove pacotes não utilizados e atualiza as versões)

// go get github.com/badoux/checkmail (aparece nesse arquivo automaticamente depois)
require github.com/badoux/checkmail v1.2.4
