package main

import "fmt"

// Posso utilizar para defini um objeto
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{
		Nome:  "Thiago",
		Idade: 14,
	}
	fmt.Println(pessoa.Nome)

}
