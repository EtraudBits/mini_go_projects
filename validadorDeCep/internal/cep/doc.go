/*
Package cep fornece validação de CEPs (Código de Endereçamento Postal) brasileiros.

# Visão Geral

Este pacote implementa validação do formato de CEP brasileiro, verificando
se o código está no formato correto e contém apenas dígitos válidos.

# Estrutura de Arquivos

  - doc.go: Documentação do pacote
  - validador.go: Função principal de validação
  - errors.go: Definição de todos os erros do pacote
  - validador_test.go: Testes automatizados

# Formato do CEP Brasileiro

O CEP (Código de Endereçamento Postal) é composto por 8 dígitos:

Formato padrão: XXXXX-XXX

Estrutura:
  - 5 primeiros dígitos: identificam região, subregião, setor, subsetor e divisor
  - 3 últimos dígitos: identificam a localidade específica

Exemplo: 01310-100 (Avenida Paulista, São Paulo - SP)

# Validações Implementadas

1. CEP não vazio:
  - Verifica se o CEP fornecido não é uma string vazia

2. Remove formatação:
  - Remove hífens, espaços e outros caracteres não numéricos
  - Exemplo: "01310-100" → "01310100"

3. Exatamente 8 dígitos:
  - Valida se o CEP tem exatamente 8 dígitos após remover formatação
  - CEPs com mais ou menos dígitos são inválidos

4. Apenas números:
  - Verifica se todos os caracteres (após remover formatação) são dígitos
  - Letras ou caracteres especiais são inválidos

# Uso Básico

	import "your-module/validadorDeCep/internal/cep"

	err := cep.Validate("01310-100")
	if err != nil {
		log.Fatal("CEP inválido:", err)
	}
	fmt.Println("CEP válido!")

# Exemplos de Validação

CEPs Válidos:
  - 01310-100 (com hífen)
  - 01310100 (sem formatação)
  - 20040-020 (outro exemplo com hífen)
  - 30130-000 (CEP com zeros)

CEPs Inválidos:
  - "" (vazio) → ErrCEPEmpty
  - "0131010" (7 dígitos) → ErrCEPInvalidLength
  - "013101000" (9 dígitos) → ErrCEPInvalidLength
  - "0131A-100" (contém letra) → ErrCEPInvalidFormat
  - "abcde-fgh" (não numérico) → ErrCEPInvalidFormat

# Tratamento de Erros

O pacote define erros específicos para cada tipo de falha:
  - ErrCEPEmpty: CEP vazio
  - ErrCEPInvalidLength: CEP não tem 8 dígitos
  - ErrCEPInvalidFormat: CEP contém caracteres não numéricos

# Limitações

Esta implementação valida apenas o formato do CEP. Não verifica:
  - Se o CEP realmente existe nos Correios
  - Se o CEP corresponde a um endereço real
  - Se o CEP está ativo ou foi desativado

Para validações mais robustas, considere integrar com a API dos Correios
ou usar serviços como ViaCEP (https://viacep.com.br/) que fornecem:
  - Consulta de CEP
  - Endereço completo
  - Bairro, cidade, estado
  - Validação de existência

# Estrutura do CEP

Os 8 dígitos do CEP têm significados específicos:

Exemplo: 01310-100

  - 01: Região (Grande São Paulo)
  - 3: Sub-região (zona central de São Paulo)
  - 10: Setor (Bela Vista)
  - 100: Identificador da localidade (Av. Paulista)

Principais regiões:
  - 01000-000 a 05999-999: São Paulo - SP
  - 20000-000 a 28999-999: Rio de Janeiro - RJ
  - 30000-000 a 39999-999: Belo Horizonte - MG
  - 40000-000 a 48999-999: Salvador - BA
  - 50000-000 a 54999-999: Recife - PE
  - 60000-000 a 61999-999: Fortaleza - CE
  - 70000-000 a 72799-999: Brasília - DF
  - 80000-000 a 82999-999: Curitiba - PR
  - 90000-000 a 91999-999: Porto Alegre - RS

# Exemplo Completo

	package main

	import (
		"fmt"
		"your-module/validadorDeCep/internal/cep"
	)

	func main() {
		// Lista de CEPs para validar
		ceps := []string{
			"01310-100",    // Válido: Av. Paulista, SP
			"20040-020",    // Válido: Centro, RJ
			"0131010",      // Inválido: 7 dígitos
			"01310-10A",    // Inválido: contém letra
			"",             // Inválido: vazio
		}

		for _, cepNum := range ceps {
			err := cep.Validate(cepNum)
			if err != nil {
				fmt.Printf("❌ CEP '%s' inválido: %v\n", cepNum, err)
			} else {
				fmt.Printf("✅ CEP '%s' válido\n", cepNum)
			}
		}
	}

# Integração com ViaCEP

Para consultar endereços completos, você pode integrar com ViaCEP:

	import (
		"encoding/json"
		"fmt"
		"net/http"
	)

	type Address struct {
		CEP         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		UF          string `json:"uf"`
	}

	func ConsultarCEP(cep string) (*Address, error) {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var addr Address
		if err := json.NewDecoder(resp.Body).Decode(&addr); err != nil {
			return nil, err
		}
		return &addr, nil
	}
*/
package cep
