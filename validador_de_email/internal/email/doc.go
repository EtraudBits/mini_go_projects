/*
Package email fornece validação de endereços de email.

# Visão Geral

Este pacote implementa validação básica de endereços de email, verificando
o formato e a estrutura essencial de um email válido.

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - validador.go: Função principal de validação
  - errors.go: Definição de todos os erros do pacote
  - validador_test.go: Testes automatizados

# Validações Implementadas

1. Email não vazio:
  - Verifica se o email fornecido não é uma string vazia

2. Exatamente um @:
  - Valida se o email contém apenas um caractere '@'
  - Emails com zero ou múltiplos '@' são inválidos

3. Parte local (antes do @):
  - Verifica se existe texto antes do '@'
  - Exemplo: em "usuario@dominio.com", "usuario" é a parte local

4. Domínio (depois do @):
  - Verifica se existe texto após o '@'
  - Exemplo: em "usuario@dominio.com", "dominio.com" é o domínio

5. Ponto no domínio:
  - Valida se o domínio contém pelo menos um ponto '.'
  - Exemplo: "dominio.com" tem um ponto, "dominio" não tem

# Uso Básico

	import "your-module/validador_de_email/internal/email"

	err := email.Validate("usuario@exemplo.com")
	if err != nil {
		log.Fatal("Email inválido:", err)
	}
	fmt.Println("Email válido!")

# Exemplos de Validação

Emails Válidos:
  - usuario@exemplo.com
  - nome.sobrenome@empresa.com.br
  - contato123@dominio.org
  - user+tag@mail.example.com

Emails Inválidos:
  - "" (vazio) → ErrEmptyEmail
  - "@exemplo.com" (sem parte local) → ErrInvalidLocalPart
  - "usuario@" (sem domínio) → ErrIvalidDomain
  - "usuario@@exemplo.com" (dois @) → ErrIvalidAtSymbol
  - "usuario@exemplo" (sem ponto no domínio) → ErrIvalidDotDomain
  - "usuario.exemplo.com" (sem @) → ErrIvalidAtSymbol

# Tratamento de Erros

O pacote define erros específicos para cada tipo de falha:
  - ErrEmptyEmail: Email vazio
  - ErrIvalidAtSymbol: Email deve conter exatamente um @
  - ErrInvalidLocalPart: Email deve conter texto antes do @
  - ErrIvalidDomain: Email deve conter domínio válido
  - ErrIvalidDotDomain: Domínio deve conter pelo menos um ponto

# Limitações

Esta implementação é uma validação básica e não cobre todos os casos
especificados na RFC 5322 (formato completo de email). Para validações
mais robustas em produção, considere usar bibliotecas especializadas
ou expressões regulares mais complexas.

Casos não cobertos:
  - Validação de caracteres especiais na parte local
  - Validação de TLD (Top Level Domain) válido
  - Verificação de existência do domínio (DNS)
  - Verificação de existência da caixa de email

# Exemplo Completo

	package main

	import (
		"fmt"
		"log"
		"your-module/validador_de_email/internal/email"
	)

	func main() {
		// Lista de emails para validar
		emails := []string{
			"usuario@exemplo.com",
			"@exemplo.com",
			"usuario@",
			"usuario.exemplo.com",
			"",
		}

		for _, emailAddr := range emails {
			err := email.Validate(emailAddr)
			if err != nil {
				fmt.Printf("❌ Email '%s' inválido: %v\n", emailAddr, err)
			} else {
				fmt.Printf("✅ Email '%s' válido\n", emailAddr)
			}
		}
	}
*/
package email
