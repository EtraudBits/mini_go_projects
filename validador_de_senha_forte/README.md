# 🔐 Validador de Senha Forte | Strong Password Validator

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Golang Logo" width="200"/>
  
  ### Sistema de validação de senhas seguras com regras de segurança
  ### Secure password validation system with security rules
</div>

---

## 📖 Sobre o Projeto | About the Project

**🇧🇷 Português:**
Mini projeto desenvolvido em Go para validar senhas fortes, aplicando 5 regras de segurança essenciais. Este projeto foi criado para aprender conceitos fundamentais de programação backend e boas práticas de desenvolvimento.

**🇺🇸 English:**
Mini project developed in Go to validate strong passwords, applying 5 essential security rules. This project was created to learn fundamental backend programming concepts and development best practices.

---

## 🎯 Regras de Validação | Validation Rules

✅ **Mínimo de 8 caracteres** | Minimum 8 characters  
✅ **Pelo menos 1 letra maiúscula** | At least 1 uppercase letter  
✅ **Pelo menos 1 letra minúscula** | At least 1 lowercase letter  
✅ **Pelo menos 1 número** | At least 1 number  
✅ **Pelo menos 1 caractere especial** | At least 1 special character

---

## 🛠️ Tecnologias e Conceitos | Technologies and Concepts

### 🔹 Linguagem | Language

- **Go (Golang)** - Versão 1.x

### 🔹 Pacotes Utilizados | Packages Used

- `unicode` - Identificação de tipos de caracteres
- `errors` - Tratamento customizado de erros
- `testing` - Testes unitários nativos do Go

---

## 📚 O Que Foi Aprendido | What Was Learned

### 1️⃣ **Lógica Sequencial | Sequential Logic**

🇧🇷 Estruturação do código em etapas lógicas: verificação de tamanho → análise de caracteres → validação de regras → retorno do resultado.

🇺🇸 Code structuring in logical steps: size verification → character analysis → rule validation → result return.

### 2️⃣ **Decisões Condicionais | Conditional Decisions**

🇧🇷 Uso de `if`, `switch` e operadores lógicos para tomar decisões baseadas em condições.

🇺🇸 Using `if`, `switch`, and logical operators to make decisions based on conditions.

```go
// Switch sem expressão (avalia condições booleanas)
switch {
case unicode.IsUpper(char):
    hasUpper = true
case unicode.IsLower(char):
    hasLower = true
}
```

### 3️⃣ **Leitura de Regras Como Sistema | Reading Rules as a System**

🇧🇷 Transformar requisitos de negócio em código validável e testável.

🇺🇸 Transforming business requirements into validatable and testable code.

- Cada regra = uma verificação específica
- Regras independentes e reutilizáveis
- Fácil manutenção e extensão

### 4️⃣ **Tratamento de Erro Real | Real Error Handling**

🇧🇷 Criação de erros customizados e específicos para cada tipo de falha.

🇺🇸 Creating custom and specific errors for each type of failure.

```go
var (
    ErrTooShort      = errors.New("senha deve ter no mínimo 8 caracteres")
    ErrNoUppercase   = errors.New("senha deve conter pelo menos uma letra maiúscula")
    ErrNoLowercase   = errors.New("senha deve conter pelo menos uma letra minúscula")
    // ...
)
```

### 5️⃣ **Separação de Responsabilidades | Separation of Concerns**

🇧🇷 Organização do código em camadas com responsabilidades bem definidas.

🇺🇸 Code organization in layers with well-defined responsibilities.

```
internal/password/
├── errors.go      → Definição de erros
├── validator.go   → Lógica de validação
└── validator_test.go → Testes unitários
```

### 6️⃣ **Arquitetura Básica de Backend | Basic Backend Architecture**

🇧🇷 Estrutura modular seguindo padrões de organização de projetos Go.

🇺🇸 Modular structure following Go project organization patterns.

- `cmd/` → Ponto de entrada da aplicação
- `internal/` → Código privado do projeto
- `go.mod` → Gerenciamento de dependências

### 7️⃣ **Sintaxe Go Usada no Dia a Dia | Go Syntax Used Daily**

#### 🔸 Variáveis Booleanas | Boolean Variables

```go
var hasUpper, hasLower, hasNumber, hasSpecial bool
```

#### 🔸 Loop com Range

```go
for _, char := range password {
    // processa cada caractere
}
```

#### 🔸 Funções com Retorno de Erro | Functions with Error Return

```go
func Validate(password string) error {
    // ...
    return nil // ou return error
}
```

#### 🔸 Testes Unitários | Unit Tests

```go
func TestValidate(t *testing.T) {
    err := Validate("Senha84&")
    if err != nil {
        t.Errorf("esperava senha válida, mas recebeu erro: %v", err)
    }
}
```

---

## 🚀 Como Executar | How to Run

### Pré-requisitos | Prerequisites

- Go 1.x instalado | Go 1.x installed

### Executar o Programa | Run the Program

```bash
cd validador_de_senha_forte
go run cmd/main.go
```

### Executar os Testes | Run Tests

```bash
cd internal/password
go test -v
```

### Resultado Esperado dos Testes | Expected Test Results

```
✓ TestValidate
✓ TestValidate_TooShort
✓ TestValidate_NoUppercase
✓ TestValidate_NoLowercase
✓ TestValidate_NoNumber
✓ TestValidate_NoSpecialChar
PASS
```

---

## 📂 Estrutura do Projeto | Project Structure

```
validador_de_senha_forte/
│
├── cmd/
│   └── main.go                 # Ponto de entrada | Entry point
│
├── internal/
│   └── password/
│       ├── errors.go           # Definição de erros | Error definitions
│       ├── validator.go        # Lógica de validação | Validation logic
│       └── validator_test.go   # Testes unitários | Unit tests
│
├── go.mod                      # Módulo Go | Go module
└── README.md                   # Documentação | Documentation
```

---

## 💡 Exemplos de Uso | Usage Examples

### ✅ Senha Válida | Valid Password

```go
err := Validate("Senha84&")
// err = nil
```

### ❌ Senha Inválida | Invalid Password

```go
err := Validate("senha")
// err = ErrTooShort

err := Validate("senhaSemNumero!")
// err = ErrNoUppercase

err := Validate("Senha123")
// err = ErrNoSpecialChar
```

---

## 🎓 Conceitos de Go Aplicados | Go Concepts Applied

| Conceito     | Onde Foi Usado         | Where It Was Used     |
| ------------ | ---------------------- | --------------------- |
| 🔹 Pacotes   | `package password`     | Package organization  |
| 🔹 Imports   | `import "unicode"`     | External dependencies |
| 🔹 Funções   | `func Validate()`      | Main validation logic |
| 🔹 Variáveis | `var hasUpper bool`    | State tracking        |
| 🔹 Loops     | `for _, char := range` | Character iteration   |
| 🔹 Switch    | `switch { case ... }`  | Conditional branching |
| 🔹 Errors    | `errors.New()`         | Error handling        |
| 🔹 Testing   | `func Test...()`       | Unit testing          |

---

## 📈 Próximos Passos | Next Steps

- [ ] Adicionar validação de senhas comuns (dicionário) | Add common password validation (dictionary)
- [ ] Implementar gerador de senhas fortes | Implement strong password generator
- [ ] Criar API REST para o validador | Create REST API for the validator
- [ ] Adicionar suporte a diferentes níveis de segurança | Add support for different security levels

---

## 🤝 Contribuições | Contributions

🇧🇷 Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

🇺🇸 Contributions are welcome! Feel free to open issues or pull requests.

---

## 📄 Licença | License

Este projeto é de código aberto e está disponível para fins educacionais.

This project is open source and available for educational purposes.

---

Author | Autor

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

<div align="center">
  
  **Desenvolvido com** ❤️ **usando** <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go" width="40"/>
  
  **Developed with** ❤️ **using Go**

</div>
