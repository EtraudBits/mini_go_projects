/*
Package cpf fornece validação completa de CPF (Cadastro de Pessoas Físicas).

# Visão Geral

Este pacote implementa a validação de CPF seguindo as regras oficiais da Receita Federal
do Brasil, incluindo validação de formato, verificação de dígitos repetidos e cálculo
dos dígitos verificadores.

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - model.go: Definição do tipo CPF
  - factory.go: Factory para criação e validação de CPF
  - service.go: Funções auxiliares de validação (checksum, dígitos repetidos)
  - errors.go: Definição de todos os erros do pacote
  - service_test.go: Testes automatizados

# Algoritmo de Validação

O CPF é composto por 11 dígitos, sendo os dois últimos dígitos verificadores.
O algoritmo de validação segue os seguintes passos:

1. Primeiro Dígito Verificador:
  - Multiplica os 9 primeiros dígitos por pesos de 10 a 2
  - Soma todos os resultados
  - Multiplica a soma por 10 e calcula o resto da divisão por 11
  - Se o resto for 10, o dígito verificador é 0

2. Segundo Dígito Verificador:
  - Multiplica os 10 primeiros dígitos por pesos de 11 a 2
  - Soma todos os resultados
  - Multiplica a soma por 10 e calcula o resto da divisão por 11
  - Se o resto for 10, o dígito verificador é 0

# Exemplo de Cálculo

CPF: 123.456.789-09

Primeiro dígito verificador (9º posição):

	1×10 + 2×9 + 3×8 + 4×7 + 5×6 + 6×5 + 7×4 + 8×3 + 9×2 = 210
	(210 × 10) % 11 = 2100 % 11 = 0

Segundo dígito verificador (10º posição):

	1×11 + 2×10 + 3×9 + 4×8 + 5×7 + 6×6 + 7×5 + 8×4 + 9×3 + 0×2 = 255
	(255 × 10) % 11 = 2550 % 11 = 9

# Uso Básico

	import "your-module/cpf_validator/internal/cpf"

	// Validar um CPF
	cpfValue, err := cpf.NewCPF("123.456.789-09")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CPF válido:", cpfValue)

# Validações Implementadas

  - Remove toda formatação (pontos, hífens, espaços)
  - Valida se tem exatamente 11 dígitos
  - Verifica se todos os dígitos são iguais (CPFs inválidos conhecidos)
  - Valida os dois dígitos verificadores usando o algoritmo oficial

# CPFs Inválidos Conhecidos

Os seguintes CPFs são considerados inválidos mesmo tendo checksum correto:
  - 000.000.000-00
  - 111.111.111-11
  - 222.222.222-22
  - 333.333.333-33
  - 444.444.444-44
  - 555.555.555-55
  - 666.666.666-66
  - 777.777.777-77
  - 888.888.888-88
  - 999.999.999-99

# Type Safety

O tipo CPF é um tipo nomeado (string) que garante que:
  - Só pode ser criado através do factory NewCPF
  - Sempre contém um CPF validado
  - O compilador impede atribuições diretas de string para CPF
  - O código é mais seguro e auto-documentado

# Tratamento de Erros

O pacote define erros específicos para cada tipo de falha:
  - ErrCPFEmpty: CPF vazio
  - ErrCPFInvalidLength: CPF não tem 11 dígitos
  - ErrCPFAllSameDigits: Todos os dígitos são iguais
  - ErrCPFInvalidChecksum: Dígitos verificadores inválidos

# Exemplo Completo

	package main

	import (
		"fmt"
		"log"
		"your-module/cpf_validator/internal/cpf"
	)

	func main() {
		// Exemplos de CPFs para validar
		cpfs := []string{
			"123.456.789-09",
			"111.111.111-11",  // Inválido: dígitos repetidos
			"123.456.789-00",  // Inválido: checksum incorreto
			"12345678909",     // Válido: sem formatação
		}

		for _, cpfStr := range cpfs {
			cpfValue, err := cpf.NewCPF(cpfStr)
			if err != nil {
				fmt.Printf("❌ CPF %s inválido: %v\n", cpfStr, err)
			} else {
				fmt.Printf("✅ CPF %s válido: %s\n", cpfStr, cpfValue)
			}
		}
	}
*/
package cpf
