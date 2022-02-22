package main

import (
	"errors"
	"fmt"
)

type Pessoa_2 struct {
	Nome  string
	Idade int
}

func (p Pessoa_2) andou() (string, error) {
	if p.Nome != "Wesley" {
		return "", errors.New("Nome tem que ser Wesley")
	}

	fmt.Println(p.Nome, "andou")

	return p.Nome + "andou", nil

}

func object() {
	pessoa := Pessoa_2{
		Nome:  "Thiago",
		Idade: 44,
	}
	// Posso tratar como se fosse um método da POO
	res, err := pessoa.andou()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)

}
