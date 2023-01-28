package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < len(sites); i++ {
		fmt.Println("Testando site: ", sites[i])
		resp, _ := http.Get(sites[i])

		if resp.StatusCode == 200 {
			fmt.Println("O endereço respondeu corretamente ", resp.StatusCode)
		} else {
			fmt.Println("O endereço ", sites[i], "está com problemas. Status Code: ", resp.StatusCode)
		}
	}

}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
