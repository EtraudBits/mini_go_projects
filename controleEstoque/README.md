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
- ✅ **Testes Unitários**
- ✅ **Mocking e Test Doubles**
- ✅ **Tratamento de Erros**
- ✅ **Validação de Regras de Negócio**
- ✅ **Persistência em JSON**
- ✅ **Geração Automática de IDs**
- ✅ **Padrão Factory**

---

## 📂 Estrutura do Projeto

```
controleEstoque/
├── go.mod                 # Gerenciamento de módulo
├── main.go               # Ponto de entrada da aplicação
├── estoque/              # Pacote de lógica de negócio
│   ├── produto.go        # Estrutura e métodos de Produto + geração de ID
│   ├── interface.go      # Interface RepositorioEstoque (contrato)
│   ├── memoria.go        # Implementação em memória do repositório
│   ├── arquivo.go        # Implementação com persistência em JSON
│   ├── servico.go        # Camada de serviço (lógica de negócio)
│   └── servico_test.go   # Testes unitários do serviço
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
  - Métodos `CadastrarProduto()` e `ListarEstoque()`
- ✅ **Refatoração completa da arquitetura**:
  - Remoção de código redundante (`estoque.go`)
  - Aplicação de princípios SOLID
  - Código mais testável e manutenível

### **Versão 4.0 - Testes Unitários e Qualidade de Código**

- ✅ Implementação de **Testes Unitários** (`servico_test.go`):
  - Testes para `CadastrarProduto()`
  - Testes para `ListarEstoque()`
  - Uso do pacote `testing` do Go
- ✅ **Mock Objects** para testes isolados:
  - Criação de `mockRepositorioEstoque`
  - Implementação dos métodos da interface para testes
  - Testes sem dependência de implementações reais
- ✅ **Documentação detalhada do código de teste**:
  - Comentários explicando origem de cada função/tipo
  - Referências aos arquivos fonte
  - Facilita compreensão do fluxo de testes
- ✅ **Validação de comportamento**:
  - Verificação de quantidade de produtos cadastrados
  - Validação de listagem de estoque
  - Uso de `t.Errorf()` para mensagens de erro descritivas

### **Versão 5.0 - Tratamento de Erros e Validação de Negócio**

- ✅ **Sistema de Erros Customizados**:
  - `ErrEstoqueInsuficiente` - Erro quando há tentativa de venda sem estoque
  - `ErrValorInvalido` - Erro para valores inválidos (negativos ou zero)
  - Uso de `errors.New()` para criação de erros semânticos
- ✅ **Refatoração do método `DiminuirQuantidade()`**:
  - Retorna `error` em vez de silenciosamente falhar
  - Valida se o valor é positivo antes de processar
  - Verifica disponibilidade de estoque antes da operação
  - Permite tratamento adequado de erros no código cliente
- ✅ **Implementação do método `VenderProduto()`**:
  - Busca produto por nome no repositório
  - Utiliza validação do método `DiminuirQuantidade()`
  - Retorna erro específico se produto não for encontrado
  - Propaga erros de estoque insuficiente adequadamente
- ✅ **Testes de Validação**:
  - `TestVenderProdutoComEstoqueInsuficiente()` verifica tratamento de erro
  - Garante que regras de negócio são respeitadas
  - Valida comportamento em cenários de erro
- ✅ **Correções e Melhorias**:
  - Corrigido typo: `DiminiurQuantidade` → `DiminuirQuantidade`
  - Código mais robusto e previsível
  - Melhor experiência para quem usa a API

### **Versão 6.0 - Persistência em Arquivo e Geração de IDs**

- ✅ **Implementação do `RepositorioArquivo`**:
  - Nova implementação da interface `RepositorioEstoque`
  - Persistência de dados em arquivo JSON
  - Métodos `Listar()`, `Adicionar()` e `Atualizar()`
  - Leitura e escrita automática no arquivo
  - Tratamento de arquivo inexistente (estoque vazio)
  - Uso de `json.Marshal` e `json.Unmarshal` para serialização
- ✅ **Sistema de Geração de IDs**:
  - Função `NovoProduto()` com padrão Factory
  - Geração automática de IDs únicos usando timestamp (nanossegundos)
  - Campo `ID` adicionado à estrutura `Produto`
  - IDs baseados em `time.Now().UnixNano()` garantem unicidade
- ✅ **Refatoração do Produto**:
  - Adicionado campo `ID string` à struct `Produto`
  - Função `gerarID()` privada para criação de identificadores
  - `NovoProduto(nome, quantidade)` substitui criação manual de produtos
- ✅ **Atualização do Fluxo Principal**:
  - `main.go` usa `NovoProduto()` para criação de produtos
  - Integração com `ServicoEstoque` usando `RepositorioArquivo`
  - Demonstração de persistência em arquivo JSON
- ✅ **Correções de Nomenclatura**:
  - `EstoqueMemoria` renomeado para `RepositorioMemoria`
  - `NovoEstoqueMemoria()` renomeado para `NovoRepositorioMemoria()`
  - Consistência de nomes entre repositórios (Memória e Arquivo)

### **Versão 6.1 - IDs Determinísticos e Prevenção de Duplicatas**

- ✅ **Refatoração da Geração de IDs**:
  - IDs agora baseados em hash SHA256 do nome do produto
  - Mesmo nome sempre gera o mesmo ID (determinístico)
  - Previne duplicatas: produtos com mesmo nome têm mesmo ID
  - Usa `crypto/sha256` e `encoding/hex` para hash confiável
  - Retorna 16 caracteres hexadecimais do hash
- ✅ **Vantagens do Sistema de ID Determinístico**:
  - Evita cadastro duplicado de produtos
  - ID permanece consistente entre execuções do programa
  - Permite busca tanto por nome quanto por ID
  - Facilita identificação e comparação de produtos
- ✅ **Atualização dos Testes**:
  - Testes refatorados para usar `NovoProduto()` consistentemente
  - Mock atualizado com método `Atualizar()` completo
  - Implementação total da interface `RepositorioEstoque` no mock
  - Garantia de que testes seguem as mesmas práticas do código de produção

### **Versão 7.0 - Concorrência no Repositório em Memória**

- ✅ **Proteção de Estado Compartilhado**:
  - Adicionado `sync.Mutex` ao `RepositorioMemoria`
  - `Lock()` e `Unlock()` aplicados em `Adicionar()`, `Atualizar()` e `Listar()`
  - Evita condições de corrida em acesso concorrente

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

### Executando os testes

```bash
# Execute todos os testes do pacote estoque
go test ./estoque

# Execute com saída detalhada (verbose)
go test -v ./estoque

# Execute um teste específico
go test -v ./estoque -run TestCadastrarProduto
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

### **Testes Unitários**

```go
func TestCadastrarProduto(t *testing.T) {
    mockRepo := &mockRepositorioEstoque{}
    servico := NovoServicoEstoque(mockRepo)

    produto := Produto{Nome: "viga", Quantidade: 12}
    servico.CadastrarProduto(produto)

    if len(mockRepo.produtos) != 1 {
        t.Errorf("Esperava 1 produto, mas encontrei %d", len(mockRepo.produtos))
    }
}
```

**Benefícios dos Testes:**

- Validam o comportamento esperado do código
- Detectam regressões e bugs rapidamente
- Servem como documentação viva do sistema
- Facilitam refatorações com segurança

**Mock Objects:**

- Simulam implementações reais para testes isolados
- Não dependem de banco de dados ou recursos externos
- Testam apenas a lógica do serviço
- Permitem controle total sobre o comportamento do repositório

---

## 📚 Aprendizados e Notas

Este projeto serve como documentação viva do processo de aprendizado em Go. Cada commit representa um passo na jornada de compreensão da linguagem, desde conceitos básicos até padrões mais avançados de desenvolvimento.

**Principais Lições da Versão 1.0:**

- **Structs são a base da programação em Go**: Permitem agrupar dados relacionados
- **Métodos com receivers conectam funções a tipos**: Sintaxe `func (p *Produto) Metodo()`
- **Ponteiros são essenciais para modificar estado**: `*Produto` permite alterar o produto original
- **Slices são arrays dinâmicos**: Crescem conforme necessário com `append()`
- **Validação de dados é importante**: Sempre verificar valores antes de operações (ex: não permitir quantidade negativa)
- **Funções auxiliares organizam o código**: `cadastrarProduto()` separa responsabilidades

**Principais Lições da Versão 2.0:**

- **Pacotes organizam código em módulos**: Separar lógica em arquivos diferentes facilita manutenção
- **Exportação com letra maiúscula**: `Produto` é exportado, `produto` seria privado ao pacote
- **Factory functions são o padrão Go**: `NovoEstoque()` substitui construtores de outras linguagens
- **Separação de responsabilidades melhora código**: `produto.go` cuida de produtos, `estoque.go` cuida do estoque
- **`go.mod` gerencia dependências**: Define o módulo e suas dependências externas
- **Importação de pacotes locais**: Usar o caminho do módulo (ex: `controleEstoque/estoque`)

**Principais Lições da Versão 3.0:**

- **Interfaces são contratos**: Definem o que precisa ser feito, não como fazer
- **Qualquer tipo que implemente os métodos da interface automaticamente a satisfaz** (não precisa declarar explicitamente)
- **Interfaces facilitam testes**: Permite criar mocks sem alterar código de produção
- **Repository Pattern desacopla persistência**: Trocar de memória para banco de dados não afeta o resto do código
- **Service Layer centraliza regras de negócio**: Mantém a lógica separada da camada de dados
- **Dependency Injection através de construtores**: Aumenta flexibilidade e testabilidade
- **Refatoração é importante**: Remover código redundante mantém o projeto limpo e manutenível

**Principais Lições da Versão 4.0:**

- **Testes unitários são essenciais**: Garantem que o código funciona conforme esperado
- **Mocks isolam testes**: Testam apenas a unidade de código desejada sem dependências externas
- **Pacote `testing` do Go é simples e poderoso**: Não requer frameworks externos para testes básicos
- **Interfaces facilitam mocking**: Criar um mock é apenas implementar os métodos da interface
- **Testes são documentação executável**: Mostram como o código deve ser usado e o comportamento esperado
- **Nomenclatura de testes**: Funções de teste devem começar com `Test` seguido do nome descritivo
- **`t.Errorf()` fornece feedback claro**: Mensagens descritivas ajudam a identificar falhas rapidamente
- **Cada teste deve ser independente**: Não deve depender de outros testes ou ordem de execução
- **Comentários nos testes auxiliam compreensão**: Especialmente útil para aprendizado, indicando origem de tipos e funções

**Principais Lições da Versão 5.0:**

- **Erros são valores em Go**: Não há exceções, erros são retornados como valores
- **Erros customizados melhoram semântica**: `ErrEstoqueInsuficiente` é mais descritivo que uma string genérica
- **Validação preventiva evita estados inválidos**: Verificar valores antes de modificar estado
- **Métodos devem comunicar falhas**: Retornar `error` permite que o chamador decida como tratar
- **Múltiplos valores de retorno são comuns**: `(resultadoDesejado, error)` é um padrão idiomático em Go
- **Propagação de erros é explícita**: Não há tratamento automático, cada nível decide se trata ou propaga
- **Testes devem cobrir cenários de erro**: Validar comportamento em situações anormais é tão importante quanto sucesso
- **Erros específicos facilitam debugging**: Saber exatamente o que falhou acelera correções
- **Regras de negócio devem ser validadas**: Estoque insuficiente é uma regra de negócio, não um bug
- **Nomenclatura clara previne erros**: `DiminuirQuantidade` é mais claro que `DiminiurQuantidade`

**Principais Lições da Versão 6.0:**

- **Múltiplas implementações de uma interface**: `RepositorioMemoria` e `RepositorioArquivo` implementam o mesmo contrato
- **Persistência em JSON é simples em Go**: Pacote `encoding/json` fornece `Marshal` e `Unmarshal`
- **Factory Functions padronizam criação**: `NovoProduto()` garante que produtos sempre tenham IDs válidos
- **Timestamps são úteis para IDs únicos**: `time.Now().UnixNano()` gera identificadores únicos sem conflitos
- **Interfaces permitem trocar implementações facilmente**: Mudar de memória para arquivo não requer alteração no serviço
- **Leitura e escrita de arquivos é direta**: `os.ReadFile` e `os.WriteFile` simplificam operações com arquivos
- **Tratamento de arquivo inexistente**: Retornar slice vazio quando arquivo não existe evita erros na primeira execução
- **Indentação melhora legibilidade do JSON**: `json.MarshalIndent` cria arquivos JSON formatados e fáceis de ler
- **Consistência de nomenclatura é importante**: Renomear para `RepositorioMemoria` mantém padrão com `RepositorioArquivo`
- **Padrão Factory encapsula complexidade**: Cliente não precisa saber como ID é gerado, apenas chama `NovoProduto()`

**Principais Lições da Versão 6.1:**

- **Hash determinístico resolve problema de duplicatas**: Usar hash do nome garante mesmo ID para mesmo produto
- **SHA256 é confiável para geração de IDs**: Hash criptográfico garante unicidade e consistência
- **IDs determinísticos facilitam comparação**: Produtos com mesmo nome sempre terão mesmo ID, evitando duplicatas
- **Conversão hexadecimal torna IDs legíveis**: 16 caracteres hex são suficientes e fáceis de visualizar
- **Separação de responsabilidades**: `gerarID()` recebe nome como parâmetro em vez de usar timestamp global
- **Testes devem usar mesmas práticas que código de produção**: Mock deve implementar interface completamente
- **Atualizar testes junto com código**: Refatorações no código devem refletir nos testes para manter consistência
- **Hash resolve trade-off entre simplicidade e flexibilidade**: Nome único + ID fixo = solução equilibrada

**Principais Lições da Versão 7.0:**

- **Concorrência exige proteção de estado**: Acesso simultâneo a slices pode causar race conditions
- **Mutex garante exclusão mútua**: `sync.Mutex` protege a região crítica
- **`Lock()`/`Unlock()` devem estar pareados**: `defer` reduz risco de esquecer o unlock
- **Leitura também precisa de proteção**: `Listar()` trava antes de copiar os dados
- **Cópia evita modificações externas**: Retornar uma cópia protege a estrutura interna
- **Mudanças pequenas geram robustez**: Poucas linhas de mutex evitam bugs difíceis de reproduzir

---

## 📄 Licença

Este é um projeto educacional de código aberto para fins de aprendizado.

---

**Última atualização:** Fevereiro 2026  
**Versão atual:** 7.0 - Concorrência no Repositório em Memória
