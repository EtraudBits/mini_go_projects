package phone

import (
	"errors"
	"unicode"
)

// ValidadorPhone valida números de telefone brasileiros.
// Coleta todos os erros de validação para retornar ao usuário de uma vez.
func ValidadorPhone(phone string) error {
	var errs []error

	if phone == "" {
		return ErrEmptyPhone
	}

	var digits []rune
	for _, char := range phone {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	if len(digits) != 11 {
		errs = append(errs, ErrInvalidDigits)
	}

	// Valida DDD se tiver pelo menos 2 dígitos
	if len(digits) >= 2 {
		ddd := string(digits[0:2])

		validateDDDs := map[string]bool{
			// Região Norte
			"68": true, "82": true, "96": true, "92": true, "97": true,
			// Região Nordeste
			"71": true, "73": true, "74": true, "75": true, "77": true,
			"85": true, "88": true, "98": true, "99": true, "83": true,
			"81": true, "87": true, "86": true, "89": true, "84": true, "79": true,
			// Região Centro-Oeste
			"61": true, "62": true, "64": true, "65": true, "66": true, "67": true,
			// Região Sudeste
			"27": true, "28": true, "31": true, "32": true, "33": true, "34": true,
			"35": true, "37": true, "38": true, "21": true, "22": true, "24": true,
			"11": true, "12": true, "13": true, "14": true, "15": true, "16": true,
			"17": true, "18": true, "19": true,
			// Região Sul
			"41": true, "42": true, "43": true, "44": true, "45": true, "46": true,
			"47": true, "48": true, "49": true, "51": true, "53": true, "54": true, "55": true,
			// Norte
			"69": true, "95": true, "63": true, "91": true, "93": true, "94": true,
		}

		if !validateDDDs[ddd] {
			errs = append(errs, ErrInvalidDDD)
		}
	}

	// O nono dígito deve ser 9 (celulares)
	if len(digits) >= 3 && digits[2] != '9' {
		errs = append(errs, ErrInvalidStart)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
