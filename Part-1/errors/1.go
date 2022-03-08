package main

import (
	"errors"
	"fmt"
)

type PessoaInFunction struct {
	Nome  string
	Idade int
}

func andou(pessoa PessoaInFunction) (string, error) {
	if pessoa.Nome != "Wesley" {
		return "", errors.New("Nome tem que ser Wesley")
	}

	fmt.Println(pessoa.Nome, "andou")

	return pessoa.Nome + "andou", nil

}

// Forma de tratart error.
func functions() {
	pessoa := PessoaInFunction{
		Nome:  "Thiago",
		Idade: 44,
	}
	// Forma de tratar error em Go
	res, err := andou(pessoa)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)

}
