# 🇧🇷 Validador de CEP em Go | 🇺🇸 ZIP Code Validator in Go

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

</div>

---

## 🇧🇷 Português

### 📖 Sobre o Projeto

Este projeto é um mini-projeto de estudo desenvolvido em **Golang**, focado em **validação de dados**, **organização de código limpo**, **testes automatizados abrangentes** e **refatoração inteligente**.

O objetivo principal não é apenas validar um CEP, mas entender **como pensar validações de forma genérica**, **como criar testes eficientes** e **por que escolher soluções mais elegantes ao invés de complexas**.

### 📌 O que este projeto faz

- ✅ Valida CEPs brasileiros (8 dígitos numéricos)
- 🔍 Verifica se o CEP não está vazio
- 🔢 Garante que contém apenas números
- 📏 Valida o tamanho exato de 8 dígitos
- 🚫 Rejeita sequências repetidas (00000000, 11111111, etc.)
- ❌ Retorna erros específicos e claros para cada tipo de falha
- 🧪 Possui testes automatizados gerados dinamicamente
- 🏗️ Mantém a lógica de negócio separada do ponto de entrada (`main`)

---

## 🎯 O que foi aprendido (e POR QUE)

### 1. **Validação em Camadas**

#### Por que validar em múltiplas regras?

Ao invés de fazer uma validação "tudo ou nada", separamos em **4 regras distintas**:

1. **CEP vazio** → Validação básica de entrada
2. **Apenas números** → Garantia de tipo de dado correto
3. **Tamanho exato** → Conformidade com padrão brasileiro
4. **Sequências repetidas** → Validação de valor lógico

**Por que isso importa:**

- ✅ **Debugging facilitado** → Sabemos exatamente qual regra falhou
- ✅ **Mensagens de erro precisas** → Usuário entende o que corrigir
- ✅ **Manutenção simples** → Cada regra é independente
- ✅ **Testabilidade** → Podemos testar cada regra isoladamente

```go
// Cada erro específico ajuda o desenvolvedor
ErrCEPEmpty        // "CEP está vazio"
ErrOnlyNumbers     // "CEP deve conter apenas números"
ErrInvalidLength   // "CEP deve ter exatamente 8 dígitos"
ErrIvalidValue     // "CEP não pode ser uma sequência repetida"
```

---

### 2. **Elegância vs Complexidade**

#### Evolução da Regra 4 (Sequências Repetidas)

**❌ Primeira versão (manual e específica):**

```go
if value == "00000000" {
    errs = append(errs, ErrIvalidValue)
}
```

**Problema:** Valida apenas `00000000`, ignora `11111111`, `22222222`, etc.

**⚠️ Segunda versão (genérica mas verbosa):**

```go
allSame := true
firstChar := value[0]
for i := 1; i < len(value); i++ {
    if value[i] != firstChar {
        allSame = false
        break
    }
}
if allSame {
    errs = append(errs, ErrIvalidValue)
}
```

**Problema:** Funciona, mas é muito código para algo simples.

**✅ Versão final (elegante e idiomática):**

```go
firstChar := string(value[0])
if value == strings.Repeat(firstChar, 8) {
    errs = append(errs, ErrIvalidValue)
}
```

**Por que essa versão é melhor:**

- 🎯 **Clareza** → Qualquer desenvolvedor entende imediatamente
- 🚀 **Idiomática** → Usa recursos nativos do Go (`strings.Repeat`)
- 🔧 **Manutenção** → Menos código = menos bugs
- 📖 **Legibilidade** → Se lê quase como linguagem natural

---

### 3. **Table-Driven Tests (Testes Orientados a Tabela)**

#### Por que este padrão é poderoso?

**❌ Forma tradicional (repetitiva):**

```go
func TestValidCEP(t *testing.T) { /* ... */ }
func TestEmptyCEP(t *testing.T) { /* ... */ }
func TestCEPWithLetters(t *testing.T) { /* ... */ }
// ... 30 funções diferentes
```

**✅ Forma com Table-Driven Tests:**

```go
tests := []struct {
    name        string
    cep         string
    shouldError bool
}{
    {"CEP válido", "12345678", false},
    {"CEP vazio", "", true},
    {"CEP com letras", "12A45678", true},
    // ... todos os casos em uma estrutura
}
```

**Por que isso importa:**

- 🔄 **Reutilização** → Mesma lógica de teste para todos os casos
- 📊 **Organização** → Casos de teste são dados, não código
- ⚡ **Eficiência** → Fácil adicionar novos casos
- 🎯 **Foco** → Separação entre dados de teste e lógica de verificação

---

### 4. **Geração Dinâmica de Testes com Loop**

#### A sacada mais importante do projeto

**Por que gerar testes dinamicamente?**

Ao invés de escrever manualmente:

```go
{"CEP com 1 dígito", "1", true},
{"CEP com 2 dígitos", "11", true},
{"CEP com 3 dígitos", "111", true},
// ... 14 casos manuais
```

Usamos um **loop para gerar automaticamente**:

```go
for i := 1; i <= 15; i++ {
    if i == 8 { continue } // pula o tamanho válido
    tests = append(tests, struct{...}{
        name: fmt.Sprintf("CEP com %d dígitos", i),
        cep: strings.Repeat("1", i),
        shouldError: true,
    })
}
```

**Resultado:** De 4 testes manuais para **14 testes gerados automaticamente**.

**Por que isso é revolucionário:**

- ✨ **DRY (Don't Repeat Yourself)** → Zero repetição de código
- 🎯 **Cobertura completa** → Testa todos os tamanhos de 1 a 15
- 🛡️ **Prova de futuro** → Mudou a regra? Muda só o loop
- 🧠 **Pensamento algorítmico** → Dados gerados por lógica, não digitação

---

### 5. **Mesma lógica para testes E produção**

#### Consistência é fundamental

Percebemos que a mesma lógica usada nos testes (`strings.Repeat`) poderia ser aplicada no código de produção:

**Nos testes:**

```go
cep: strings.Repeat("1", i),  // gera "111...1"
```

**No validador:**

```go
if value == strings.Repeat(firstChar, 8) { // verifica se é repetido
```

**Por que isso importa:**

- 🔄 **Consistência** → Mesma abordagem em teste e produção
- 🤝 **Confiança** → Se funciona nos testes, funciona no código
- 📚 **Aprendizado** → Técnicas de teste melhoram código de produção

---

## 🛠️ Tecnologias e Recursos Utilizados

### Pacotes Go

| Pacote    | Uso                                         | Por que usar                                        |
| --------- | ------------------------------------------- | --------------------------------------------------- |
| `unicode` | Validar se caracteres são dígitos numéricos | Forma idiomática em Go para validação de caracteres |
| `strings` | Repetir strings para comparação             | Solução elegante para detectar sequências repetidas |
| `fmt`     | Formatar strings dinamicamente              | Criar nomes de testes descritivos automaticamente   |
| `testing` | Framework de testes nativo do Go            | Testes integrados sem dependências externas         |

### Conceitos Aplicados

- ✅ **Clean Code** → Código auto-explicativo com nomes claros
- ✅ **Separation of Concerns** → Lógica separada da interface
- ✅ **Error Handling** → Erros customizados e específicos
- ✅ **Table-Driven Tests** → Padrão idiomático de Go
- ✅ **Refatoração** → Evolução de código verboso para elegante
- ✅ **Test Coverage** → Cobertura completa de casos válidos e inválidos

---

## 📂 Estrutura do Projeto

```
validadorDeCep/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   └── cep/
│       ├── validador.go     # Lógica de validação
│       ├── validador_test.go # Testes automatizados
│       └── errors.go        # Erros customizados
├── go.mod
└── README.md
```

### Por que essa estrutura?

- **`cmd/`** → Ponto de entrada, não contém lógica de negócio
- **`internal/cep/`** → Lógica isolada, reutilizável e testável
- **Separação de concerns** → Mudanças na UI não afetam validação
- **Testabilidade** → Lógica pode ser testada independentemente

---

## 🎮 Como Usar

### Clonar o Repositório

```bash
git clone https://github.com/seu-usuario/mini_go_projects.git
cd mini_go_projects/validadorDeCep
```

### Executar a Aplicação

```bash
go run cmd/main.go
```

### Executar os Testes

```bash
# Todos os testes
go test ./internal/cep/

# Com detalhes
go test -v ./internal/cep/

# Com cobertura
go test -cover ./internal/cep/
```

### Exemplo de Uso

```go
package main

import (
    "fmt"
    "validadorDeCep/internal/cep"
)

func main() {
    // Testa um CEP válido
    errs := cep.Validate("12345678")
    if len(errs) == 0 {
        fmt.Println("✅ CEP válido!")
    }

    // Testa um CEP inválido
    errs = cep.Validate("00000000")
    if len(errs) > 0 {
        fmt.Println("❌ Erros encontrados:")
        for _, err := range errs {
            fmt.Println(" -", err)
        }
    }
}
```

---

## 🧪 Cobertura de Testes

### Casos Testados

| Categoria                | Quantidade | Geração  |
| ------------------------ | ---------- | -------- |
| CEPs válidos             | 2          | Manual   |
| CEP vazio                | 1          | Manual   |
| Caracteres não numéricos | 5          | Manual   |
| Tamanhos inválidos       | 14         | **Loop** |
| Sequências repetidas     | 10         | **Loop** |
| **Total**                | **32**     | -        |

### Por que 32 testes?

- **2 válidos** → Garantem que CEPs corretos passam
- **1 vazio** → Valida entrada mínima
- **5 caracteres** → Letras, espaços, símbolos, etc.
- **14 tamanhos** → De 1 a 15 dígitos (exceto 8)
- **10 sequências** → 00000000 até 99999999

**Resultado:** Cobertura completa com código mínimo graças aos loops!

---

## 💡 Principais Lições

| Aspecto                 | Aprendizado                                                         |
| ----------------------- | ------------------------------------------------------------------- |
| **Pensamento Genérico** | Não resolver só o problema de hoje, pensar em todos os casos        |
| **Elegância**           | Código simples > código complexo. Use recursos nativos da linguagem |
| **Automação de Testes** | Loops reduzem código manual e aumentam cobertura                    |
| **Refatoração**         | Código evolui. Primeira versão pode não ser a melhor                |
| **Clean Code**          | Legibilidade é mais importante que esperteza                        |

---

## 🚀 Próximos Passos

- [ ] Integração com API de CEP (ViaCEP)
- [ ] Validação com formatação (00000-000)
- [ ] Benchmark de performance
- [ ] Adicionar cache de CEPs válidos
- [ ] REST API para validação de CEP

---

## 🇺🇸 English

### 📖 About the Project

This project is a mini study project developed in **Golang**, focused on **data validation**, **clean code organization**, **comprehensive automated testing**, and **smart refactoring**.

The main goal is not just to validate a ZIP code, but to understand **how to think about validations generically**, **how to create efficient tests**, and **why to choose elegant solutions over complex ones**.

### 📌 What this project does

- ✅ Validates Brazilian ZIP codes (8 numeric digits)
- 🔍 Checks if the ZIP code is not empty
- 🔢 Ensures it contains only numbers
- 📏 Validates the exact length of 8 digits
- 🚫 Rejects repeated sequences (00000000, 11111111, etc.)
- ❌ Returns specific and clear errors for each type of failure
- 🧪 Has dynamically generated automated tests
- 🏗️ Keeps business logic separated from the entry point (`main`)

---

## 🎯 What was learned (and WHY)

### 1. **Layered Validation**

#### Why validate in multiple rules?

Instead of doing an "all or nothing" validation, we separated it into **4 distinct rules**:

1. **Empty ZIP** → Basic input validation
2. **Only numbers** → Guarantee of correct data type
3. **Exact length** → Compliance with Brazilian standard
4. **Repeated sequences** → Logical value validation

**Why this matters:**

- ✅ **Easier debugging** → We know exactly which rule failed
- ✅ **Precise error messages** → User understands what to fix
- ✅ **Simple maintenance** → Each rule is independent
- ✅ **Testability** → We can test each rule in isolation

---

### 2. **Elegance vs Complexity**

#### Evolution of Rule 4 (Repeated Sequences)

**❌ First version (manual and specific):**

```go
if value == "00000000" {
    errs = append(errs, ErrIvalidValue)
}
```

**Problem:** Only validates `00000000`, ignores `11111111`, `22222222`, etc.

**⚠️ Second version (generic but verbose):**

```go
allSame := true
firstChar := value[0]
for i := 1; i < len(value); i++ {
    if value[i] != firstChar {
        allSame = false
        break
    }
}
if allSame {
    errs = append(errs, ErrIvalidValue)
}
```

**Problem:** Works, but too much code for something simple.

**✅ Final version (elegant and idiomatic):**

```go
firstChar := string(value[0])
if value == strings.Repeat(firstChar, 8) {
    errs = append(errs, ErrIvalidValue)
}
```

**Why this version is better:**

- 🎯 **Clarity** → Any developer understands immediately
- 🚀 **Idiomatic** → Uses Go's native features (`strings.Repeat`)
- 🔧 **Maintenance** → Less code = fewer bugs
- 📖 **Readability** → Reads almost like natural language

---

### 3. **Table-Driven Tests**

#### Why is this pattern powerful?

**❌ Traditional way (repetitive):**

```go
func TestValidCEP(t *testing.T) { /* ... */ }
func TestEmptyCEP(t *testing.T) { /* ... */ }
func TestCEPWithLetters(t *testing.T) { /* ... */ }
// ... 30 different functions
```

**✅ Table-Driven Tests way:**

```go
tests := []struct {
    name        string
    cep         string
    shouldError bool
}{
    {"Valid ZIP", "12345678", false},
    {"Empty ZIP", "", true},
    {"ZIP with letters", "12A45678", true},
    // ... all cases in one structure
}
```

**Why this matters:**

- 🔄 **Reusability** → Same test logic for all cases
- 📊 **Organization** → Test cases are data, not code
- ⚡ **Efficiency** → Easy to add new cases
- 🎯 **Focus** → Separation between test data and verification logic

---

### 4. **Dynamic Test Generation with Loops**

#### The most important insight of the project

**Why generate tests dynamically?**

Instead of writing manually:

```go
{"ZIP with 1 digit", "1", true},
{"ZIP with 2 digits", "11", true},
{"ZIP with 3 digits", "111", true},
// ... 14 manual cases
```

We use a **loop to generate automatically**:

```go
for i := 1; i <= 15; i++ {
    if i == 8 { continue } // skip valid length
    tests = append(tests, struct{...}{
        name: fmt.Sprintf("ZIP with %d digits", i),
        cep: strings.Repeat("1", i),
        shouldError: true,
    })
}
```

**Result:** From 4 manual tests to **14 automatically generated tests**.

**Why this is revolutionary:**

- ✨ **DRY (Don't Repeat Yourself)** → Zero code repetition
- 🎯 **Complete coverage** → Tests all sizes from 1 to 15
- 🛡️ **Future-proof** → Rule changed? Just change the loop
- 🧠 **Algorithmic thinking** → Data generated by logic, not typing

---

### 5. **Same logic for tests AND production**

#### Consistency is fundamental

We realized that the same logic used in tests (`strings.Repeat`) could be applied in production code:

**In tests:**

```go
cep: strings.Repeat("1", i),  // generates "111...1"
```

**In validator:**

```go
if value == strings.Repeat(firstChar, 8) { // checks if repeated
```

**Why this matters:**

- 🔄 **Consistency** → Same approach in test and production
- 🤝 **Confidence** → If it works in tests, it works in code
- 📚 **Learning** → Test techniques improve production code

---

## 💡 Key Lessons

| Aspect               | Learning                                                 |
| -------------------- | -------------------------------------------------------- |
| **Generic Thinking** | Don't solve just today's problem, think about all cases  |
| **Elegance**         | Simple code > complex code. Use native language features |
| **Test Automation**  | Loops reduce manual code and increase coverage           |
| **Refactoring**      | Code evolves. First version may not be the best          |
| **Clean Code**       | Readability is more important than cleverness            |

---

## 👤 Author | Autor

**Duarte Rodrigo Santos de Oliveira**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte84oliviera)
[![Email](https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:du84arte@gmail.com)

---

## 📎 Final Notes | Observações Finais

🇧🇷 Este é um projeto de estudo focado em **qualidade, refatoração e pensamento algorítmico**. O código evoluiu de uma solução simples para uma implementação elegante através de aprendizado contínuo. Sugestões e feedbacks são sempre bem-vindos!

🇺🇸 This is a study project focused on **quality, refactoring, and algorithmic thinking**. The code evolved from a simple solution to an elegant implementation through continuous learning. Suggestions and feedback are always welcome!

---

<div align="center">

**⭐ Se este projeto foi útil para você, considere dar uma estrela!**

**⭐ If this project was helpful, consider giving it a star!**

</div>
