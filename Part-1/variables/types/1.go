package main

import "fmt"

type nome []string // Posso utilizar para tipar as variáveis

func index() {
	var nome nome // define o tipo da variável nome

	nome = append(nome, "Teste 1")
	nome = append(nome, "Teste 2")

	for _, v := range nome {
		fmt.Println(v)
	}

}
