package phone

import (
	"strings"
	"testing"
)

// Table-driven test - forma genérica de testar múltiplos cenários / Table-driven test - generic way to test multiple scenarios
func TestValidadorPhoneTable(t *testing.T) {
	// Define uma tabela com casos de teste (input e erro esperado) - Defines a table with test cases (input and expected error)
	testCases := []struct {
		name          string // nome descritivo do teste - descriptive test name
		phone         string // entrada do telefone - phone input
		expectedError error  // erro esperado - expected error
	}{
		// Casos válidos / Valid cases
		{name: "Telefone válido com máscara", phone: "(84) 98765-4321", expectedError: nil},
		{name: "Telefone válido sem máscara", phone: "84987654321", expectedError: nil},
		{name: "Telefone válido SP", phone: "11987654321", expectedError: nil},
		
		// Regra 1: vazio / Rule 1: empty
		{name: "Telefone vazio", phone: "", expectedError: ErrEmptyPhone},
		
		// Regra 2: quantidade de dígitos / Rule 2: number of digits
		{name: "Menos de 11 dígitos", phone: "123456789", expectedError: ErrInvalidDigits},
		{name: "Mais de 11 dígitos", phone: "849876543219999", expectedError: ErrInvalidDigits},
		
		// Regra 3: DDD inválido / Rule 3: invalid DDD
		{name: "DDD 00 inválido", phone: "00987654321", expectedError: ErrInvalidDDD},
		{name: "DDD 99 inexistente", phone: "09987654321", expectedError: ErrInvalidDDD},
		
		// Regra 4: não começa com 9 / Rule 4: does not start with 9
		{name: "Começa com 8", phone: "84887654321", expectedError: ErrInvalidStart},
		{name: "Começa com 7", phone: "11787654321", expectedError: ErrInvalidStart},
		
		// Múltiplas regras violadas - retorna o primeiro erro da validação sequencial / Multiple rules violated - returns first error from sequential validation
		{name: "DDD inválido + poucos dígitos", phone: "(00) 1234124", expectedError: ErrInvalidDigits},
		{name: "DDD inválido + não começa com 9", phone: "00812345678", expectedError: ErrInvalidDDD},
		{name: "Poucos dígitos + não começa com 9", phone: "848765432", expectedError: ErrInvalidDigits},
	}

	// Percorre cada caso de teste / Iterates over each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // executa cada teste com nome descritivo - runs each test with descriptive name
			err := ValidadorPhone(tc.phone) // chama o validador - calls the validator
			
			// Para casos com múltiplos erros, verifica se o erro esperado está contido no erro retornado
			if tc.expectedError == nil {
				if err != nil {
					t.Errorf("Teste '%s': Esperava telefone válido, mas recebeu erro: %v", tc.name, err)
				}
			} else {
				if err == nil {
					t.Errorf("Teste '%s': Esperava erro %v, mas não recebeu erro", tc.name, tc.expectedError)
				} else if !strings.Contains(err.Error(), tc.expectedError.Error()) {
					t.Errorf("Teste '%s': Esperava erro contendo '%v', mas recebeu '%v'", tc.name, tc.expectedError, err)
				}
			}
		})
	}
}
