//Ponteiro

package main

import (
	"fmt"
)

func prox() {
	nome := "Wesley"
	// Cria váriavel que recebe um ponteiro
	var nome2 *string
	// & traz a posição da várivel na memória
	nome2 = &nome

	// Por nome2 está referenciado na posição 2 o valor de nome altera para "FullCycle"
	*nome2 = "FullCycle"

	fmt.Println(*nome2)
	// -> FullCycle
}

func pointers() {
	nome := "Wesley"
	var nome2 *string
	nome2 = &nome

	fmt.Println(*nome2)

	prox()
}
