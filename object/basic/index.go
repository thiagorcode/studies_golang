package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
} // Posso utilizar para defini um objeto

func main() {
	pessoa := Pessoa{
		Nome:  "Thiago",
		Idade: 14,
	}
	fmt.Println(pessoa.Nome)

}
