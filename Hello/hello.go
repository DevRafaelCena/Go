package main

import (
	"fmt"
)

func main() {
	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}

}

func exibeIntroducao() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá sr, ", nome)
	fmt.Println("Este programa está na versão ", versao)
}

func leComando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereço de memória de comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	return comando
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
