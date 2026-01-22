/*
Package client fornece funcionalidades para normalização de dados de clientes.

# Visão Geral

Este pacote implementa um sistema completo de normalização e validação de dados de clientes,
garantindo que todos os dados estejam em um formato consistente e válido antes de serem
processados ou armazenados.

# Arquitetura

O pacote segue os princípios:
  - Type Safety: Usa tipos nomeados (Name, CPF, Phone, Email) para garantir segurança de tipos
  - Factory Pattern: Usa factories para criar tipos validados
  - Fail-Fast: Retorna todos os erros de validação de uma vez
  - Separação de Responsabilidades: Cada arquivo tem uma responsabilidade específica

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - model.go: Definição da estrutura Client e tipos de domínio
  - types.go: Tipos nomeados para garantir type safety
  - normalizer.go: Função principal de normalização
  - factoryCPF.go: Factory para criação e validação de CPF
  - factoryPhone.go: Factory para criação e validação de telefone
  - factoryEmail.go: Factory para criação e validação de email
  - errors.go: Definição de todos os erros do pacote
  - normalizer_test.go: Testes automatizados

# Uso Básico

	client, errs := client.Normalize("joão silva", "123.456.789-09", "(11) 98765-4321", "joao@example.com")
	if len(errs) > 0 {
		// Trata erros de validação
		for _, err := range errs {
			log.Println(err)
		}
		return
	}
	// Usa o cliente normalizado
	fmt.Printf("Cliente: %+v\n", client)

# Validações Implementadas

CPF:
  - Remove formatação (pontos, hífens, espaços)
  - Valida se tem exatamente 11 dígitos
  - Verifica dígitos repetidos (ex: 111.111.111-11)
  - Valida dígitos verificadores usando o algoritmo oficial

Telefone:
  - Remove formatação (parênteses, hífens, espaços)
  - Valida se tem 10 ou 11 dígitos
  - Verifica se o DDD é válido (códigos existentes no Brasil)
  - Para celulares (11 dígitos), valida o nono dígito (deve ser 9)

Email:
  - Remove espaços em branco
  - Valida se não está vazio
  - Verifica se tem exatamente um @
  - Valida se tem texto antes do @
  - Valida se tem domínio após o @
  - Verifica se o domínio contém pelo menos um ponto

Nome:
  - Remove espaços em branco extras
  - Valida se não está vazio
  - Normaliza capitalização (primeira letra de cada palavra maiúscula)

# Tratamento de Erros

O pacote nunca retorna um cliente parcialmente válido sem avisar. Se houver qualquer erro
de validação, a função Normalize retorna um Client vazio e uma lista com todos os erros
encontrados, permitindo que o usuário veja todos os problemas de uma vez.

# Type Safety

Os tipos nomeados (Name, CPF, Phone, Email) garantem que:
  - Não é possível criar um CPF inválido acidentalmente
  - O compilador impede atribuições incorretas
  - O código é auto-documentado
  - A IDE fornece autocompletar e verificação de tipos

# Exemplo Completo

	package main

	import (
		"fmt"
		"log"
		"your-module/customerDataNormalizer/internal/client"
	)

	func main() {
		// Dados crus do cliente
		rawName := "  joão SILVA  "
		rawCPF := "123.456.789-09"
		rawPhone := "(11) 98765-4321"
		rawEmail := "  joao@example.com  "

		// Normaliza todos os dados
		client, errs := client.Normalize(rawName, rawCPF, rawPhone, rawEmail)

		// Verifica se houve erros
		if len(errs) > 0 {
			log.Println("Erros de validação:")
			for _, err := range errs {
				log.Printf("  - %v\n", err)
			}
			return
		}

		// Cliente válido e normalizado
		fmt.Printf("Cliente normalizado:\n")
		fmt.Printf("  Nome: %s\n", client.Name)
		fmt.Printf("  CPF: %s\n", client.CPF)
		fmt.Printf("  Telefone: %s\n", client.Phone)
		fmt.Printf("  Email: %s\n", client.Email)
	}
*/
package client
