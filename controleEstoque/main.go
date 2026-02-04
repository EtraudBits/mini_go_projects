package main

import (
	"controleEstoque/estoque"
)

// Função principal do programa
func main() {

	repo := estoque.NovoEstoqueMemoria() // cria um novo repositório de estoque em memória chamando a função NovoEstoqueMemoria do pacote memoria.go

	servico := estoque.NovoServicoEstoque(repo) // cria um novo serviço de estoque chamando a função NovoServicoEstoque do pacote servico.go e passando o repositório criado acima
		
	
		viga := estoque.Produto { // cria um novo produto no pacote estoque no tipo Produto
			Nome: "viga",
			Quantidade: 11,
		}
		coluna := estoque.Produto{
			Nome: "coluna",
			Quantidade: 10,
		}
		estacaTipoMourao := estoque.Produto{
			Nome: "estaca tipo mourao",
			Quantidade: 100,
		}
		estacaCurvada := estoque.Produto{
			Nome: "estaca curvada",
			Quantidade: 15,
		}

		//
		viga.AumentarQuantidade(6) // Aumenta a quantidade da viga em 6
		coluna.DiminuirQuantidade(2) // Diminui a quantidade da coluna em 2

		// Adicionando os produtos ao estoque
		servico.CadastrarProduto(viga) // Chama o método CadastrarProduto do serviço para adicionar o produto viga
		servico.CadastrarProduto(coluna)
		servico.CadastrarProduto(estacaTipoMourao)
		servico.CadastrarProduto(estacaCurvada)


	// Listando os produtos no estoque
	produtos := servico.ListarEstoque() // Chama o método ListarProdutos do serviço para obter a lista de produtos

	// Imprimindo os produtos no estoque
	for _, produto := range produtos {
		println("Produto:", produto.Nome, "Quantidade:", produto.Quantidade)
	}

}

