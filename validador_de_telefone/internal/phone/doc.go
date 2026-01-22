/*
Package phone fornece validação de números de telefone brasileiros.

# Visão Geral

Este pacote implementa validação completa de números de telefone brasileiros,
incluindo telefones fixos (10 dígitos) e celulares (11 dígitos), com verificação
de DDD válido e formato correto.

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - validador.go: Função principal de validação
  - errors.go: Definição de todos os erros do pacote
  - validador_test.go: Testes automatizados

# Formato de Telefones Brasileiros

Telefone Fixo (10 dígitos):
  - Formato: (XX) XXXX-XXXX
  - Exemplo: (11) 3333-4444
  - Estrutura: DDD + 8 dígitos

Celular (11 dígitos):
  - Formato: (XX) 9XXXX-XXXX
  - Exemplo: (11) 98765-4321
  - Estrutura: DDD + 9 + 8 dígitos
  - O nono dígito deve ser 9

# DDDs Válidos no Brasil

O pacote valida os seguintes DDDs (códigos de área):
  - Região Sudeste: 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 24, 27, 28
  - Região Sul: 41, 42, 43, 44, 45, 46, 47, 48, 49, 51, 53, 54, 55
  - Região Centro-Oeste: 61, 62, 63, 64, 65, 66, 67
  - Região Nordeste: 71, 73, 74, 75, 77, 79, 81, 82, 83, 84, 85, 86, 87, 88, 89
  - Região Norte: 91, 92, 93, 94, 95, 96, 97, 98, 99

# Validações Implementadas

1. Remove formatação:
  - Remove parênteses, hífens, espaços e outros caracteres não numéricos
  - Exemplo: "(11) 98765-4321" → "11987654321"

2. Verifica tamanho:
  - Aceita 10 dígitos (fixo) ou 11 dígitos (celular)
  - Outros tamanhos são inválidos

3. Valida DDD:
  - Verifica se os dois primeiros dígitos formam um DDD válido
  - Rejeita DDDs que não existem no Brasil

4. Valida nono dígito (apenas celulares):
  - Para telefones com 11 dígitos, o terceiro dígito deve ser 9
  - Exemplo: 11 9 8765-4321 (o 9 é obrigatório)

# Uso Básico

	import "your-module/validador_de_telefone/internal/phone"

	err := phone.Validate("(11) 98765-4321")
	if err != nil {
		log.Fatal("Telefone inválido:", err)
	}
	fmt.Println("Telefone válido!")

# Exemplos de Validação

Telefones Válidos:
  - (11) 98765-4321 (celular SP)
  - (21) 3333-4444 (fixo RJ)
  - 11987654321 (sem formatação)
  - (85) 91234-5678 (celular CE)

Telefones Inválidos:
  - "" (vazio) → ErrPhoneEmpty
  - "(11) 8765-4321" (9 dígitos) → ErrPhoneInvalidLength
  - "(99) 98765-4321" (DDD inválido) → ErrPhoneInvalidDDD
  - "(11) 88765-4321" (nono dígito não é 9) → ErrPhoneInvalidNinthDigit
  - "123" (muito curto) → ErrPhoneInvalidLength

# Tratamento de Erros

O pacote define erros específicos para cada tipo de falha:
  - ErrPhoneEmpty: Telefone vazio
  - ErrPhoneInvalidLength: Telefone não tem 10 ou 11 dígitos
  - ErrPhoneInvalidDDD: DDD não é válido no Brasil
  - ErrPhoneInvalidNinthDigit: Para celular, terceiro dígito deve ser 9

# Sobre o Nono Dígito

Em 2012, o Brasil começou a adicionar o nono dígito (9) aos números de celular.
Este dígito foi implementado gradualmente em diferentes regiões até 2016.
Atualmente, TODOS os celulares brasileiros têm 11 dígitos começando com 9.

Exemplos:
  - Celular antigo: (11) 8765-4321 (inválido hoje)
  - Celular atual: (11) 98765-4321 (válido)

# Exemplo Completo

	package main

	import (
		"fmt"
		"your-module/validador_de_telefone/internal/phone"
	)

	func main() {
		// Lista de telefones para validar
		phones := []string{
			"(11) 98765-4321",  // Válido: celular SP
			"(21) 3333-4444",   // Válido: fixo RJ
			"(99) 98765-4321",  // Inválido: DDD não existe
			"(11) 88765-4321",  // Inválido: nono dígito errado
			"11 987654321",     // Válido: sem formatação
		}

		for _, phoneNum := range phones {
			err := phone.Validate(phoneNum)
			if err != nil {
				fmt.Printf("❌ Telefone '%s' inválido: %v\n", phoneNum, err)
			} else {
				fmt.Printf("✅ Telefone '%s' válido\n", phoneNum)
			}
		}
	}
*/
package phone
