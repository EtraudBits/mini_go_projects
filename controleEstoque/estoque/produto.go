package estoque

import (
	"fmt"
)

type Produto struct {

	Nome string
	Quantidade int

}

// Métodos para aumentar e diminuir a quantidade do produto
func (p *Produto) AumentarQuantidade(valor int) { 
	p.Quantidade += valor
}

// Método para diminuir a quantidade do produto
func (p *Produto) DiminiurQuantidade(valor int) {
	if p.Quantidade >= valor { // Verifica se há quantidade suficiente
		p.Quantidade -= valor // Diminui a quantidade
	}
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