# Calculadora Básica em Go

Uma calculadora de linha de comando simples e interativa, desenvolvida em Go, que realiza operações matemáticas básicas (adição, subtração, multiplicação e divisão).

## 📋 Índice

- [Características](#-características)
- [Arquitetura do Projeto](#-arquitetura-do-projeto)
- [Lógica Implementada](#-lógica-implementada)
- [Organização do Código](#-organização-do-código)
- [Como Usar](#-como-usar)
- [Testes](#-testes)
- [Estrutura de Diretórios](#-estrutura-de-diretórios)

## ✨ Características

- ✅ Operações matemáticas básicas: `+`, `-`, `*`, `/`
- ✅ Suporte a números decimais
- ✅ Validação de entrada do usuário
- ✅ Detecção de divisão por zero
- ✅ Detecção de operações inválidas
- ✅ Interface interativa em loop
- ✅ Opção de sair a qualquer momento
- ✅ Tratamento robusto de erros
- ✅ Código totalmente comentado
- ✅ Testes unitários completos

## 🏗️ Arquitetura do Projeto

Este projeto segue os princípios de **Clean Architecture** e **Separation of Concerns** (Separação de Responsabilidades), garantindo um código:

- **Modular**: Cada parte tem uma responsabilidade específica
- **Testável**: Lógica de negócio isolada e fácil de testar
- **Manutenível**: Fácil de entender e modificar
- **Escalável**: Estrutura preparada para crescimento

### Camadas da Aplicação

```
┌─────────────────────────────────────┐
│         cmd/ (Camada de UI)         │
│  - main.go: Ponto de entrada        │
│  - app.go: Lógica da CLI            │
└─────────────────┬───────────────────┘
                  │
                  ▼
┌─────────────────────────────────────┐
│   internal/calculator (Lógica)      │
│  - calculator.go: Cálculos          │
│  - errors.go: Erros personalizados  │
│  - calculator_test.go: Testes       │
└─────────────────────────────────────┘
```

## 🧠 Lógica Implementada

### 1. **Pacote `calculator` (Lógica de Negócio)**

**Arquivo: `internal/calculator/calculator.go`**

A função principal `Calculate(a, b float64, op string)` implementa:

```go
func Calculate(a, b float64, op string) (float64, error)
```

**Lógica:**

- Recebe dois números (`a` e `b`) e uma operação (`op`)
- Usa `switch` para determinar a operação
- Para divisão, valida se `b == 0` antes de calcular
- Retorna o resultado ou um erro apropriado

**Arquivo: `internal/calculator/errors.go`**

Define erros personalizados:

- `ErrDivisionByZero`: Para divisão por zero
- `ErrInvalidOperation`: Para operações não suportadas

### 2. **Camada de Interface (CLI)**

**Arquivo: `cmd/main.go`**

O `main.go` foi mantido **extremamente enxuto** (apenas 11 linhas):

```go
func main() {
    if err := Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Erro ao ler entrada: %v\n", err)
        os.Exit(1)
    }
}
```

**Por que isso é importante?**

- O ponto de entrada deve ser simples e delegar responsabilidades
- Facilita testes e manutenção
- Segue o princípio de Single Responsibility

**Arquivo: `cmd/app.go`**

Contém toda a lógica da interface, separada em funções pequenas:

1. **`Run()`**: Orquestra a aplicação
   - Cria o scanner para ler entrada
   - Inicia o loop de processamento
   - Trata erros do scanner

2. **`printHeader()`**: Exibe cabeçalho
   - Mostra título e instruções

3. **`readInput(scanner, prompt)`**: Lê entrada genérica
   - Exibe prompt
   - Lê linha do usuário
   - Remove espaços em branco
   - Verifica comando "sair"
   - Retorna (input, continue)

4. **`readNumber(scanner, prompt)`**: Lê e valida números
   - Usa `readInput()` para ler
   - Converte string para float64
   - Valida a conversão
   - Retorna (número, continue)

5. **`processCalculation(scanner)`**: Processa uma operação completa
   - Lê primeiro número
   - Lê operação
   - Lê segundo número
   - Chama `displayResult()`
   - Retorna se deve continuar

6. **`displayResult(a, b, op)`**: Calcula e exibe resultado
   - Chama `calculator.Calculate()`
   - Exibe resultado ou erro

## 📁 Organização do Código

### Por que separamos em múltiplos arquivos?

#### ✅ **Vantagens da Separação:**

1. **Responsabilidade Única**
   - Cada arquivo tem um propósito claro
   - Facilita localizar código específico

2. **Reutilização**
   - Funções pequenas podem ser usadas em outros lugares
   - `readInput()` elimina código duplicado

3. **Testabilidade**
   - Lógica de negócio isolada em `calculator/`
   - Fácil escrever testes unitários

4. **Manutenibilidade**
   - Mudanças na UI não afetam a lógica
   - Mudanças na lógica não afetam a UI

5. **Legibilidade**
   - Funções pequenas são mais fáceis de entender
   - Código bem documentado

### Padrão Utilizado: **Table-Driven Tests**

Os testes utilizam o padrão recomendado em Go:

```go
tests := []struct {
    name     string
    a        float64
    b        float64
    op       string
    expected float64
    wantErr  error
}{
    // casos de teste aqui
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // executa teste
    })
}
```

**Vantagens:**

- Fácil adicionar novos casos
- Cada caso é um sub-teste independente
- Output claro de qual teste falhou

## 🚀 Como Usar

### Pré-requisitos

- Go 1.22.2 ou superior

### Executando a Calculadora

**Opção 1: A partir do diretório `cmd/`**

```bash
cd mini_go_projects/calculadoraBasica/cmd
go run .
```

**Opção 2: Especificando os arquivos**

```bash
cd mini_go_projects/calculadoraBasica/cmd
go run main.go app.go
```

**Opção 3: A partir da raiz do projeto**

```bash
cd mini_go_projects/calculadoraBasica
go run ./cmd
```

### Exemplo de Uso

```
=== Calculadora Básica ===
Digite 'sair' para encerrar

Digite o primeiro número: 10
Digite a operação (+, -, *, /): +
Digite o segundo número: 5
Resultado: 10.00 + 5.00 = 15.00

Digite o primeiro número: 20
Digite a operação (+, -, *, /): /
Digite o segundo número: 4
Resultado: 20.00 / 4.00 = 5.00

Digite o primeiro número: 10
Digite a operação (+, -, *, /): /
Digite o segundo número: 0
Erro: division by zero

Digite o primeiro número: 5
Digite a operação (+, -, *, /): %
Digite o segundo número: 3
Erro: invalid operation

Digite o primeiro número: sair
Encerrando...
```

### Compilando o Executável

```bash
cd mini_go_projects/calculadoraBasica/cmd
go build -o calculadora
./calculadora
```

## 🧪 Testes

### Executando os Testes

**Todos os testes:**

```bash
cd mini_go_projects/calculadoraBasica
go test ./...
```

**Apenas testes do calculator:**

```bash
go test ./internal/calculator
```

**Com output detalhado:**

```bash
go test ./internal/calculator -v
```

**Com cobertura:**

```bash
go test ./internal/calculator -cover
```

**Relatório de cobertura HTML:**

```bash
go test ./internal/calculator -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Casos de Teste Implementados

O arquivo `calculator_test.go` contém **20 casos de teste**:

#### `TestCalculate` (17 casos):

- ✅ **Soma**: positivos, negativos, decimais
- ✅ **Subtração**: positivos, resultado negativo, decimais
- ✅ **Multiplicação**: positivos, negativos, por zero, decimais
- ✅ **Divisão**: positivos, decimais, negativos, **por zero**
- ✅ **Operações Inválidas**: %, string vazia, caracteres desconhecidos

#### `TestCalculateEdgeCases` (4 casos):

- ✅ Zero dividido por número
- ✅ Números muito grandes (1e10)
- ✅ Números muito pequenos (0.0001)
- ✅ Subtração resultando em zero

### Por que usar Table-Driven Tests?

```go
// ❌ Forma tradicional (repetitiva)
func TestAdd(t *testing.T) {
    result, _ := Calculate(2, 3, "+")
    if result != 5 {
        t.Error("...")
    }
}

func TestSubtract(t *testing.T) {
    result, _ := Calculate(5, 3, "-")
    if result != 2 {
        t.Error("...")
    }
}

// ✅ Table-Driven (organizado e escalável)
tests := []struct{...}{...}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // teste aqui
    })
}
```

## 📂 Estrutura de Diretórios

```
calculadoraBasica/
├── go.mod                          # Módulo Go do projeto
├── README.md                       # Esta documentação
├── cmd/                            # Código executável (entry point)
│   ├── main.go                     # Ponto de entrada (11 linhas)
│   └── app.go                      # Lógica da CLI
└── internal/                       # Código interno (não exportável)
    └── calculator/                 # Pacote de cálculos
        ├── calculator.go           # Lógica de cálculo
        ├── calculator_test.go      # Testes unitários (20 casos)
        └── errors.go               # Erros personalizados
```

### Por que `internal/`?

O diretório `internal/` é uma convenção do Go que:

- Torna o pacote **não importável** de fora do projeto
- Protege o código interno de uso externo
- Evita dependências acidentais

## 🔧 Decisões de Design

### 1. **Uso de `bufio.Scanner` ao invés de `fmt.Scan`**

**Por quê?**

- `fmt.Scan` para no primeiro espaço em branco
- `bufio.Scanner` lê a linha completa
- Melhor controle sobre entrada do usuário
- Permite validação mais robusta

### 2. **Float64 para números**

**Por quê?**

- Suporta números decimais
- Precisão suficiente para calculadora básica
- Tipo padrão para cálculos em Go

### 3. **Erros Personalizados**

**Por quê?**

- Mais semântico que strings
- Permite verificação de tipo de erro
- Facilita tratamento específico

### 4. **Loop Infinito com Break**

**Por quê?**

- Permite múltiplas operações
- Usuário controla quando sair
- Mais interativo e prático

## 📝 Conceitos de Go Aplicados

1. **Packages e Modules**
   - Organização em pacotes (`main`, `calculator`)
   - `go.mod` para gerenciar dependências

2. **Error Handling**
   - Múltiplos retornos `(resultado, erro)`
   - Tratamento explícito de erros

3. **Structs e Table-Driven Tests**
   - Structs anônimos para casos de teste
   - Loop sobre slice de structs

4. **Sub-tests**
   - `t.Run()` para testes organizados
   - Output claro de falhas

5. **Float Comparison**
   - Função `floatEquals` para comparar com epsilon
   - Evita problemas de precisão

## 🎯 Boas Práticas Aplicadas

- ✅ Código comentado em português
- ✅ Nomes de funções descritivos
- ✅ Single Responsibility Principle
- ✅ DRY (Don't Repeat Yourself)
- ✅ Separation of Concerns
- ✅ Clean Architecture
- ✅ Error Handling adequado
- ✅ Testes abrangentes
- ✅ Documentação completa

## 🚦 Melhorias Futuras

Possíveis expansões do projeto:

- [ ] Operações avançadas (potência, raiz quadrada, módulo)
- [ ] Histórico de cálculos
- [ ] Interface gráfica (GUI)
- [ ] Salvar/carregar sessões
- [ ] Suporte a expressões (ex: "2 + 3 \* 4")
- [ ] Constantes matemáticas (π, e)
- [ ] Conversão de bases (binário, hexadecimal)

## 📄 Licença

Este projeto é parte de um repositório de mini projetos em Go para fins educacionais.

## 👨‍💻 Autor

Duarte Rodrigo Santos de Oliveira
Desenvolvido como parte do portfólio de projetos em Go.

---

**Dica:** Sempre use `go run .` ao invés de listar arquivos individualmente! 🚀
