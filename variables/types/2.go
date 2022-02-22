package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{
		Nome:  "Thiago",
		Idade: 44,
	}

	fmt.Println(pessoa.Idade)

}
