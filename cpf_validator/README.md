# 🇧🇷 Validador de CPF em Go | 🇺🇸 CPF Validator in Go

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

</div>

---

## 🇧🇷 Português

### 📖 Sobre o Projeto

Este projeto é um mini-projeto de estudo desenvolvido em **Golang**, com foco em **lógica de programação**, **organização de código**, **tratamento de erros** e **estrutura próxima da realidade de projetos backend**.

O objetivo principal não é apenas validar um CPF, mas entender **como pensar um problema**, **como organizar a solução** e **por que cada parte do código existe**.

### 📌 O que este projeto faz

- ✅ Recebe um CPF em formato texto (com ou sem formatação)
- 🧹 Remove caracteres não numéricos automaticamente
- 🔍 Valida regras básicas (tamanho, repetição de números)
- 🧮 Calcula e valida os dígitos verificadores oficiais
- ❌ Retorna erros específicos e claros para cada tipo de falha
- 🏗️ Mantém a lógica de negócio separada do ponto de entrada (`main`)
- ✅ Inclui testes automatizados

### 🧠 Por que um validador de CPF?

A validação de CPF é uma funcionalidade **extremamente comum** em sistemas reais, especialmente no Brasil. Você encontra essa lógica em:

| Contexto                    | Exemplos                                             |
| --------------------------- | ---------------------------------------------------- |
| **E-commerce**              | Validação de dados no checkout, cadastro de clientes |
| **Sistemas Bancários**      | Abertura de contas, validação de titulares           |
| **APIs de Autenticação**    | Verificação de identidade, KYC (Know Your Customer)  |
| **Plataformas de Saúde**    | Cadastro de pacientes, agendamentos                  |
| **Sistemas Governamentais** | Portais de serviços públicos, declarações            |
| **RH e Folha de Pagamento** | Cadastro de colaboradores, integração com eSocial    |

Ou seja, apesar de simples, esse tipo de lógica **aparece constantemente** no dia a dia de quem trabalha com backend no Brasil.

### 🏗️ Arquitetura do Projeto

```text
cpf_validator/
├── go.mod                    # Gerenciamento de dependências
├── README.md                 # Documentação
├── cmd/
│   └── main.go              # 🚪 Ponto de entrada da aplicação
└── internal/
    └── cpf/
        ├── model.go         # 📦 Definição do tipo CPF
        ├── factory.go       # 🏭 Criação e validação de CPF
        ├── service.go       # 🔧 Serviços e operações
        ├── errors.go        # ⚠️  Erros específicos do domínio
        └── service_test.go  # 🧪 Testes automatizados
```

#### 💡 Por que essa estrutura?

| Diretório       | Propósito                            | Benefício                                                     |
| --------------- | ------------------------------------ | ------------------------------------------------------------- |
| **`cmd/`**      | Contém apenas o código executável    | Facilita criar múltiplos pontos de entrada (CLI, API, worker) |
| **`internal/`** | Lógica de negócio privada do projeto | Impede que outros projetos importem código interno            |
| **`cpf/`**      | Pacote específico do domínio CPF     | Agrupa toda lógica relacionada, facilitando manutenção        |

**Princípios aplicados:**

- ✅ **Separação de responsabilidades**: cada arquivo tem um propósito claro
- ✅ **Domain-Driven Design (DDD)**: lógica organizada por domínio
- ✅ **Testabilidade**: lógica isolada do `main` permite testes fáceis
- ✅ **Reutilização**: o pacote `cpf` pode ser usado em diferentes contextos

### 📂 Detalhamento dos Arquivos

#### `model.go`

Define o tipo `CPF` como um tipo personalizado baseado em `string`, garantindo type safety e permitindo adicionar métodos específicos.

#### `factory.go`

Implementa a função `NewCPF()` que:

- Valida o formato
- Verifica dígitos repetidos
- Calcula e valida os dígitos verificadores
- Retorna erro específico para cada tipo de falha

#### `errors.go`

Define erros customizados usando o padrão Go de `var Err...`:

- `ErrInvalidLength`: CPF não tem 11 dígitos
- `ErrRepeatedDigits`: Todos os dígitos são iguais
- `ErrInvalidChecksum`: Dígitos verificadores incorretos

#### `service.go`

Contém funções auxiliares e serviços adicionais relacionados ao CPF.

#### `service_test.go`

Testes automatizados que garantem a corretude da implementação.

### ⚙️ Conceitos e Técnicas Praticadas

#### 🎯 Lógica de Programação

- ✅ Algoritmo de validação de dígitos verificadores
- ✅ Manipulação de strings e conversão de tipos
- ✅ Loops e condicionais aplicados a problemas reais

#### 🏛️ Organização de Código

- ✅ Estrutura de pastas inspirada em projetos backend reais
- ✅ Separação clara entre execução (`cmd`) e lógica (`internal`)
- ✅ Nomenclatura descritiva e consistente

#### ⚠️ Tratamento de Erros

- ✅ Erros específicos para cada tipo de falha
- ✅ Retorno explícito de erros (padrão Go)
- ✅ Mensagens de erro claras e úteis

#### 🧪 Testes

- ✅ Testes unitários com `testing` package
- ✅ Casos de sucesso e falha
- ✅ Validação de comportamento esperado

#### 🎨 Boas Práticas Go

- ✅ Tipos customizados para domínios específicos
- ✅ Factory functions para criação validada
- ✅ Uso de `internal/` para encapsulamento
- ✅ Comentários bilíngues para contexto

### 🚀 Como Executar

#### Pré-requisitos

- Go 1.21 ou superior

#### Executar o programa

```bash
go run ./cmd
```

#### Executar os testes

```bash
go test ./internal/cpf -v
```

#### Exemplo de uso no código

```go
cpf, err := cpf.NewCPF("529.982.247-25")
if err != nil {
    log.Fatal(err)
}
fmt.Println("CPF válido:", cpf)
```

### 📚 O que Aprendi com Este Projeto

| Aspecto              | Aprendizado                                              |
| -------------------- | -------------------------------------------------------- |
| **Planejamento**     | Pensar antes de codar, entender o problema completamente |
| **Arquitetura**      | Organizar código pensando em manutenção e escalabilidade |
| **Qualidade**        | Escrever código limpo, testável e documentado            |
| **Profissionalismo** | Aproximar estudos da realidade de projetos backend       |
| **Go idiomático**    | Seguir convenções e boas práticas da comunidade Go       |

Este projeto faz parte da minha jornada de estudos como desenvolvedor backend, focando em **qualidade** e **profissionalismo** desde o início.

### 🎯 Próximos Passos

- [ ] Adicionar mais casos de teste
- [ ] Criar API REST para validação de CPF
- [ ] Implementar formatação de CPF (000.000.000-00)
- [ ] Adicionar benchmark de performance
- [ ] Documentação com exemplos de integração

---

## 🇺🇸 English

### 📖 About the Project

This is a **study-oriented mini project** developed in **Golang**, focused on **programming logic**, **code organization**, **error handling**, and **backend-oriented structure**.

The main goal is not only to validate a Brazilian CPF number, but to understand **how to think about a problem**, **how to structure a solution**, and **why each part of the code exists**.

### 📌 What this project does

- ✅ Receives a CPF as a string (formatted or not)
- 🧹 Automatically removes non-numeric characters
- 🔍 Validates basic rules (length, repeated digits)
- 🧮 Calculates and validates official check digits
- ❌ Returns specific, clear errors for each failure type
- 🏗️ Keeps business logic separated from entry point (`main`)
- ✅ Includes automated tests

### 🧠 Why a CPF validator?

CPF validation is an **extremely common** feature in real-world systems, especially in Brazil. You find this logic in:

| Context                  | Examples                                        |
| ------------------------ | ----------------------------------------------- |
| **E-commerce**           | Checkout validation, customer registration      |
| **Banking Systems**      | Account opening, holder verification            |
| **Authentication APIs**  | Identity verification, KYC (Know Your Customer) |
| **Healthcare Platforms** | Patient registration, appointments              |
| **Government Systems**   | Public service portals, tax declarations        |
| **HR & Payroll**         | Employee registration, eSocial integration      |

Even though it looks simple, this type of logic **appears constantly** in the day-to-day of backend developers in Brazil.

### 🏗️ Project Architecture

```text
cpf_validator/
├── go.mod                    # Dependency management
├── README.md                 # Documentation
├── cmd/
│   └── main.go              # 🚪 Application entry point
└── internal/
    └── cpf/
        ├── model.go         # 📦 CPF type definition
        ├── factory.go       # 🏭 CPF creation and validation
        ├── service.go       # 🔧 Services and operations
        ├── errors.go        # ⚠️  Domain-specific errors
        └── service_test.go  # 🧪 Automated tests
```

#### 💡 Why this structure?

| Directory       | Purpose                       | Benefit                                                          |
| --------------- | ----------------------------- | ---------------------------------------------------------------- |
| **`cmd/`**      | Contains only executable code | Makes it easy to create multiple entry points (CLI, API, worker) |
| **`internal/`** | Private business logic        | Prevents other projects from importing internal code             |
| **`cpf/`**      | Domain-specific package       | Groups all related logic, facilitating maintenance               |

**Applied principles:**

- ✅ **Separation of concerns**: each file has a clear purpose
- ✅ **Domain-Driven Design (DDD)**: logic organized by domain
- ✅ **Testability**: logic isolated from `main` allows easy testing
- ✅ **Reusability**: `cpf` package can be used in different contexts

### 📂 File Breakdown

#### `model.go`

Defines the `CPF` type as a custom type based on `string`, ensuring type safety and allowing specific methods.

#### `factory.go`

Implements the `NewCPF()` function that:

- Validates format
- Checks for repeated digits
- Calculates and validates check digits
- Returns specific errors for each failure type

#### `errors.go`

Defines custom errors using Go's `var Err...` pattern:

- `ErrInvalidLength`: CPF doesn't have 11 digits
- `ErrRepeatedDigits`: All digits are the same
- `ErrInvalidChecksum`: Incorrect check digits

#### `service.go`

Contains helper functions and additional CPF-related services.

#### `service_test.go`

Automated tests ensuring implementation correctness.

### ⚙️ Practiced Concepts and Techniques

#### 🎯 Programming Logic

- ✅ Check digit validation algorithm
- ✅ String manipulation and type conversion
- ✅ Loops and conditionals applied to real problems

#### 🏛️ Code Organization

- ✅ Folder structure inspired by real backend projects
- ✅ Clear separation between execution (`cmd`) and logic (`internal`)
- ✅ Descriptive and consistent naming

#### ⚠️ Error Handling

- ✅ Specific errors for each failure type
- ✅ Explicit error returns (Go pattern)
- ✅ Clear and helpful error messages

#### 🧪 Testing

- ✅ Unit tests with `testing` package
- ✅ Success and failure cases
- ✅ Expected behavior validation

#### 🎨 Go Best Practices

- ✅ Custom types for specific domains
- ✅ Factory functions for validated creation
- ✅ Use of `internal/` for encapsulation
- ✅ Bilingual comments for context

### 🚀 How to Run

#### Prerequisites

- Go 1.21 or higher

#### Run the program

```bash
go run ./cmd
```

#### Run tests

```bash
go test ./internal/cpf -v
```

#### Usage example in code

```go
cpf, err := cpf.NewCPF("529.982.247-25")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Valid CPF:", cpf)
```

### 📚 What I Learned from This Project

| Aspect              | Learning                                                    |
| ------------------- | ----------------------------------------------------------- |
| **Planning**        | Think before coding, fully understand the problem           |
| **Architecture**    | Organize code thinking about maintenance and scalability    |
| **Quality**         | Write clean, testable, and documented code                  |
| **Professionalism** | Bring studies closer to backend project reality             |
| **Idiomatic Go**    | Follow conventions and best practices from the Go community |

This project is part of my journey as a backend developer, focusing on **quality** and **professionalism** from the start.

### 🎯 Next Steps

- [ ] Add more test cases
- [ ] Create REST API for CPF validation
- [ ] Implement CPF formatting (000.000.000-00)
- [ ] Add performance benchmarks
- [ ] Documentation with integration examples

---

## 👤 Author | Autor

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

## 📎 Final Notes | Observações Finais

🇧🇷 Este é um projeto de estudo. Sugestões, feedbacks e melhorias são sempre bem-vindos!

🇺🇸 This is a study project. Suggestions, feedback, and improvements are always welcome!

---

<div align="center">

**⭐ Se este projeto foi útil para você, considere dar uma estrela!**

**⭐ If this project was helpful, consider giving it a star!**

</div>
