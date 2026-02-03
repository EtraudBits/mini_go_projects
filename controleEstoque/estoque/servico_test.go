package estoque

import "testing" // pacote padrão do Go para testes

// mockRepositorioEstoque é uma implementação falsa do RepositorioEstoque para testes
// serve para apenas testar a lógica do serviço de estoque sem depender de um banco de dados real
// implementa a interface RepositorioEstoque do arquivo interface.go
type mockRepositorioEstoque struct {
	produtos []Produto // usa o tipo Produto do arquivo produto.go
}

// Adicionar implementa o método da interface RepositorioEstoque do arquivo interface.go
func (m *mockRepositorioEstoque) Adicionar(produto Produto) { // usa o tipo Produto do arquivo produto.go
	m.produtos = append(m.produtos, produto)
}

// Listar implementa o método da interface RepositorioEstoque do arquivo interface.go
func (m *mockRepositorioEstoque) Listar() []Produto { // usa o tipo Produto do arquivo produto.go
	return m.produtos
}

func TestCadastrarProduto( t *testing.T) { // t *testing.T vem do pacote padrão "testing"

	// cria o mock do repositório para os testes
	mockRepo := &mockRepositorioEstoque{}

	// NovoServicoEstoque vem do arquivo servico.go
	servico := NovoServicoEstoque(mockRepo)

	// Produto vem do arquivo produto.go
	produto := Produto {
		Nome: "viga",
		Quantidade: 12,
	}

	// CadastrarProduto vem do arquivo servico.go
	servico.CadastrarProduto(produto)

	// t.Errorf vem do pacote padrão "testing"
	if len(mockRepo.produtos) != 1 {
		t.Errorf("Esperava 1 produto no repositório, mas encontrei %d", len(mockRepo.produtos),)
	}

}

func TestListarEstoque(t *testing.T) { // t *testing.T vem do pacote padrão "testing"

	// cria o mock e o serviço para os testes
	mockRepo := &mockRepositorioEstoque{}
	servico := NovoServicoEstoque(mockRepo) // NovoServicoEstoque vem do arquivo servico.go

	// adiciona alguns produtos ao mock diretamente usando o método Adicionar (linha 12 deste arquivo)
	mockRepo.Adicionar(Produto{Nome: "tijolo", Quantidade: 50}) // Produto vem do arquivo produto.go
	mockRepo.Adicionar(Produto{Nome: "cimento", Quantidade: 30}) // Produto vem do arquivo produto.go

	// ListarEstoque vem do arquivo servico.go
	produtos := servico.ListarEstoque()

	// verifica se a lista retornada está correta
	if len(produtos) != 2 { // esperamos 2 produtos
		t.Errorf("Esperava 2 produtos na lista, mas encontrei %d", len(produtos),) // t.Errorf vem do pacote padrão "testing"
	}
}
