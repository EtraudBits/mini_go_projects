package estoque

import (
	"errors"
	"sync"
)

// RepositorioMemoria implementa o RepositorioEstoque armazenando produtos em memória
type RepositorioMemoria struct {
	produtos []Produto
	mu sync.Mutex // mutex para garantir acesso seguro ao repositório em caso de concorrência
}

// NovoRepositorioMemoria cria um repositório de estoque em memória
func NovoRepositorioMemoria() *RepositorioMemoria {
	return &RepositorioMemoria {
		produtos: []Produto{},
	}
}

// Adicionar insere um produto no estoque em memória.
func (r *RepositorioMemoria) Adicionar(produto Produto) {
	r.mu.Lock() // bloqueia o mutex para garantir que apenas uma goroutine possa acessar o repositório ao mesmo tempo
	defer r.mu.Unlock() // desbloqueia o mutex após a função ser executada, garantindo que outros goroutines possam acessar o repositório
	r.produtos = append(r.produtos, produto) // adiciona o produto à lista de produtos do repositório
}

// Atualizar modifica um produto existente no estoque em memória.
func (r *RepositorioMemoria) Atualizar(produto Produto) error { // percorre a lista de produtos para encontrar o produto com o ID correspondente
	r.mu.Lock() // bloqueia o mutex para garantir que apenas uma goroutine possa acessar o repositório ao mesmo tempo
	defer r.mu.Unlock() // desbloqueia o mutex após a função ser executada, garantindo que outros goroutines possam acessar o repositório

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
func (r *RepositorioMemoria) Listar() []Produto {
	r.mu.Lock() // bloqueia o mutex para garantir que apenas uma goroutine possa acessar o repositório ao mesmo tempo
	defer r.mu.Unlock() // desbloqueia o mutex após a função ser executada, garantindo que outros goroutines possam acessar o repositório

	copia := make([]Produto, len(r.produtos)) // cria uma cópia da lista de produtos para evitar que o chamador modifique diretamente a lista interna do repositório
	copy(copia, r.produtos) // copia os produtos para a nova lista

	return copia // retorna a cópia da lista de produtos
}


