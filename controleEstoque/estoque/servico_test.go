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

// Atualizar implementa o método da interface RepositorioEstoque do arquivo interface.go
func (m *mockRepositorioEstoque) Atualizar(produto Produto) error {
	for i := range m.produtos {
		if m.produtos[i].ID == produto.ID {
			m.produtos[i] = produto
			return nil
		}
	}
	return nil // para testes simples, retorna nil se não encontrar
}

func TestCadastrarProduto( t *testing.T) { // t *testing.T vem do pacote padrão "testing"

	// cria o mock do repositório para os testes
	mockRepo := &mockRepositorioEstoque{}

	// NovoServicoEstoque vem do arquivo servico.go
	servico := NovoServicoEstoque(mockRepo)

	// NovoProduto vem do arquivo produto.go
	produto := NovoProduto("viga", 12)

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
	mockRepo.Adicionar(NovoProduto("tijolo", 50)) // NovoProduto vem do arquivo produto.go
	mockRepo.Adicionar(NovoProduto("cimento", 30)) // NovoProduto vem do arquivo produto.go

	// ListarEstoque vem do arquivo servico.go
	produtos := servico.ListarEstoque()

	// verifica se a lista retornada está correta
	if len(produtos) != 2 { // esperamos 2 produtos
		t.Errorf("Esperava 2 produtos na lista, mas encontrei %d", len(produtos),) // t.Errorf vem do pacote padrão "testing"
	}
}

func TestVenderProdutoComEstoqueInsuficiente(t *testing.T) {
	mockRepo := &mockRepositorioEstoque{}
	servico := NovoServicoEstoque(mockRepo)

	mockRepo.Adicionar(NovoProduto("areia", 5))

	err := servico.VenderProduto("areia", 10) // VenderProduto vem do arquivo servico.go

	if err != nil {
		t.Errorf("Esperava erro de estoque insuficiente, mas não recebi erro")
	}
}


