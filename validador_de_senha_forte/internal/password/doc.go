/*
Package password fornece validação de senhas fortes.

# Visão Geral

Este pacote implementa um validador de senhas que verifica se uma senha
atende aos critérios de complexidade necessários para ser considerada
uma senha forte e segura.

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - validator.go: Função principal de validação
  - errors.go: Definição de todos os erros do pacote
  - validator_test.go: Testes automatizados

# Critérios de Senha Forte

Uma senha forte deve atender a TODOS os seguintes critérios:

1. Comprimento Mínimo:
  - Pelo menos 8 caracteres

2. Letra Maiúscula:
  - Pelo menos uma letra maiúscula (A-Z)

3. Letra Minúscula:
  - Pelo menos uma letra minúscula (a-z)

4. Número:
  - Pelo menos um dígito (0-9)

5. Caractere Especial:
  - Pelo menos um caractere especial (!@#$%^&*()_+-=[]{}|;:,.<>?)

# Uso Básico

	import "your-module/validador_de_senha_forte/internal/password"

	err := password.Validate("SenhaForte123!")
	if err != nil {
		log.Fatal("Senha fraca:", err)
	}
	fmt.Println("Senha forte!")

# Exemplos de Validação

Senhas Válidas:
  - SenhaForte123!
  - P@ssw0rd
  - Abcd1234#
  - MinhaSenha@2024

Senhas Inválidas:
  - "" (vazia) → ErrPasswordEmpty
  - "curta1!" (menos de 8 caracteres) → ErrPasswordTooShort
  - "senhafraca123!" (sem maiúscula) → ErrPasswordNoUppercase
  - "SENHAFRACA123!" (sem minúscula) → ErrPasswordNoLowercase
  - "SenhaFraca!" (sem número) → ErrPasswordNoDigit
  - "SenhaFraca123" (sem especial) → ErrPasswordNoSpecialChar

# Tratamento de Erros

O pacote define erros específicos para cada critério não atendido:
  - ErrPasswordEmpty: Senha vazia
  - ErrPasswordTooShort: Senha com menos de 8 caracteres
  - ErrPasswordNoUppercase: Senha sem letra maiúscula
  - ErrPasswordNoLowercase: Senha sem letra minúscula
  - ErrPasswordNoDigit: Senha sem número
  - ErrPasswordNoSpecialChar: Senha sem caractere especial

A função Validate retorna apenas o primeiro erro encontrado, mas você
pode facilmente modificá-la para retornar todos os erros de uma vez.

# Segurança

Esta implementação verifica apenas a complexidade da senha.
Para um sistema completo de segurança, considere também:

  - Verificar contra senhas comuns (dictionary attack)
  - Verificar contra dados pessoais do usuário
  - Implementar limite de tentativas de login
  - Usar hash seguro (bcrypt, argon2) para armazenar senhas
  - Implementar autenticação de dois fatores (2FA)
  - Exigir troca periódica de senha
  - Verificar histórico de senhas usadas

# Boas Práticas

  - Nunca armazene senhas em texto plano
  - Use sempre algoritmos de hash seguros (bcrypt, argon2, scrypt)
  - Implemente rate limiting para tentativas de login
  - Registre tentativas de login falhadas
  - Use HTTPS para transmissão de senhas
  - Eduque usuários sobre senhas seguras

# Exemplo Completo

	package main

	import (
		"fmt"
		"your-module/validador_de_senha_forte/internal/password"
	)

	func main() {
		// Lista de senhas para validar
		passwords := []string{
			"SenhaForte123!",
			"senha123",
			"SENHA123!",
			"SenhaForte!",
			"SenhaForte123",
			"Curta1!",
		}

		for _, pwd := range passwords {
			err := password.Validate(pwd)
			if err != nil {
				fmt.Printf("❌ Senha '%s' fraca: %v\n", pwd, err)
			} else {
				fmt.Printf("✅ Senha '%s' forte\n", pwd)
			}
		}
	}
*/
package password
