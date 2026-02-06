package estoque

import (
	"errors"
)

// ServicoEstoque chama a interface RepositorioEstoque para gerenciar produtos no estoque
type ServicoEstoque struct {
	repositorio RepositorioEstoque // campo que armazena o repositório de estoque que esta implementa a interface RepositorioEstoque
}


// NovoServicoEstoque cria um novo serviço de estoque com o repositório fornecido
func NovoServicoEstoque(repo RepositorioEstoque) *ServicoEstoque {
	return &ServicoEstoque { // retorna um ponteiro para ServicoEstoque
		repositorio: repo, // repo significa o repositório passado como argumento que é atribuído ao campo repositorio
	}
}

// CadastrarProduto adiciona um novo produto ao estoque usando o repositório substituindo o método Adicionar da interface
func (s *ServicoEstoque) CadastrarProduto(produto Produto) {
	s.repositorio.Adicionar(produto) // chama o método Adicionar do repositório para adicionar o produto que esta no arquivo main.go
}

// ListarEstoque retorna a lista de produtos no estoque usando o repositório substituindo o método Listar da interface
func (s *ServicoEstoque) ListarEstoque() []Produto { 
	return s.repositorio.Listar() // chama o método Listar do repositório para listar os produtos que estão no estoque que estar no aruivo main.go
}

// VenderProduto diminui a quantidade de um produto no estoque
func (s *ServicoEstoque) VenderProduto(id string, quantidade int) error {
	produtos := s.repositorio.Listar() // obtém a lista de produtos do repositório

	for i := range produtos {  // faz um loop pelos os produtos
		if produtos[i].ID == id { // se o nome do produto for igual ao nome passado como argumento
			return produtos[i].DiminuirQuantidade(quantidade) // retorna o resultado da chamada do método DiminuirQuantidade do produto, onde podutos[i] significa o produto encontrado no estoque .DiminuirQuantidade(quantidade) vem do arquivo produto.go
		}
	}

	return errors.New("produto não encontrado") // retorna um erro se o produto não for encontrado
}