package estoque

import (
	"errors"
	"fmt"
)

// Erro para indicar que o estoque é insuficiente
var ErrEstoqueInsuficiente = errors.New("estoque insuficiente")

// Erro para indicar que o valor fornecido é inválido
var ErrValorInvalido = errors.New("valor inválido")

type Produto struct {

	Nome string
	Quantidade int

}

// Métodos para aumentar e diminuir a quantidade do produto
func (p *Produto) AumentarQuantidade(valor int) { 
	p.Quantidade += valor
}

// Método para diminuir a quantidade do produto
func (p *Produto) DiminuirQuantidade(valor int) error {
	if valor <= 0 { // valida se o valor é positivo
		return ErrValorInvalido // retorna erro se o valor for inválido
	}
	// verifica se há estoque suficiente (regra de negócio)
	if p.Quantidade < valor {
		return ErrEstoqueInsuficiente // retorna erro se o estoque for insuficiente
	}

	// serve para diminuir a quantidade do produto - altera segura do estado do objeto
	p.Quantidade -= valor // diminui a quantidade do produto
	return nil // se não houver erro, retorna nil
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