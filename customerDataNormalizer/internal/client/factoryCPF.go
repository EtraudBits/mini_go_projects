package client

import "unicode"

// isValidChecksum valida os dois dígitos verificadores do CPF.
func isValidChecksum(digits []rune) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(digits[i]-'0') * (10 - i)
	}

	first := (sum * 10) % 11
	if first == 10 {
		first = 0
	}

	if first != int(digits[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(digits[i]-'0') * (11 - i)
	}

	second := (sum * 10) % 11
	if second == 10 {
		second = 0
	}

	return second == int(digits[10]-'0')
}

// NewCPF cria um novo CPF válido.
// Valida formato, tamanho, dígitos repetidos e checksum.
func NewCPF(value string) (CPF, error) {
	var digits []rune
	for _, r := range value {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	if len(digits) != 11 {
		return "", ErrCPFinvalid
	}

	// Verifica se todos os dígitos são iguais
	allEqual := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return "", ErrCPFinvalid
	}

	// Valida os dígitos verificadores
	if !isValidChecksum(digits) {
		return "", ErrCPFinvalid
	}

	return CPF(string(digits)), nil
}


