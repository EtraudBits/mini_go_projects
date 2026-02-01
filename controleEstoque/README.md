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

- Estruturas de dados (structs)
- Métodos e receivers
- Ponteiros e gerenciamento de memória
- Organização de código em pacotes
- Slices e manipulação de coleções
- Modularização e arquitetura de software

---

## 📂 Estrutura do Projeto

```
controleEstoque/
├── go.mod                 # Gerenciamento de módulo
├── main.go               # Ponto de entrada da aplicação
├── estoque/              # Pacote de lógica de negócio
│   ├── produto.go        # Estrutura e métodos de Produto
│   └── estoque.go        # Estrutura e métodos de Estoque
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
Produto: viga | Quantidade: 17
Produto: coluna | Quantidade: 8
Produto: estaca tipo mourao | Quantidade: 100
Produto: estaca curvada | Quantidade: 15
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

### **Métodos com Receivers**

```go
func (p *Produto) AumentarQuantidade(valor int) {
    p.Quantidade += valor
}
```

### **Ponteiros**

- Uso de ponteiros (`*Estoque`) para modificar o estado original
- Factory functions retornando ponteiros para novas instâncias

### **Pacotes**

- Organização modular do código
- Exportação de tipos e funções (primeira letra maiúscula)
- Encapsulamento de lógica de negócio

---

## 📚 Aprendizados e Notas

Este projeto serve como documentação viva do processo de aprendizado em Go. Cada commit representa um passo na jornada de compreensão da linguagem, desde conceitos básicos até padrões mais avançados de desenvolvimento.

---

## 📄 Licença

Este é um projeto educacional de código aberto para fins de aprendizado.

---

**Última atualização:** Janeiro 2026  
**Versão atual:** 2.0 - Refatoração e Organização em Pacotes
