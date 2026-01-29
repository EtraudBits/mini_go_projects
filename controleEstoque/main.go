package main

import "fmt"

type Produto struct {

	Nome string
	Quantidade int

}

func cadastrarProduto(nome string, quantidade int) Produto {
	return Produto{
		Nome: nome,
		Quantidade: quantidade, 
	}
}

func adicionarEstoque(estoque []Produto, produto Produto) []Produto {
	estoque  = append(estoque, produto)
	return estoque
}

func listarEstoque(estoque []Produto) {
	for _, produto := range estoque {
		fmt.Printf("Produto: %s | Quantidade: %d\n", produto.Nome, produto.Quantidade)
	}
}

func main() {

	estoque := []Produto{}

		viga := cadastrarProduto("viga", 11)
		coluna := cadastrarProduto("coluna", 10)
		estacaTipoMourao := cadastrarProduto("estaca tipo mourao", 100)
		estacaCurvada := cadastrarProduto("estaca curvada", 15)

		estoque = adicionarEstoque(estoque, viga)
		estoque = adicionarEstoque(estoque, coluna)
		estoque = adicionarEstoque(estoque, estacaTipoMourao)
		estoque = adicionarEstoque(estoque, estacaCurvada)
		listarEstoque(estoque)

}

