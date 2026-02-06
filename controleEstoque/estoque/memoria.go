package estoque

import "errors"

// RepositorioMemoria implementa o RepositorioEstoque armazenando produtos em memória
type RepositorioMemoria struct {
	produtos []Produto
}

// NovoRepositorioMemoria cria um repositório de estoque em memória
func NovoRepositorioMemoria() *RepositorioMemoria {
	return &RepositorioMemoria {
		produtos: []Produto{},
	}
}

// Adicionar insere um produto no estoque em memória.
func (e *RepositorioMemoria) Adicionar(produto Produto) {
	e.produtos = append(e.produtos, produto)	
}

// Atualizar modifica um produto existente no estoque em memória.
func (r *RepositorioMemoria) Atualizar(produto Produto) error { // percorre a lista de produtos para encontrar o produto com o ID correspondente
	for i:= range r.produtos { // itera sobre os produtos do repositório
		if r.produtos[i].ID == produto.ID { // se encontrar o produto com o ID correspondente
			// atualiza o produto no repositório
			r.produtos[i] = produto
			return nil // retorna nil se a atualização for bem-sucedida
		}
	}
	return errors.New("produto não encontrado") // retorna erro se o produto não for encontrado
}
// Listar devolve todos os produtos armazenados no estoque em memória.
func (e *RepositorioMemoria) Listar() []Produto {
	return e.produtos
}


