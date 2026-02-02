package estoque

// EstoqueMemoria implementa o RepositorioEstoque armazenando produtos em memória
type EstoqueMemoria struct {
	produtos []Produto
}

// NovoEstoqueMemoria cria um repositório de estoque em memória
func NovoEstoqueMemoria() *EstoqueMemoria {
	return &EstoqueMemoria {
		produtos: []Produto{},
	}
}

// Adicionar insere um produto no estoque em memória.
func (e *EstoqueMemoria) Adicionar(produto Produto) {
	e.produtos = append(e.produtos, produto)	
}

// Lista devolve todos os produtos armazenados no estoque em memória.
func (e *EstoqueMemoria) Listar() []Produto {
	return e.produtos
}

