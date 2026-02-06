package main

import (
	"controleEstoque/estoque"
)

// Função principal do programa
func main() {

	repo := estoque.NovoRepositorioArquivo("estoque.json") // cria um novo repositório de estoque em arquivo chamando a função NovoRepositorioArquivo do pacote arquivo.go, passando o nome do arquivo onde os dados serão armazenados
	
	servico := estoque.NovoServicoEstoque(repo) // cria um novo serviço de estoque passando o repositório como parâmetro

	// Criando alguns produtos usando a função NovoProduto
	viga := estoque.NovoProduto("viga", 11) // cria um novo produto chamando a função NovoProduto
	coluna := estoque.NovoProduto("coluna", 10)
	estacaTipoMourao := estoque.NovoProduto("estaca tipo mourao", 100)
	estacaCurvada := estoque.NovoProduto("estaca curvada", 15)
	cobogoFlor := estoque.NovoProduto("cobogo flor", 55)

	//
	viga.AumentarQuantidade(6) // Aumenta a quantidade da viga em 6 (de 11 para 17)
	coluna.DiminuirQuantidade(2) // Diminui a quantidade da coluna em 2 (de 10 para 8)

		// Adicionando os produtos ao estoque
		servico.CadastrarProduto(viga) // Chama o método CadastrarProduto do serviço para adicionar o produto viga
		servico.CadastrarProduto(coluna)
		servico.CadastrarProduto(estacaTipoMourao)
		servico.CadastrarProduto(estacaCurvada)
		servico.CadastrarProduto(cobogoFlor)


	// Listando os produtos no estoque
	produtos := servico.ListarEstoque() // Chama o método ListarProdutos do serviço para obter a lista de produtos

	// Imprimindo os produtos no estoque
	for _, produto := range produtos {
		println("Produto:", produto.Nome, "Quantidade:", produto.Quantidade)
	}

}

