package client

import (
	"testing"
)

// TestNormalize testa função principal de normalização | tests main normalization function
func TestNormalize(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		inputCPF      string
		inputPhone    string
		inputEmail    string
		expectedName  Name
		expectedCPF   CPF
		expectedPhone Phone
		expectedEmail Email
		expectedError error
		errorCount    int
	}{
		{
			name:          "sucesso - todos os campos válidos",
			inputName:     "  joão SILVA santos  ",
			inputCPF:      "123.456.789-09",
			inputPhone:    "(11) 98765-4321",
			inputEmail:    "JOAO@EXAMPLE.COM",
			expectedName:  Name("João Silva Santos"),
			expectedCPF:   CPF("12345678909"),
			expectedPhone: Phone("11987654321"),
			expectedEmail: Email("joao@example.com"),
			errorCount:    0,
		},
		{
			name:          "erro - nome vazio",
			inputName:     "   ",
			inputCPF:      "123.456.789-09",
			inputPhone:    "(11) 98765-4321",
			inputEmail:    "test@example.com",
			expectedError: ErrNameEmpty,
			errorCount:    1,
		},
		{
			name:          "erro - CPF inválido (dígitos repetidos)",
			inputName:     "João Silva",
			inputCPF:      "111.111.111-11",
			inputPhone:    "(11) 98765-4321",
			inputEmail:    "test@example.com",
			expectedError: ErrCPFinvalid,
			errorCount:    1,
		},
		{
			name:          "erro - telefone inválido (sem o 9)",
			inputName:     "João Silva",
			inputCPF:      "123.456.789-09",
			inputPhone:    "(11) 1234-5678",
			inputEmail:    "test@example.com",
			expectedError: ErrPhoneInvalid,
			errorCount:    1,
		},
		{
			name:          "erro - email inválido (sem @)",
			inputName:     "João Silva",
			inputCPF:      "123.456.789-09",
			inputPhone:    "(11) 98765-4321",
			inputEmail:    "invalidemail",
			expectedError: ErrEmailInvalid,
			errorCount:    1,
		},
		{
			name:       "erro - múltiplos erros",
			inputName:  "   ",
			inputCPF:   "111.111.111-11",
			inputPhone: "(11) 1234-5678",
			inputEmail: "invalidemail",
			errorCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, errs := Normalize(tt.inputName, tt.inputCPF, tt.inputPhone, tt.inputEmail)

			// Verifica quantidade de erros | check error count
			if len(errs) != tt.errorCount {
				t.Errorf("esperava %d erros, got %d: %v", tt.errorCount, len(errs), errs)
			}

			// Se espera sucesso, valida os campos | if expecting success, validate fields
			if tt.errorCount == 0 {
				if client.Name != tt.expectedName {
					t.Errorf("Name: esperava %q, got %q", tt.expectedName, client.Name)
				}
				if client.CPF != tt.expectedCPF {
					t.Errorf("CPF: esperava %q, got %q", tt.expectedCPF, client.CPF)
				}
				if client.Phone != tt.expectedPhone {
					t.Errorf("Phone: esperava %q, got %q", tt.expectedPhone, client.Phone)
				}
				if client.Email != tt.expectedEmail {
					t.Errorf("Email: esperava %q, got %q", tt.expectedEmail, client.Email)
				}
			}

			// Se espera erro específico, valida | if expecting specific error, validate
			if tt.expectedError != nil && len(errs) > 0 {
				if errs[0] != tt.expectedError {
					t.Errorf("esperava erro %v, got %v", tt.expectedError, errs[0])
				}
			}
		})
	}
}

// TestNormalizeName testa normalização de nome isoladamente | tests name normalization in isolation
func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Name
		hasError bool
	}{
		{"nome simples", "joão silva", Name("João Silva"), false},
		{"nome com espaços extras", "  maria   SANTOS  ", Name("Maria Santos"), false},
		{"nome vazio", "   ", Name(""), true},
		{"nome composto", "ana PAULA de OLIVEIRA", Name("Ana Paula De Oliveira"), false},
		{"string vazia", "", Name(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errs []error
			result := normalizeName(tt.input, &errs)

			if tt.hasError && len(errs) == 0 {
				t.Error("esperava erro mas não obteve nenhum")
			}

			if !tt.hasError && len(errs) > 0 {
				t.Errorf("não esperava erro mas obteve: %v", errs)
			}

			if result != tt.expected {
				t.Errorf("esperava %q, got %q", tt.expected, result)
			}
		})
	}
}

// ========== TESTES DO FACTORY CPF ==========

// TestNewCPF_Valid testa criação de CPFs válidos | tests valid CPF creation
func TestNewCPF_Valid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected CPF
	}{
		{"CPF formatado válido", "123.456.789-09", CPF("12345678909")},
		{"CPF sem formatação", "12345678909", CPF("12345678909")},
		{"CPF com espaços", "123 456 789 09", CPF("12345678909")},
		{"CPF com caracteres mistos", "123-456-789.09", CPF("12345678909")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewCPF(tt.input)

			if err != nil {
				t.Errorf("não esperava erro mas obteve: %v", err)
			}

			if result != tt.expected {
				t.Errorf("esperava %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestNewCPF_Invalid testa todos os casos de CPF inválido | tests all invalid CPF cases
func TestNewCPF_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"CPF vazio", ""},
		{"CPF com menos de 11 dígitos", "123456789"},
		{"CPF com mais de 11 dígitos", "123456789012"},
		{"CPF com todos dígitos iguais - 000", "000.000.000-00"},
		{"CPF com todos dígitos iguais - 111", "111.111.111-11"},
		{"CPF com todos dígitos iguais - 222", "222.222.222-22"},
		{"CPF com todos dígitos iguais - 999", "999.999.999-99"},
		{"CPF com checksum inválido", "123.456.789-10"},
		{"CPF com primeiro dígito verificador errado", "123.456.789-19"},
		{"CPF com segundo dígito verificador errado", "123.456.789-08"},
		{"CPF apenas com letras", "abcdefghijk"},
		{"CPF com apenas caracteres especiais", "...---...--"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewCPF(tt.input)

			if err == nil {
				t.Errorf("esperava erro mas não obteve nenhum. Got: %v", result)
			}

			if err != ErrCPFinvalid {
				t.Errorf("esperava ErrCPFinvalid, got %v", err)
			}
		})
	}
}

// ========== TESTES DO FACTORY EMAIL ==========

// TestNewEmail_Valid testa criação de emails válidos | tests valid email creation
func TestNewEmail_Valid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Email
	}{
		{"email simples", "user@example.com", Email("user@example.com")},
		{"email com maiúsculas", "USER@EXAMPLE.COM", Email("user@example.com")},
		{"email com espaços", "  user@example.com  ", Email("user@example.com")},
		{"email com números", "user123@example456.com", Email("user123@example456.com")},
		{"email com caracteres especiais", "user.name+tag@example.co.uk", Email("user.name+tag@example.co.uk")},
		{"email com subdomínio", "user@mail.example.com", Email("user@mail.example.com")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewEmail(tt.input)

			if err != nil {
				t.Errorf("não esperava erro mas obteve: %v", err)
			}

			if result != tt.expected {
				t.Errorf("esperava %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestNewEmail_Invalid testa todos os casos de email inválido | tests all invalid email cases
func TestNewEmail_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"email vazio", ""},
		{"email com apenas espaços", "   "},
		{"email sem @", "userexample.com"},
		{"email com múltiplos @", "user@@example.com"},
		{"email com @ duplicado", "user@exam@ple.com"},
		{"email sem parte local", "@example.com"},
		{"email sem domínio", "user@"},
		{"email sem ponto no domínio", "user@example"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewEmail(tt.input)

			if err == nil {
				t.Errorf("esperava erro mas não obteve nenhum. Got: %v", result)
			}

			if err != ErrEmailInvalid {
				t.Errorf("esperava ErrEmailInvalid, got %v", err)
			}
		})
	}
}

// ========== TESTES DO FACTORY PHONE ==========

// TestNewPhone_Valid testa criação de telefones válidos | tests valid phone creation
func TestNewPhone_Valid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Phone
	}{
		{"celular formatado SP", "(11) 98765-4321", Phone("11987654321")},
		{"celular sem formatação SP", "11987654321", Phone("11987654321")},
		{"celular com espaços", "11 9 8765 4321", Phone("11987654321")},
		{"celular RJ", "(21) 99876-5432", Phone("21998765432")},
		{"celular BA", "(71) 91234-5678", Phone("71912345678")},
		{"celular RS", "(51) 99999-8888", Phone("51999998888")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewPhone(tt.input)

			if err != nil {
				t.Errorf("não esperava erro mas obteve: %v", err)
			}

			if result != tt.expected {
				t.Errorf("esperava %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestNewPhone_Invalid testa todos os casos de telefone inválido | tests all invalid phone cases
func TestNewPhone_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"telefone vazio", ""},
		{"telefone com apenas espaços", "   "},
		{"telefone com menos de 11 dígitos", "1198765432"},
		{"telefone com mais de 11 dígitos", "119876543210"},
		{"telefone sem DDD", "987654321"},
		{"telefone com DDD inválido", "(00) 98765-4321"},
		{"telefone com DDD zero", "(00) 98765-4321"},
		{"telefone sem o 9", "(11) 88765-4321"},
		{"telefone fixo (10 dígitos)", "(11) 3456-7890"},
		{"terceiro dígito não é 9", "11887654321"},
		{"telefone apenas com letras", "abcdefghijk"},
		{"telefone com caracteres especiais apenas", "()..-"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewPhone(tt.input)

			if err == nil {
				t.Errorf("esperava erro mas não obteve nenhum. Got: %v", result)
			}

			if err != ErrPhoneInvalid {
				t.Errorf("esperava ErrPhoneInvalid, got %v", err)
			}
		})
	}
}
