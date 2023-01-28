package main

import "fmt"

func main() {
	var nome string = "Rafael"
	var idade int = 32
	var versao float32 = 1.1
	fmt.Println("Olá sr, ", nome, "sua idade é ", idade)
	fmt.Println("Este programa esta na versão ", versao)
}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
