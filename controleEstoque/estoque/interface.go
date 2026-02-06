package estoque

//RepositorioEstoque define o contrato de armazenamento de produtos no estoque
// Qualquer estrutura que implemente esse método pode ser usada como repositório de estoque
type RepositorioEstoque interface { 
	Adicionar(produto Produto) // se conter esse método, pode ser usado como repositório
	Atualizar(produto Produto) error // se conter esse método, pode ser usado como repositório
	Listar() []Produto // se conter esse método, pode ser usado como repositório
}