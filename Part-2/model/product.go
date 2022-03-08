package model

import uuid "github.com/satori/go.uuid"

/*
O GO permite que geramos tag para referência valores caso esteja em JSON
*/
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Product []Product
}

// p recebe o array de product já o Add recebe os dados do product
func (p *Products) Add(product Product) {
	p.Product = append(p.Product, product)

}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}
