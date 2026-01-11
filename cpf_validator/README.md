# Validador de CPF em Go

Este projeto é um mini-projeto de estudo desenvolvido em **Golang**, com foco em **lógica de programação**, **organização de código**, **tratamento de erros** e **estrutura próxima da realidade de projetos backend**.

O objetivo principal não é apenas validar um CPF, mas entender **como pensar um problema**, **como organizar a solução** e **por que cada parte do código existe**.

---

## 📌 O que este projeto faz

- Recebe um CPF em formato texto
- Remove caracteres inválidos
- Valida regras básicas (tamanho, repetição de números)
- Calcula e valida os dígitos verificadores
- Retorna erro quando o CPF é inválido
- Mantém a lógica separada do ponto de entrada (`main`)

---

## 🧠 Por que um validador de CPF?

Validação de CPF é algo muito comum em sistemas reais, como:

- cadastros de usuários
- sistemas bancários
- e-commerce
- APIs de autenticação
- sistemas governamentais

Ou seja, apesar de simples, esse tipo de lógica aparece com frequência no dia a dia de quem trabalha com backend.

---

## 🏗️ Estrutura do projeto

````text
.
├── go.mod
├── cmd/
│   └── main.go          # Ponto de entrada da aplicação
└── internal/
    └── cpf/
        ├── validator.go # Lógica de validação do CPF
        └── errors.go    # Erros específicos do domínio

---

Por que essa estrutura?

cmd/
Contém apenas o código que executa o programa.
O main.go chama funções externas, como acontece em projetos reais.

internal/
Contém a regra de negócio.
Nada aqui depende do main, o que facilita testes, manutenção e reutilização.

Separação de responsabilidades
Cada arquivo e função tem um propósito claro.

---

⚙️ Conceitos praticados neste projeto

Lógica de programação aplicada

Funções e retorno de valores

Tratamento explícito de erros em Go

Organização de pacotes

Nomes de funções e variáveis com significado

Leitura de código pensando em fluxo e invariantes

Pensar em falhas antes do sucesso

---

🚀 Como executar o projeto
go run ./cmd

---

📚 O que estou aprendendo com este projeto

Pensar antes de codar

Entender o problema antes da solução

Separar lógica de execução

Ler e escrever código com mais clareza

Me aproximar da forma como projetos reais são organizados

Este projeto faz parte da minha rotina de estudos como desenvolvedor backend iniciante.

---

👤 Autor

Duarte Rodrigo Santos de Oliveira

LinkedIn: https://www.linkedin.com/in/duarte84oliviera

Email: du84arte@gmail.com

---

📎 Observação final

Este é um projeto de estudo.
Sugestões, feedbacks e melhorias são sempre bem-vindos.

---

# CPF Validator in Go

This project is a **study-oriented mini project** developed in **Golang**, focused on **programming logic**, **code organization**, **error handling**, and **backend-oriented structure**.

The main goal is not only to validate a CPF number, but to understand **how to think about a problem**, **how to structure a solution**, and **why each part of the code exists**.

---

## 📌 What this project does

- Receives a CPF as a string
- Removes invalid characters
- Validates basic rules (length, repeated digits)
- Calculates and validates check digits
- Returns explicit errors for invalid cases
- Keeps business logic separated from the application entry point (`main`)

---

## 🧠 Why a CPF validator?

CPF validation is very common in real-world systems such as:
- user registration
- banking systems
- e-commerce platforms
- authentication APIs
- government systems

Even though it looks simple, this kind of logic appears frequently in backend development.

---

## 🏗️ Project structure

```text
.
├── go.mod
├── cmd/
│   └── main.go
└── internal/
    └── cpf/
        ├── validator.go
        └── errors.go

---

Structure rationale

cmd/
Contains only the application entry point.
The main.go file calls external logic, similar to real backend projects.

internal/
Contains business logic, isolated from execution details.
This improves maintainability and reuse.

---

⚙️ Concepts practiced

Programming logic

Functions and return values

Explicit error handling in Go

Package organization

Meaningful naming

Thinking about failure before success

Reading code as a system

---

🚀 How to run
go run ./cmd

---

📚 What I’m learning

Thinking before coding

Understanding the problem first

Separating execution from logic

Writing clearer and more maintainable code

Approaching backend projects realistically

---

👤 Author

Duarte Rodrigo Santos de Oliveira

LinkedIn: https://www.linkedin.com/in/duarte84oliviera

Email: du84arte@gmail.com

---

📎 Notes

This is a study project.
Feedback and suggestions are welcome.
````
