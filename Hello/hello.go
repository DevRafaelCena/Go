package main

import (
	"fmt"
)

func main() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá sr, ", nome)
	fmt.Println("Este programa está na versão ", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int

	fmt.Scanf("%d", &comando)
	fmt.Println("O endereço de memória de comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	// outra forma
	var comando2 int
	fmt.Scan(&comando2)
	fmt.Println("O endereço de memória de comando2 é ", &comando2)
	fmt.Println("O comando escolhido foi ", comando2)
}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
