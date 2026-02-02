# 📦 Sistema de Controle de Estoque

> Mini-projeto de aprendizado em **Golang** - Do básico ao avançado

**Autor:** Duarte Rodrigo Santos de Oliveira  
**LinkedIn:** [www.linkedin.com/in/duarte-backend-golang](https://www.linkedin.com/in/duarte-backend-golang)

---

## 📋 Sobre o Projeto

Este é um projeto educacional desenvolvido para aprender e praticar conceitos fundamentais da linguagem Go (Golang). O projeto está sendo construído de forma incremental, começando com conceitos básicos e evoluindo gradualmente para um sistema mais completo e robusto.

**Status:** 🚧 Em desenvolvimento ativo - Este README será atualizado a cada nova funcionalidade implementada.

---

## 🎯 Objetivos de Aprendizado

- ✅ Estruturas de dados (structs)
- ✅ Métodos e receivers
- ✅ Ponteiros e gerenciamento de memória
- ✅ Organização de código em pacotes
- ✅ Slices e manipulação de coleções
- ✅ Modularização e arquitetura de software
- ✅ **Interfaces e contratos**
- ✅ **Padrões de projeto (Repository, Service Layer)**
- ✅ **Dependency Injection**
- ✅ **Princípios SOLID**

---

## 📂 Estrutura do Projeto

```
controleEstoque/
├── go.mod                 # Gerenciamento de módulo
├── main.go               # Ponto de entrada da aplicação
├── estoque/              # Pacote de lógica de negócio
│   ├── produto.go        # Estrutura e métodos de Produto
│   ├── interface.go      # Interface RepositorioEstoque (contrato)
│   ├── memoria.go        # Implementação em memória do repositório
│   └── servico.go        # Camada de serviço (lógica de negócio)
└── README.md            # Este arquivo
```

---

## 🚀 Evolução do Projeto

### **Versão 1.0 - Fundamentos Básicos**

- ✅ Criação da estrutura `Produto` com campos Nome e Quantidade
- ✅ Implementação de métodos básicos:
  - `AumentarQuantidade()` - Adiciona unidades ao produto
  - `DiminuirQuantidade()` - Remove unidades do produto (com validação)
  - `Exibir()` - Exibe informações do produto
- ✅ Função `cadastrarProduto()` para criar novos produtos
- ✅ Sistema básico de estoque usando slices
- ✅ Funções para adicionar e listar produtos

### **Versão 2.0 - Refatoração e Organização**

- ✅ Reorganização do código em pacotes separados
- ✅ Criação do pacote `estoque` para modularização
- ✅ Separação de responsabilidades:
  - `produto.go` - Lógica relacionada a produtos
  - `estoque.go` - Lógica de gerenciamento do estoque
- ✅ Implementação da estrutura `Estoque` com métodos:
  - `NovoEstoque()` - Factory function para criar estoque
  - `Adicionar()` - Adiciona produtos ao estoque
  - `Listar()` - Lista todos os produtos
- ✅ Uso adequado de ponteiros para modificação de estado
- ✅ Adição de `go.mod` para gerenciamento de dependências

### **Versão 3.0 - Arquitetura em Camadas com Interfaces**

- ✅ Implementação de **Interfaces** (`RepositorioEstoque`):
  - Define contratos para operações de estoque
  - Permite múltiplas implementações do repositório
  - Facilita testes e manutenção
- ✅ **Padrão Repository** com `EstoqueMemoria`:
  - Implementação concreta da interface
  - Armazenamento em memória
  - Preparado para futuras implementações (banco de dados, arquivo, etc.)
- ✅ **Camada de Serviço** (`ServicoEstoque`):
  - Separa lógica de negócio da camada de dados
  - Usa a interface `RepositorioEstoque` (inversão de dependência)
  - Métodos `CadastrarProduto()` e `ListarProdutos()`
- ✅ **Refatoração completa da arquitetura**:
  - Remoção de código redundante (`estoque.go`)
  - Aplicação de princípios SOLID
  - Código mais testável e manutenível

---

## 💻 Como Executar

### Pré-requisitos

- Go 1.22.2 ou superior instalado

### Executando o projeto

```bash
# Navegue até o diretório do projeto
cd controleEstoque

# Execute o programa
go run main.go
```

### Exemplo de Saída

```
Produto: viga Quantidade: 17
Produto: coluna Quantidade: 8
Produto: estaca tipo mourao Quantidade: 100
Produto: estaca curvada Quantidade: 15
```

---

## 📝 Conceitos Aplicados

### **Structs**

```go
type Produto struct {
    Nome       string
    Quantidade int
}
```

### **Interfaces**

```go
type RepositorioEstoque interface {
    Adicionar(produto Produto)
    Listar() []Produto
}
```

**Benefícios das Interfaces:**

- Define contratos entre componentes
- Permite trocar implementações sem alterar código cliente
- Facilita testes com mocks/stubs
- Reduz acoplamento entre camadas

### **Métodos com Receivers**

```go
func (p *Produto) AumentarQuantidade(valor int) {
    p.Quantidade += valor
}

func (e *EstoqueMemoria) Adicionar(produto Produto) {
    e.produtos = append(e.produtos, produto)
}
```

### **Ponteiros**

- Uso de ponteiros (`*Estoque`, `*ServicoEstoque`) para modificar o estado original
- Factory functions retornando ponteiros para novas instâncias
- Receivers com ponteiros para métodos que modificam estado

### **Pacotes**

- Organização modular do código
- Exportação de tipos e funções (primeira letra maiúscula)
- Encapsulamento de lógica de negócio

### **Padrões de Arquitetura**

**Repository Pattern:**

- Abstrai a camada de persistência
- Implementações específicas (`EstoqueMemoria`)
- Facilita adição de novos meios de armazenamento

**Service Layer:**

- Centraliza lógica de negócio
- Usa interfaces para desacoplar da implementação
- Facilita testes e manutenção

**Dependency Injection:**

- Serviço recebe repositório via construtor
- Inversão de dependência (depende de interface, não de implementação)
- Mais flexível e testável

---

## 📚 Aprendizados e Notas

Este projeto serve como documentação viva do processo de aprendizado em Go. Cada commit representa um passo na jornada de compreensão da linguagem, desde conceitos básicos até padrões mais avançados de desenvolvimento.

**Principais Lições da Versão 3.0:**

- **Interfaces são contratos**: Definem o que precisa ser feito, não como fazer
- **Qualquer tipo que implemente os métodos da interface automaticamente a satisfaz** (não precisa declarar explicitamente)
- **Interfaces facilitam testes**: Permite criar mocks sem alterar código de produção
- **Repository Pattern desacopla persistência**: Trocar de memória para banco de dados não afeta o resto do código
- **Service Layer centraliza regras de negócio**: Mantém a lógica separada da camada de dados
- **Dependency Injection através de construtores**: Aumenta flexibilidade e testabilidade
- **Refatoração é importante**: Remover código redundante mantém o projeto limpo e manutenível

---

## 📄 Licença

Este é um projeto educacional de código aberto para fins de aprendizado.

---

**Última atualização:** Fevereiro 2026  
**Versão atual:** 3.0 - Arquitetura em Camadas com Interfaces
