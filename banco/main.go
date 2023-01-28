package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	fmt.Println("Banco CenaBank")

	contaDoBruno := ContaCorrente{titular: "Bruno", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	fmt.Println(contaDoBruno)

}
