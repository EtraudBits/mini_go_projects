package main

import "fmt"

// Definição da estrutura do produto
type Produto struct {

	Nome string
	Quantidade int

}

// Métodos para aumentar e diminuir a quantidade do produto
func (p *Produto) AumentarQuantidade(valor int) { 
	p.Quantidade += valor
}

// Método para diminuir a quantidade do produto
func (p *Produto) DiminiurQuantidade(valor int) {
	if p.Quantidade >= valor { // Verifica se há quantidade suficiente
		p.Quantidade -= valor // Diminui a quantidade
	}
}

func (p *Produto) Exibir() {
	fmt.Printf("Produto: %s | Quantidade: %d\n", p.Nome, p.Quantidade)
}

// Função para cadastrar um novo produto
func cadastrarProduto(nome string, quantidade int) Produto {
	return Produto{
		Nome: nome,
		Quantidade: quantidade, 
	}
}
// Função para adicionar um produto ao estoque
// Usamos um ponteiro para modificar o slice original
func adicionarEstoque(estoque *[]Produto, produto Produto) { // recebe o endereço do slice ([]Produto)
	*estoque  = append(*estoque, produto) // altere o valor dentro desse endereço 
	
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

		//
		viga.AumentarQuantidade(6) // Aumenta a quantidade da viga em 6
		coluna.DiminiurQuantidade(2) // Diminui a quantidade da coluna em 2

		// Adicionando produtos ao estoque chamando a função e passando o endereço do slice
		adicionarEstoque(&estoque, viga) // &estoque passa o endereço do slice
		adicionarEstoque(&estoque, coluna)
		adicionarEstoque(&estoque, estacaTipoMourao)
		adicionarEstoque(&estoque, estacaCurvada)

	// Listando os produtos no estoque
		listarEstoque(estoque)

}

