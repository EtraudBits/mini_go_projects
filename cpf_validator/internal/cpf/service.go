package cpf

// isValidChecksum valida os dois dígitos verificadores do CPF.
func isValidChecksum(digits []rune) bool {
	// Calcula o primeiro dígito verificador
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

	// Calcula o segundo dígito verificador
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