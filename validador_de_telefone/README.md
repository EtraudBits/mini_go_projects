# 📱 Validador de Telefone Brasileiro | Brazilian Phone Validator

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Golang Logo" width="200"/>
  
  ### Sistema de validação de números de telefone celular brasileiros
  ### Brazilian cell phone number validation system
</div>

---

## 📖 Sobre o Projeto | About the Project

**🇧🇷 Português:**
Mini projeto desenvolvido em Go para validar telefones celulares brasileiros, aplicando 4 regras de validação essenciais. Durante o desenvolvimento, realizamos **refatoração importante** do código para **retornar todos os erros possíveis de uma vez**, proporcionando melhor experiência ao usuário ao identificar múltiplas falhas simultaneamente. Este projeto foi criado para aprender conceitos fundamentais de programação backend e tratamento avançado de erros.

**🇺🇸 English:**
Mini project developed in Go to validate Brazilian cell phone numbers, applying 4 essential validation rules. During development, we performed **important code refactoring** to **return all possible errors at once**, providing better user experience by identifying multiple failures simultaneously. This project was created to learn fundamental backend programming concepts and advanced error handling.

---

## 🎯 Regras de Validação | Validation Rules

✅ **Não pode ser vazio** | Cannot be empty  
✅ **Deve conter exatamente 11 dígitos** | Must contain exactly 11 digits  
✅ **Deve ter DDD válido (2 primeiros dígitos)** | Must have valid area code (first 2 digits)  
✅ **Deve começar com 9 após o DDD** | Must start with 9 after area code

### Formatos Aceitos | Accepted Formats

- `84987654321` (sem máscara | without mask)
- `(84) 98765-4321` (com máscara | with mask)

---

## 🛠️ Tecnologias e Conceitos | Technologies and Concepts

### 🔹 Linguagem | Language

- **Go (Golang)** - Versão 1.x

### 🔹 Pacotes Utilizados | Packages Used

- `unicode` - Identificação e validação de caracteres numéricos
- `errors` - Tratamento customizado de erros + `errors.Join()` para múltiplos erros
- `testing` - Testes unitários com Table-Driven Tests

---

## 📚 O Que Foi Aprendido | What Was Learned

### 1️⃣ **Validação com Múltiplos Erros | Validation with Multiple Errors**

🇧🇷 **Refatoração Importante:** No início, o código retornava apenas o primeiro erro encontrado. Durante o aprendizado, **refatoramos o código para coletar e retornar TODOS os erros possíveis**, usando `errors.Join()` para unir múltiplos erros em uma única mensagem.

🇺🇸 **Important Refactoring:** Initially, the code returned only the first error found. During learning, **we refactored the code to collect and return ALL possible errors**, using `errors.Join()` to combine multiple errors into a single message.

```go
var errs []error // slice para coletar todos os erros

if len(digits) != 11 {
    errs = append(errs, ErrInvalidDigits)
}

if !validateDDDs[ddd] {
    errs = append(errs, ErrInvalidDDD)
}

if len(errs) > 0 {
    return errors.Join(errs...) // retorna todos os erros juntos
}
```

**Benefício:** O usuário recebe feedback completo de uma vez, em vez de descobrir os erros um por um.

**Benefit:** User receives complete feedback at once, instead of discovering errors one by one.

### 2️⃣ **Manipulação de Runes e Unicode | Rune and Unicode Manipulation**

🇧🇷 Aprendizado sobre o tipo `rune` em Go (representação de caracteres Unicode) para extrair apenas dígitos, ignorando máscaras como parênteses e hífens.

🇺🇸 Learning about Go's `rune` type (Unicode character representation) to extract only digits, ignoring masks like parentheses and hyphens.

```go
var digits []rune
for _, char := range phone {
    if unicode.IsDigit(char) {
        digits = append(digits, char)
    }
}
```

### 3️⃣ **Uso de Maps para Validação | Using Maps for Validation**

🇧🇷 Utilização de `map[string]bool` para armazenar e validar todos os DDDs (códigos de área) válidos do Brasil de forma eficiente.

🇺🇸 Using `map[string]bool` to store and validate all valid Brazilian area codes (DDDs) efficiently.

```go
validateDDDs := map[string]bool{
    "11": true, "21": true, "84": true, // etc...
}

if !validateDDDs[ddd] {
    errs = append(errs, ErrInvalidDDD)
}
```

### 4️⃣ **Table-Driven Tests | Table-Driven Tests**

🇧🇷 Implementação de testes usando a técnica Table-Driven Tests, permitindo testar múltiplos cenários com código limpo e organizado.

🇺🇸 Implementation of tests using the Table-Driven Tests technique, allowing testing multiple scenarios with clean and organized code.

```go
testCases := []struct {
    name          string
    phone         string
    expectedError error
}{
    {name: "Telefone válido", phone: "84987654321", expectedError: nil},
    {name: "DDD inválido", phone: "00987654321", expectedError: ErrInvalidDDD},
    // ...
}
```

### 5️⃣ **Slicing e Indexação de Strings | String Slicing and Indexing**

🇧🇷 Extração de substrings para validar partes específicas do telefone (DDD, nono dígito).

🇺🇸 Extracting substrings to validate specific parts of the phone number (area code, ninth digit).

```go
ddd := string(digits[0:2])    // primeiros 2 dígitos
if digits[2] != '9' {         // terceiro dígito deve ser 9
    errs = append(errs, ErrInvalidStart)
}
```

### 6️⃣ **Separação de Responsabilidades | Separation of Concerns**

🇧🇷 Organização do código em camadas com responsabilidades bem definidas.

🇺🇸 Code organization in layers with well-defined responsibilities.

```
internal/phone/
├── errors.go         → Definição de erros customizados
├── validador.go      → Lógica de validação
└── validador_test.go → Testes unitários completos
```

### 7️⃣ **Tratamento Avançado de Erros | Advanced Error Handling**

🇧🇷 Criação de erros customizados e específicos + uso de `errors.Join()` para combinar múltiplos erros.

🇺🇸 Creating custom and specific errors + using `errors.Join()` to combine multiple errors.

```go
var (
    ErrEmptyPhone    = errors.New("telefone não pode ser vazio")
    ErrInvalidDigits = errors.New("telefone deve conter 11 dígitos...")
    ErrInvalidDDD    = errors.New("telefone deve conter DDD válido...")
    ErrInvalidStart  = errors.New("telefone deve começar com 9...")
)
```

---

## 🚀 Como Executar | How to Run

### Pré-requisitos | Prerequisites

- Go 1.x instalado | Go 1.x installed

### Executar o Programa | Run the Program

```bash
cd validador_de_telefone
go run cmd/main.go
```

### Executar os Testes | Run Tests

```bash
cd internal/phone
go test -v
```

### Resultado Esperado dos Testes | Expected Test Results

```
✓ Telefone válido com máscara
✓ Telefone válido sem máscara
✓ Telefone válido SP
✓ Telefone vazio
✓ Menos de 11 dígitos
✓ Mais de 11 dígitos
✓ DDD 00 inválido
✓ DDD 99 inexistente
✓ Começa com 8
✓ Começa com 7
✓ DDD inválido + poucos dígitos
✓ DDD inválido + não começa com 9
✓ Poucos dígitos + não começa com 9
PASS
```

---

## 📂 Estrutura do Projeto | Project Structure

```
validador_de_telefone/
│
├── cmd/
│   └── main.go                 # Ponto de entrada | Entry point
│
├── internal/
│   └── phone/
│       ├── errors.go           # Definição de erros | Error definitions
│       ├── validador.go        # Lógica de validação | Validation logic
│       └── validador_test.go   # Testes unitários | Unit tests
│
├── go.mod                      # Módulo Go | Go module
└── README.md                   # Documentação | Documentation
```

---

## 💡 Exemplos de Uso | Usage Examples

### ✅ Telefone Válido | Valid Phone

```go
err := ValidadorPhone("84987654321")
// err = nil

err := ValidadorPhone("(84) 98765-4321")
// err = nil
```

### ❌ Telefone Inválido (Um Erro) | Invalid Phone (Single Error)

```go
err := ValidadorPhone("")
// err = "telefone não pode ser vazio"

err := ValidadorPhone("11887654321")
// err = "telefone deve começar com 9 após o DDD - Ex: 9xxxxyyyyy"
```

### ❌ Telefone Inválido (Múltiplos Erros) | Invalid Phone (Multiple Errors)

```go
err := ValidadorPhone("00812345678")
// err = "telefone deve conter 11 dígitos - Ex: 84987654321 ou (84) 98765-4321
//        telefone deve conter DDD válido - Ex: 84, 11, 83 etc
//        telefone deve começar com 9 após o DDD - Ex: 9xxxxyyyyy"
```

🎯 **Diferencial:** O sistema retorna **todos os erros encontrados**, não apenas o primeiro!

🎯 **Advantage:** The system returns **all errors found**, not just the first one!

---

## 🎓 Conceitos de Go Aplicados | Go Concepts Applied

| Conceito               | Onde Foi Usado                   | Where It Was Used          |
| ---------------------- | -------------------------------- | -------------------------- |
| 🔹 Pacotes             | `package phone`                  | Package organization       |
| 🔹 Imports             | `import "unicode"`               | External dependencies      |
| 🔹 Funções             | `func ValidadorPhone()`          | Main validation logic      |
| 🔹 Runes               | `[]rune` + `range`               | Unicode character handling |
| 🔹 Slices              | `var errs []error`               | Dynamic error collection   |
| 🔹 Maps                | `map[string]bool`                | DDD validation             |
| 🔹 Errors              | `errors.New()` + `errors.Join()` | Multiple error handling    |
| 🔹 Testing             | `func Test...()`                 | Table-driven tests         |
| 🔹 Structs (em testes) | `struct { name, phone, ... }`    | Test case organization     |

---

## 🔄 Evolução do Código | Code Evolution

### Versão Inicial | Initial Version

❌ Retornava apenas o primeiro erro encontrado  
❌ Usuário descobria erros um por um

### Versão Refatorada (Atual) | Refactored Version (Current)

✅ Coleta todos os erros antes de retornar  
✅ Usa `errors.Join()` para combinar erros  
✅ Usuário recebe feedback completo de uma vez  
✅ Melhor experiência de uso

---

## 📈 Próximos Passos | Next Steps

- [ ] Adicionar validação de telefones fixos | Add landline phone validation
- [ ] Suportar formatos internacionais | Support international formats
- [ ] Criar API REST para o validador | Create REST API for the validator
- [ ] Adicionar formatação automática de números | Add automatic number formatting
- [ ] Validar números ativos via API de operadora | Validate active numbers via carrier API

---

## 🤝 Contribuições | Contributions

🇧🇷 Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

🇺🇸 Contributions are welcome! Feel free to open issues or pull requests.

---

## 📄 Licença | License

Este projeto é de código aberto e está disponível para fins educacionais.

This project is open source and available for educational purposes.

---

## 👨‍💻 Autor | Author

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte-backend-golang)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

<div align="center">
  
  **Desenvolvido com** ❤️ **usando** <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go" width="40"/>
  
  **Developed with** ❤️ **using Go**

</div>
