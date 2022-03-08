package main

import (
	uuid "github.com/satori/go.uuid"
	"github.com/thiagorcode/studies_golang/Part-2/model"
)

func main() {
	produto1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}
	produto2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)

}
