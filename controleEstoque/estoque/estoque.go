package estoque

// Estrutura para representar o estoque
type Estoque struct {
	Produtos []Produto // Slice para armazenar os produtos no estoque
}

// Função para criar um novo estoque
func NovoEstoque() *Estoque { // Retorna um ponteiro para um novo estoque
	return &Estoque{ // Retorna o endereço de um novo Estoque
		Produtos: []Produto{}, // Inicializa o slice de produtos vazio
	}
}

// Método para adicionar um produto ao estoque
func (e *Estoque) Adicionar(produto Produto) {
	e.Produtos = append(e.Produtos, produto) // Adiciona o produto ao slice de produtos
}

// Método para listar os produtos no estoque
func (e Estoque) Listar() { 
	for _, produto := range e.Produtos { // Itera sobre os produtos no estoque
		produto.Exibir() // Chama o método Exibir para cada produto
	}
}