package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Rafael"
	var idade int = 32
	var versao float32 = 1.1
	fmt.Println("Olá sr, ", nome, "sua idade é ", idade)
	fmt.Println("Este programa esta na versão ", versao)

	fmt.Println("O tipo da variavel nome é ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é ", reflect.TypeOf(versao))
}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
