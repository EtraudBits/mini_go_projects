package main

import "fmt"

type Produto struct {

	Nome string
	Quantidade int

}

func main() {

	estoque := []Produto{}

		estoque = append(estoque, Produto{
			Nome: "Cobogo Arabe",
			Quantidade: 10,
		})
		estoque = append(estoque, Produto{
			Nome: "Cobogo Taco Chines",
			Quantidade: 2,
		})
		estoque = append(estoque, Produto{
			Nome: "Cobogo Flor",
			Quantidade: 5,
		})
		estoque = append(estoque, Produto{
			Nome: "Cobogo 4 Pontas",
			Quantidade: 7,
		})

		for i, produto := range estoque {
			fmt.Printf ("%d - Produto: %s | Quantidade: %d\n", i, produto.Nome, produto.Quantidade)
		}
	}

