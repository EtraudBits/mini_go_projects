# 📧 Validador de Email | Email Validator

---

## 🇧🇷 Português

### 📝 Descrição

Um validador de email simples e educacional implementado em Go, focado em lógica clara e previsibilidade. Este projeto demonstra boas práticas de arquitetura em Go, separação de responsabilidades e validação de dados.

### ⚙️ Regras de Validação

Um email será considerado **válido** se:

1. ✅ **Não estiver vazio**
2. ✅ **Contiver exatamente um `@`**
3. ✅ **Tiver texto antes do `@`** (local part)
4. ✅ **Tiver domínio após o `@`**
5. ✅ **Domínio contiver pelo menos um ponto (`.`)**
6. ✅ **Não começar nem terminar com `@` ou `.`**

### ⚠️ Importante

**Não validamos 100% do RFC 5322.**  
O foco aqui é **lógica + previsibilidade**, como em muitos sistemas reais. Validar completamente o padrão RFC é extremamente complexo e, na prática, a maioria dos sistemas usa validações simplificadas como esta.

### 🏗️ Arquitetura do Projeto

```
validador_de_email/
├── go.mod                    # Define o módulo Go
├── README.md                 # Documentação
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
└── internal/
    └── email/
        ├── validador.go     # Lógica de validação (core)
        ├── errors.go        # Definição de erros customizados
        └── validador_test.go # Testes unitários
```

#### 📂 Estrutura Explicada

- **`cmd/main.go`**: Ponto de entrada da aplicação. Aqui você interage com o usuário e chama a lógica de negócio.

- **`internal/email/`**: Pacote interno que contém toda a lógica de validação.
  - **`validador.go`**: Contém a função `Validate()` que implementa todas as regras de validação.
  - **`errors.go`**: Define erros customizados reutilizáveis (convenção Go: nomes começam com `Err`).
  - **`validador_test.go`**: Testes unitários para garantir que a validação funciona corretamente.

### 💡 Conceitos de Go Aplicados

#### 1. **Módulos Go (`go.mod`)**

```go
module validador_email

go 1.22.2
```

- Define o **nome do módulo**
- Permite importar pacotes internos como `validador_email/internal/email`
- Gerencia dependências do projeto

#### 2. **Pacotes Internos (`internal/`)**

- Convenção Go: código em `internal/` **não pode ser importado por outros projetos**
- Garante encapsulamento e controle sobre o que é público

#### 3. **Erros Customizados**

```go
var (
    ErrEmptyEmail = errors.New("email não pode ser vazio")
    ErrIvalidAtSymbol = errors.New("email deve conter apenas um @")
    // ...
)
```

- Erros são **variáveis reutilizáveis**
- Convenção: nomes começam com `Err`
- Mensagens claras facilitam debugging

#### 4. **Manipulação de Strings**

```go
strings.Count(email, "@")      // Conta ocorrências
strings.Split(email, "@")      // Divide string
strings.Contains(domain, ".")  // Verifica presença
```

### 🔍 Lógica de Validação (Passo a Passo)

```go
func Validate(email string) error {
    // 1. Verifica se está vazio
    if email == "" {
        return ErrEmptyEmail
    }

    // 2. Verifica se tem exatamente um @
    if strings.Count(email, "@") != 1 {
        return ErrIvalidAtSymbol
    }

    // 3. Divide o email em local e domínio
    parts := strings.Split(email, "@")
    local := parts[0]   // antes do @
    domain := parts[1]  // depois do @

    // 4. Verifica se há texto antes do @
    if local == "" {
        return ErrInvalidLocalPart
    }

    // 5. Verifica se há domínio após o @
    if domain == "" {
        return ErrIvalidDomain
    }

    // 6. Verifica se o domínio tem pelo menos um ponto
    if !strings.Contains(domain, ".") {
        return ErrIvalidDotDomain
    }

    return nil // ✅ Email válido!
}
```

### 🚀 Como Usar

1. **Clone o projeto**

```bash
cd validador_de_email
```

2. **Execute o programa**

```bash
go run cmd/main.go
```

3. **Execute os testes**

```bash
go test ./internal/email -v
```

### 📚 O Que Aprendi

✅ **Arquitetura limpa em Go**: separação entre `cmd` (entrada) e `internal` (lógica)  
✅ **Gestão de módulos**: usar `go.mod` para organizar imports  
✅ **Tratamento de erros**: criar erros customizados e descritivos  
✅ **Manipulação de strings**: usar o pacote `strings` da biblioteca padrão  
✅ **Testes unitários**: garantir que a lógica funciona em diversos cenários  
✅ **Convenções Go**: nomenclatura, estrutura de pastas, e boas práticas

### 🎯 Próximos Passos

- [ ] Adicionar validação de caracteres especiais
- [ ] Implementar lista de domínios bloqueados
- [ ] Criar interface CLI interativa
- [ ] Adicionar suporte para validação em lote

---

## 👤 Autor

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

## 🇬🇧 English

### 📝 Description

A simple and educational email validator implemented in Go, focused on clear logic and predictability. This project demonstrates Go architecture best practices, separation of concerns, and data validation.

### ⚙️ Validation Rules

An email will be considered **valid** if:

1. ✅ **Not empty**
2. ✅ **Contains exactly one `@`**
3. ✅ **Has text before the `@`** (local part)
4. ✅ **Has a domain after the `@`**
5. ✅ **Domain contains at least one dot (`.`)**
6. ✅ **Does not start or end with `@` or `.`**

### ⚠️ Important

**We don't validate 100% of RFC 5322.**  
The focus here is **logic + predictability**, as in many real-world systems. Fully validating the RFC standard is extremely complex, and in practice, most systems use simplified validations like this one.

### 🏗️ Project Architecture

```
validador_de_email/
├── go.mod                    # Defines the Go module
├── README.md                 # Documentation
├── cmd/
│   └── main.go              # Application entry point
└── internal/
    └── email/
        ├── validador.go     # Validation logic (core)
        ├── errors.go        # Custom error definitions
        └── validador_test.go # Unit tests
```

#### 📂 Structure Explained

- **`cmd/main.go`**: Application entry point. Here you interact with the user and call business logic.

- **`internal/email/`**: Internal package containing all validation logic.
  - **`validador.go`**: Contains the `Validate()` function that implements all validation rules.
  - **`errors.go`**: Defines reusable custom errors (Go convention: names start with `Err`).
  - **`validador_test.go`**: Unit tests to ensure validation works correctly.

### 💡 Go Concepts Applied

#### 1. **Go Modules (`go.mod`)**

```go
module validador_email

go 1.22.2
```

- Defines the **module name**
- Allows importing internal packages like `validador_email/internal/email`
- Manages project dependencies

#### 2. **Internal Packages (`internal/`)**

- Go convention: code in `internal/` **cannot be imported by other projects**
- Ensures encapsulation and control over what is public

#### 3. **Custom Errors**

```go
var (
    ErrEmptyEmail = errors.New("email cannot be empty")
    ErrIvalidAtSymbol = errors.New("email must contain only one @")
    // ...
)
```

- Errors are **reusable variables**
- Convention: names start with `Err`
- Clear messages facilitate debugging

#### 4. **String Manipulation**

```go
strings.Count(email, "@")      // Count occurrences
strings.Split(email, "@")      // Split string
strings.Contains(domain, ".")  // Check presence
```

### 🔍 Validation Logic (Step by Step)

```go
func Validate(email string) error {
    // 1. Check if empty
    if email == "" {
        return ErrEmptyEmail
    }

    // 2. Check if has exactly one @
    if strings.Count(email, "@") != 1 {
        return ErrIvalidAtSymbol
    }

    // 3. Split email into local and domain
    parts := strings.Split(email, "@")
    local := parts[0]   // before @
    domain := parts[1]  // after @

    // 4. Check if there's text before @
    if local == "" {
        return ErrInvalidLocalPart
    }

    // 5. Check if there's a domain after @
    if domain == "" {
        return ErrIvalidDomain
    }

    // 6. Check if domain has at least one dot
    if !strings.Contains(domain, ".") {
        return ErrIvalidDotDomain
    }

    return nil // ✅ Valid email!
}
```

### 🚀 How to Use

1. **Clone the project**

```bash
cd validador_de_email
```

2. **Run the program**

```bash
go run cmd/main.go
```

3. **Run tests**

```bash
go test ./internal/email -v
```

### 📚 What I Learned

✅ **Clean architecture in Go**: separation between `cmd` (entry) and `internal` (logic)  
✅ **Module management**: using `go.mod` to organize imports  
✅ **Error handling**: creating custom and descriptive errors  
✅ **String manipulation**: using the standard library `strings` package  
✅ **Unit testing**: ensuring logic works in various scenarios  
✅ **Go conventions**: naming, folder structure, and best practices

### 🎯 Next Steps

- [ ] Add special character validation
- [ ] Implement blocked domains list
- [ ] Create interactive CLI interface
- [ ] Add batch validation support

---

## � Author

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

## �📄 License

This project is for educational purposes.

**Made with ❤️ by Duarte**
