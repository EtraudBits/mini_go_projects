package estoque

import (
	"crypto/sha256" // pacote para gerar hash SHA256
	"encoding/hex"  // pacote para codificar em hexadecimal
	"errors"        // pacote para manipulação de erros
	"fmt"           // pacote para formatação de strings
)

func NovoProduto(nome string, quantidade int) Produto {
	return Produto {
		ID: gerarID(nome),
		Nome: nome,
		Quantidade: quantidade,
	}
}

func gerarID(nome string) string {
	// Gera um hash SHA256 do nome do produto
	hash := sha256.Sum256([]byte(nome))
	// Retorna os primeiros 16 caracteres do hash em hexadecimal
	// Mesmo nome sempre gerará o mesmo ID
	return hex.EncodeToString(hash[:])[:16]
}

// Erro para indicar que o estoque é insuficiente
var ErrEstoqueInsuficiente = errors.New("estoque insuficiente")

// Erro para indicar que o valor fornecido é inválido
var ErrValorInvalido = errors.New("valor inválido")

type Produto struct {

	ID string
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