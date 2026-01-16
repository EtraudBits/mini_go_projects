package cep

import (
	"fmt"     // importa fmt para formatar strings dinamicamente
	"strings" // importa strings para criar strings repetidas
	"testing" // importa o pacote de teste do Go para criar testes automatizados
)

func TestValidate(t *testing.T) {
	// tabela de testes (table-driven tests) - padrão comum em Go para organizar múltiplos casos de teste
	// isso permite testar vários cenários sem duplicar código
	tests := []struct { // cria um slice de structs anônimos para armazenar os casos de teste
		name        string // nome descritivo do teste - aparecerá na saída quando executar os testes
		cep         string // o valor do CEP que será testado - entrada da função Validate
		shouldError bool   // indica se este CEP deve gerar erro (true) ou ser válido (false)
	}{
		// casos válidos - CEPs que devem passar na validação sem erros
		{
			name:        "CEP válido", // teste com um CEP válido simples
			cep:         "12345678", // 8 dígitos numéricos, sem repetições
			shouldError: false, // não deve retornar erro
		},
		{
			name:        "CEP válido diferente", // outro exemplo de CEP válido
			cep:         "59360000", // CEP real cidade de Parelhas - RN
			shouldError: false, // não deve retornar erro
		},

		// regra 1 - não pode ser vazio
		{
			name:        "CEP vazio", // testa se a função detecta CEP vazio
			cep:         "", // string vazia
			shouldError: true, // deve retornar erro ErrCEPEmpty
		},

		// regra 2 - deve conter apenas números
		{
			name:        "CEP com letras", // testa se detecta letras no meio do CEP
			cep:         "12A45678", // letra 'A' no meio dos números
			shouldError: true, // deve retornar erro ErrOnlyNumbers
		},
		{
			name:        "CEP com letras no início", // testa letra no começo
			cep:         "A1234567", // letra 'A' na primeira posição
			shouldError: true, // deve retornar erro ErrOnlyNumbers
		},
		{
			name:        "CEP apenas letras", // testa CEP composto só de letras
			cep:         "ABCDEFGH", // 8 letras ao invés de números
			shouldError: true, // deve retornar erro ErrOnlyNumbers
		},
		{
			name:        "CEP com caracteres especiais", // testa se detecta caracteres especiais
			cep:         "12345-78", // hífen é comum em formatação de CEP, mas aqui deve ser inválido
			shouldError: true, // deve retornar erro ErrOnlyNumbers
		},
		{
			name:        "CEP com espaços", // testa se detecta espaços
			cep:         "123 5678", // espaço no meio do CEP
			shouldError: true, // deve retornar erro ErrOnlyNumbers
		},
	}

	// regra 3 - deve ter exatamente 8 dígitos
	// usa um for para gerar dinamicamente testes com tamanhos diferentes
	// testa de 1 até 15 dígitos, exceto 8 (que é o tamanho válido)
	for i := 1; i <= 15; i++ { // percorre tamanhos de 1 a 15 dígitos
		if i == 8 { // pula o tamanho 8 (que é válido)
			continue // vai para a próxima iteração
		}
		tests = append(tests, struct { // adiciona um novo caso de teste ao slice
			name        string
			cep         string
			shouldError bool
		}{
			name:        fmt.Sprintf("CEP com %d dígitos", i), // cria nome descritivo dinamicamente
			cep:         strings.Repeat("1", i), // gera string com 'i' dígitos '1' (ex: "111" para i=3)
			shouldError: true, // qualquer tamanho diferente de 8 deve retornar erro ErrInvalidLength
		})
	}

	// regra 4 - não pode ser sequência de dígitos repetidos
	// usa um for para gerar testes com todas as sequências repetidas (00000000 até 99999999)
	for i := 0; i <= 9; i++ { // percorre os dígitos de 0 a 9
		digit := fmt.Sprintf("%d", i) // converte o número para string (ex: 0 -> "0", 5 -> "5")
		tests = append(tests, struct { // adiciona um novo caso de teste ao slice
			name        string
			cep         string
			shouldError bool
		}{
			name:        fmt.Sprintf("CEP com todos %ss", digit), // cria nome descritivo (ex: "CEP com todos 0s")
			cep:         strings.Repeat(digit, 8), // repete o dígito 8 vezes (ex: "00000000", "55555555")
			shouldError: true, // sequências repetidas devem retornar erro ErrIvalidValue
		})
	}

	// executa todos os testes da tabela acima
	for _, tt := range tests { // percorre cada caso de teste (tt = table test)
		t.Run(tt.name, func(t *testing.T) { // cria um subteste com nome descritivo - permite executar testes individualmente
			errs := Validate(tt.cep) // chama a função Validate com o CEP do caso de teste atual

			if tt.shouldError { // se este caso de teste espera que haja erro
				if len(errs) == 0 { // mas a função não retornou nenhum erro
					t.Errorf("esperava erros para CEP '%s', mas não recebi nenhum", tt.cep) // falha o teste com mensagem explicativa
				}
			} else { // se este caso de teste espera que NÃO haja erro (CEP válido)
				if len(errs) != 0 { // mas a função retornou erro(s)
					t.Errorf("não esperava erros para CEP '%s', mas recebi: %v", tt.cep, errs) // falha o teste mostrando quais erros foram retornados
				}
			}
		})
	}
}
