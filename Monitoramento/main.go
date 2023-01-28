package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

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

	sites := leArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site: ", site)
			resp, _ := http.Get(site)

			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso!")
				registraLog(site, true)
			} else {
				fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
				registraLog(site, false)
			}

			time.Sleep(delay * time.Second)

		}
	}

}

func leArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	//fmt.Println(string(arquivo))

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace((linha[:len(linha)-1]))
		sites = append(sites, linha)

		if err != nil {
			break
		}
	}

	return sites
}

// restante do código omitido

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

/* ## Compilando o programa
## go build hello.go
## ./hello
## copilando e rodando
# go run hello.go
*/
