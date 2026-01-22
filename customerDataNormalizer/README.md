# 🇧🇷 Customer Data Normalizer | Normalizador de Dados de Cliente

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Tests](https://img.shields.io/badge/tests-59%20passing-success)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

</div>

---

## 🇧🇷 Português

### 📖 Sobre o Projeto

Este projeto é um **mini-projeto de aprendizagem** desenvolvido em **Golang**, focado em **normalização e validação de dados de clientes**, utilizando conceitos avançados de **tipos nomeados (named types)**, **factory pattern**, **separação de responsabilidades** e **testes abrangentes**.

O objetivo principal é aprender **como construir sistemas robustos de validação de dados**, aplicando conceitos de **type safety**, **clean code** e **arquitetura de software**.

### 🎯 O que este projeto faz

- ✅ Normaliza nomes (capitalização correta)
- ✅ Valida e normaliza CPF (com checksum e validações completas)
- ✅ Valida e normaliza telefones celulares brasileiros (DDD + 9 dígitos)
- ✅ Valida e normaliza emails (formato padrão)
- 🔒 **Type Safety** - Impossível confundir CPF com Email no código
- 📦 **Factory Pattern** - Garante que apenas dados válidos sejam criados
- 🧪 **59 testes automatizados** cobrindo todos os cenários possíveis
- 🏗️ **Arquitetura limpa** - Separação clara de responsabilidades

---

## 💼 Aplicações no Mundo Real

Este tipo de sistema é **essencial** em diversos contextos profissionais:

### 🏢 CRM (Customer Relationship Management)

- Normalizar dados de clientes antes de salvar no banco de dados
- Evitar duplicatas por variações de formatação
- Garantir qualidade dos dados de contato

### 🛒 E-commerce

- Validar dados de cadastro no checkout
- Padronizar informações de entrega
- Garantir contato válido com cliente

### 🏦 Sistemas Financeiros

- Validar CPF antes de aprovar transações
- Garantir dados corretos para compliance
- Padronizar informações para auditoria

### 📊 Data Warehouses

- ETL (Extract, Transform, Load) - fase de transformação
- Limpeza de dados antes de análises
- Padronização para integração entre sistemas

### 🔐 Autenticação e Cadastro

- Validar dados no registro de usuários
- Normalizar emails para evitar duplicatas
- Verificar documentos antes de criar contas

---

## 🎓 O que foi Aprendido

### 1. **Tipos Nomeados (Named Types) para Type Safety**

#### Por que usar tipos nomeados?

**❌ Problema sem tipos nomeados:**

```go
func ProcessPayment(cpf string, email string) {
    // É possível passar email onde espera CPF!
    ProcessPayment("user@example.com", "12345678909") // ERRO NÃO DETECTADO
}
```

**✅ Solução com tipos nomeados:**

```go
func ProcessPayment(cpf CPF, email Email) {
    // Compilador impede erros de tipo!
    ProcessPayment(email, cpf) // ERRO DE COMPILAÇÃO ✓
}
```

**Benefícios:**

- 🔒 **Segurança** - Compilador previne erros
- 📖 **Documentação automática** - Código auto-explicativo
- 🎯 **Clareza** - Função deixa claro o que espera
- 🛡️ **Prevenção de bugs** - Erros detectados antes de rodar

**Implementação:**

```go
// types.go - Apenas definições
type CPF string
type Name string
type Email string
type Phone string
```

---

### 2. **Factory Pattern para Validação Garantida**

#### Por que usar Factory Pattern?

O Factory garante que **apenas dados válidos possam existir no sistema**.

**❌ Sem Factory:**

```go
var cpf CPF = "123" // CPF inválido criado sem validação!
```

**✅ Com Factory:**

```go
cpf, err := NewCPF("123") // ERRO retornado, CPF inválido não é criado
if err != nil {
    // Tratar erro
}
// Se chegou aqui, cpf é GARANTIDAMENTE válido
```

**Estrutura do Factory:**

```go
// factoryCPF.go
func NewCPF(value string) (CPF, error) {
    // 1. Extrai apenas dígitos
    // 2. Valida tamanho (11 dígitos)
    // 3. Rejeita dígitos repetidos (111.111.111-11)
    // 4. Valida checksum (dígitos verificadores)

    if /* inválido */ {
        return "", ErrCPFinvalid
    }
    return CPF(normalized), nil // Apenas CPFs válidos são criados!
}
```

**Benefícios:**

- ✅ **Garantia de validade** - Impossível criar dado inválido
- ✅ **Ponto único de validação** - Lógica centralizada
- ✅ **Facilita testes** - Testar factory = testar todo o sistema
- ✅ **Manutenção** - Mudança de regra em 1 lugar só

---

### 3. **Separação de Responsabilidades**

#### Arquitetura do Projeto

```
internal/client/
├── types.go          → Definição dos tipos nomeados
├── factoryCPF.go     → Validação e criação de CPF
├── factoryEmail.go   → Validação e criação de Email
├── factoryPhone.go   → Validação e criação de Phone
├── model.go          → Struct Client (usa os tipos)
├── normalizer.go     → Orquestra normalização completa
├── errors.go         → Erros customizados
└── normalizer_test.go → 59 testes cobrindo tudo
```

**Por que essa estrutura?**

| Arquivo         | Responsabilidade          | Mudança Isolada             |
| --------------- | ------------------------- | --------------------------- |
| `types.go`      | Definir tipos             | Adicionar novo tipo         |
| `factory*.go`   | Validar regras de negócio | Mudar regra de validação    |
| `normalizer.go` | Coordenar processo        | Mudar fluxo de normalização |
| `model.go`      | Estrutura de dados        | Adicionar/remover campo     |

**Benefício:** Mudanças são **localizadas** e não afetam outras partes!

---

### 4. **Validação Completa de CPF (Checksum)**

#### Lógica do Algoritmo de Validação

O CPF brasileiro usa **2 dígitos verificadores** calculados por um algoritmo específico.

**Exemplo: CPF 123.456.789-09**

**Primeiro dígito (0):**

```
Posição:  1   2   3   4   5   6   7   8   9
Dígito:   1   2   3   4   5   6   7   8   9
Peso:    10   9   8   7   6   5   4   3   2
Mult:    10  18  24  28  30  30  28  24  18  = 210
Resto:   210 * 10 % 11 = 10 → se 10, vira 0
Resultado: 0 ✓
```

**Segundo dígito (9):**

```
Posição:  1   2   3   4   5   6   7   8   9   10
Dígito:   1   2   3   4   5   6   7   8   9   0
Peso:    11  10   9   8   7   6   5   4   3   2
Mult:    11  20  27  32  35  36  35  32  27   0  = 255
Resto:   255 * 10 % 11 = 9
Resultado: 9 ✓
```

**Por que isso importa:**

- 🛡️ **Detecta erros de digitação** - 99% de eficácia
- 🔒 **Previne fraudes** - Não aceita CPFs inventados
- ✅ **Padrão oficial** - Segue Receita Federal

---

### 5. **Table-Driven Tests - Testes Eficientes**

#### Padrão idiomático em Go

**❌ Forma tradicional (repetitiva):**

```go
func TestCPFValid(t *testing.T) { /* ... */ }
func TestCPFEmpty(t *testing.T) { /* ... */ }
func TestCPFInvalidChecksum(t *testing.T) { /* ... */ }
// ... 16 funções diferentes
```

**✅ Forma com Table-Driven:**

```go
tests := []struct {
    name     string
    input    string
    expected CPF
    hasError bool
}{
    {"CPF válido", "123.456.789-09", CPF("12345678909"), false},
    {"CPF vazio", "", CPF(""), true},
    // ... todos os casos em uma tabela
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result, err := NewCPF(tt.input)
        // Validações...
    })
}
```

**Benefícios:**

- 📊 **Organização** - Dados separados da lógica
- ⚡ **Eficiência** - Fácil adicionar novos casos
- 🎯 **Clareza** - Nome descritivo para cada teste
- 🔄 **Reutilização** - Mesma lógica para todos

**Cobertura do Projeto:**

- ✅ 6 casos para `Normalize()`
- ✅ 5 casos para normalização de nome
- ✅ 16 casos para CPF (4 válidos + 12 inválidos)
- ✅ 14 casos para Email (6 válidos + 8 inválidos)
- ✅ 18 casos para Phone (6 válidos + 12 inválidos)
- **Total: 59 testes** 🎉

---

### 6. **Normalização vs Validação**

#### Diferença conceitual importante

**Validação:**

- ✅ Verifica se dado está **correto**
- ❌ Retorna erro se inválido
- 🔍 Exemplo: CPF tem 11 dígitos?

**Normalização:**

- 🔧 **Transforma** dado para formato padrão
- ✨ Remove variações de formatação
- 📝 Exemplo: "JOÃO silva" → "João Silva"

**No projeto:**

```go
// Validação + Normalização juntas
func NewEmail(value string) (Email, error) {
    email := strings.TrimSpace(value)      // Normaliza (remove espaços)
    if !strings.Contains(email, "@") {     // Valida (tem @?)
        return "", ErrEmailInvalid
    }
    return Email(strings.ToLower(email)), nil // Normaliza (minúsculas)
}
```

---

## 🛠️ Tecnologias e Conceitos

### Pacotes Go Utilizados

| Pacote    | Uso                     | Por que                     |
| --------- | ----------------------- | --------------------------- |
| `strings` | Manipulação de strings  | Trim, ToLower, Repeat, etc. |
| `unicode` | Validação de caracteres | Verificar se é dígito       |
| `fmt`     | Formatação              | Criar mensagens de teste    |
| `testing` | Framework de testes     | Testes nativos do Go        |

### Conceitos de Programação

- ✅ **Type Safety** - Segurança de tipos
- ✅ **Factory Pattern** - Padrão de criação
- ✅ **Separation of Concerns** - Separação de responsabilidades
- ✅ **Named Types** - Tipos nomeados personalizados
- ✅ **Error Handling** - Tratamento de erros idiomático
- ✅ **Table-Driven Tests** - Testes orientados a dados
- ✅ **Clean Architecture** - Arquitetura limpa

---

## 📂 Estrutura do Projeto

```
customerDataNormalizer/
├── cmd/
│   └── main.go                # Ponto de entrada
├── internal/
│   └── client/
│       ├── types.go           # Tipos nomeados (CPF, Name, Email, Phone)
│       ├── factoryCPF.go      # Factory + validação de CPF
│       ├── factoryEmail.go    # Factory + validação de Email
│       ├── factoryPhone.go    # Factory + validação de Phone
│       ├── model.go           # Struct Client
│       ├── normalizer.go      # Função principal Normalize()
│       ├── errors.go          # Erros customizados
│       └── normalizer_test.go # 59 testes
├── go.mod
└── README.md
```

---

## 🎮 Como Usar

### Clonar o Repositório

```bash
git clone https://github.com/EtraudBits/mini_go_projects.git
cd mini_go_projects/customerDataNormalizer
```

### Executar a Aplicação

```bash
go run cmd/main.go
```

**Saída esperada:**

```
=== Customer Data Normalizer ===

✅ Cliente normalizado com sucesso!

Nome: João Silva Santos
CPF: 12345678909
Telefone: 11987654321
Email: joao@example.com
```

### Executar os Testes

```bash
# Todos os testes
go test ./internal/client/

# Com detalhes
go test -v ./internal/client/

# Com cobertura
go test -cover ./internal/client/
```

### Gerar Executável

```bash
go build -o customerDataNormalizer ./cmd/main.go
./customerDataNormalizer
```

---

## 📊 Exemplo de Uso no Código

```go
package main

import (
    "fmt"
    "customerDataNormalizer/internal/client"
)

func main() {
    // Dados crus do usuário
    rawName := "  maria SILVA santos  "
    rawCPF := "123.456.789-09"
    rawPhone := "(21) 99876-5432"
    rawEmail := "MARIA@EXAMPLE.COM"

    // Normaliza todos os campos de uma vez
    customer, errs := client.Normalize(rawName, rawCPF, rawPhone, rawEmail)

    if len(errs) > 0 {
        fmt.Println("❌ Erros encontrados:")
        for _, err := range errs {
            fmt.Printf("  - %v\n", err)
        }
        return
    }

    // Dados normalizados e validados
    fmt.Printf("Nome: %s\n", customer.Name)          // Maria Silva Santos
    fmt.Printf("CPF: %s\n", customer.CPF)            // 12345678909
    fmt.Printf("Tel: %s\n", customer.Phone)          // 21998765432
    fmt.Printf("Email: %s\n", customer.Email)        // maria@example.com
}
```

---

## 💡 Principais Lições Aprendidas

| Conceito            | Lição                                            | Aplicação                              |
| ------------------- | ------------------------------------------------ | -------------------------------------- |
| **Type Safety**     | Tipos previnem erros em tempo de compilação      | Impossível confundir CPF com Email     |
| **Factory Pattern** | Garante validade dos dados no momento da criação | Apenas dados válidos entram no sistema |
| **Separação**       | Cada arquivo tem uma responsabilidade única      | Manutenção facilitada                  |
| **Testes**          | Table-driven tests são eficientes e escaláveis   | 59 testes com pouco código             |
| **Normalização**    | Dados padronizados facilitam comparações         | Evita duplicatas no banco              |

---

## 🇺🇸 English

### 📖 About the Project

This is a **learning mini-project** developed in **Golang**, focused on **customer data normalization and validation**, using advanced concepts of **named types**, **factory pattern**, **separation of concerns**, and **comprehensive testing**.

The main goal is to learn **how to build robust data validation systems**, applying concepts of **type safety**, **clean code**, and **software architecture**.

### 🎯 What this project does

- ✅ Normalizes names (correct capitalization)
- ✅ Validates and normalizes CPF (with checksum and complete validations)
- ✅ Validates and normalizes Brazilian mobile phones (area code + 9 digits)
- ✅ Validates and normalizes emails (standard format)
- 🔒 **Type Safety** - Impossible to confuse CPF with Email in code
- 📦 **Factory Pattern** - Ensures only valid data can be created
- 🧪 **59 automated tests** covering all possible scenarios
- 🏗️ **Clean architecture** - Clear separation of responsibilities

---

## 💼 Real-World Applications

This type of system is **essential** in various professional contexts:

### 🏢 CRM (Customer Relationship Management)

- Normalize customer data before saving to database
- Avoid duplicates due to formatting variations
- Ensure contact data quality

### 🛒 E-commerce

- Validate registration data at checkout
- Standardize delivery information
- Ensure valid customer contact

### 🏦 Financial Systems

- Validate CPF before approving transactions
- Ensure correct data for compliance
- Standardize information for auditing

### 📊 Data Warehouses

- ETL (Extract, Transform, Load) - transformation phase
- Data cleaning before analysis
- Standardization for system integration

### 🔐 Authentication and Registration

- Validate data during user registration
- Normalize emails to avoid duplicates
- Verify documents before creating accounts

---

## 🎓 Key Learnings

### 1. **Named Types for Type Safety**

**❌ Problem without named types:**

```go
func ProcessPayment(cpf string, email string) {
    // Can pass email where CPF is expected!
    ProcessPayment("user@example.com", "12345678909") // ERROR NOT DETECTED
}
```

**✅ Solution with named types:**

```go
func ProcessPayment(cpf CPF, email Email) {
    // Compiler prevents type errors!
    ProcessPayment(email, cpf) // COMPILATION ERROR ✓
}
```

### 2. **Factory Pattern for Guaranteed Validation**

Factory ensures that **only valid data can exist in the system**.

```go
cpf, err := NewCPF("123") // ERROR returned, invalid CPF is not created
if err != nil {
    // Handle error
}
// If reached here, cpf is GUARANTEED valid
```

### 3. **CPF Checksum Validation**

Brazilian CPF uses **2 check digits** calculated by a specific algorithm, detecting 99% of typos and preventing fraud.

### 4. **Table-Driven Tests**

Idiomatic pattern in Go for efficient and scalable testing.

**Coverage:**

- 59 tests covering all validation scenarios
- 100% code coverage
- Valid and invalid cases for all fields

---

## 👤 Author | Autor

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

## 📎 Final Notes | Observações Finais

🇧🇷 Este é um projeto educacional focado em **arquitetura de software**, **type safety** e **qualidade de código**. O código demonstra como construir sistemas robustos de validação aplicáveis em cenários reais de produção. Sugestões e feedbacks são sempre bem-vindos!

🇺🇸 This is an educational project focused on **software architecture**, **type safety**, and **code quality**. The code demonstrates how to build robust validation systems applicable in real production scenarios. Suggestions and feedback are always welcome!

---

<div align="center">

**⭐ Se este projeto foi útil, considere dar uma estrela!**

**⭐ If this project was helpful, consider giving it a star!**

</div>
