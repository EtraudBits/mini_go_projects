package main

import (
	"controleEstoque/estoque"
)


// Função principal do programa
func main() {

	estoqueLoja := estoque.NovoEstoque() // Cria um novo estoque usando a função do pacote estoque

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
		coluna.DiminiurQuantidade(2) // Diminui a quantidade da coluna em 2

		// Adicionando os produtos ao estoque
		estoqueLoja.Adicionar(viga)
		estoqueLoja.Adicionar(coluna)
		estoqueLoja.Adicionar(estacaTipoMourao)
		estoqueLoja.Adicionar(estacaCurvada)

	// Listando os produtos no estoque
	estoqueLoja.Listar() // Chama o método Listar para exibir os produtos no estoque

}

