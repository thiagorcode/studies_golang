package main

import (
	"fmt"
)

type Pessoa_2 struct {
	Nome  string
	Idade int
}

// Asterisco referencia a posição na memória de pessoa e altera o valor do mesmo
// caso não seja passado esee atributo o valor só é alterado dentro do escopo da func.
func (p *Pessoa_2) setNome(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome)

}

func finally() {
	pessoa := Pessoa_2{
		Nome:  "Thiago",
		Idade: 44,
	}

	pessoa.setNome("Beatriz")
	fmt.Println(pessoa.Nome)

}
