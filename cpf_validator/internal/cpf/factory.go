package cpf

import "unicode"

// NewCPF cria um novo cpf valido / NewCPF creates a new valid cpf
// Se algo estiver errado, falha imediatamente / if something is wrong, it fails immediately
func NewCPF(value string) (CPF, error) {

	var digits []rune // apenas digitos / only digits
	for _, r := range value {
		if unicode.IsDigit(r) { // se for digito / if it's a digit
			digits = append(digits, r) // add ao slice / add to slice
		}
	}
	if len(digits) != 11 {
		return "", ErrInvalidLength
	}

	// Verifica se todos os digitos são iguais / checks if all digits are the same
	allEqual := true
	for i := 1; i <11; i++ {
		if digits [i] != digits[0] { // se algum for diferente / if any is different
			allEqual = false
			break
		}
	}
	if allEqual {
		return "", ErrRepeatedDigits
	}

	// Valida os digitos verificadores / validates the check digits
	if !isValidChecksum(digits) {
		return "", ErrInvalidChecksum
	}

	return CPF(string(digits)), nil
}