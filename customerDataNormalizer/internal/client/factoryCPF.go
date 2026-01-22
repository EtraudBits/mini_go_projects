package client

import "unicode"

// isValidChecksum valida os dois digitos verificadores do CPF / validates the two check digits of the CPF
// É uma função interna, usada apenas no factory (letra minúscula) / it's an internal function, used only in the factory (lowercase letter)
func isValidChecksum(digits []rune) bool {
	// Calcula o primeiro digito verificador / calculates the first check digit
	sum := 0
	for i := 0; i < 9; i++ { // primeiros 9 digitos / first 9 digits
		sum += int(digits[i]-'0') * (10 - i) // peso decrescente de 10 a 2 / decreasing weight from 10 to 2
	}

	first := (sum * 10) % 11 // pega o resto da divisão por 11 / gets the remainder of division by 11
	if first == 10 {
		first = 0
	}

	if first != int(digits[9]-'0') { // compara com o digito verificador / compares with the check digit
		return false // se não bater, inválido / if it doesn't match, invalid
	}

	// calcula o segundo digito verificador / calculates the second check digit
	sum = 0
	for i := 0; i < 10; i++ { // primeiros 10 digitos / first 10 digits
		sum += int(digits[i]-'0') * (11 - i) // peso decrescente de 11 a 2 / decreasing weight from 11 to 2
	}

	second := (sum * 10) % 11 // pega o resto da divisão por 11 / gets the remainder of division by 11
	if second == 10 {
		second = 0 // se for 10, vira 0 / if it's 10, becomes 0
	}

	return second == int(digits[10]-'0') // compara com o digito verificador / compares with the check digit
}

// NewCPF cria um novo CPF válido e retorna o tipo CPF | NewCPF creates a new valid CPF and returns the CPF type
// Se algo estiver errado, falha imediatamente | if something is wrong, it fails immediately
// Garante type safety: só é possível criar um CPF se for válido | Ensures type safety: only possible to create a CPF if it's valid
func NewCPF(value string) (CPF, error) {

	var digits []rune // apenas digitos / only digits
	for _, r := range value {
		if unicode.IsDigit(r) { // se for digito / if it's a digit
			digits = append(digits, r) // add ao slice / add to slice
		}
	}
	if len(digits) != 11 {
		return "", ErrCPFinvalid
	}

	// Verifica se todos os digitos são iguais / checks if all digits are the same
	allEqual := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] { // se algum for diferente / if any is different
			allEqual = false
			break
		}
	}
	if allEqual {
		return "", ErrCPFinvalid
	}

	// Valida os digitos verificadores / validates the check digits
	if !isValidChecksum(digits) {
		return "", ErrCPFinvalid
	}

	return CPF(string(digits)), nil
}

