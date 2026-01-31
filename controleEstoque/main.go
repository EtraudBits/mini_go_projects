package main

import "fmt"

// Definição da estrutura do produto
type Produto struct {

	Nome string
	Quantidade int

}
// Função para cadastrar um novo produto
func cadastrarProduto(nome string, quantidade int) Produto {
	return Produto{
		Nome: nome,
		Quantidade: quantidade, 
	}
}
// Função para adicionar um produto ao estoque
func adicionarEstoque(estoque *[]Produto, produto Produto) {
	*estoque  = append(*estoque, produto)
	
}
// Função para listar todos os produtos no estoque	
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

		adicionarEstoque(&estoque, viga)
		adicionarEstoque(&estoque, coluna)
		adicionarEstoque(&estoque, estacaTipoMourao)
		adicionarEstoque(&estoque, estacaCurvada)
		listarEstoque(estoque)

}

