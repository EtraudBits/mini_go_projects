package cep

import (
	"strings"
	"unicode"
)

// Validate recebe um CEP e retorna uma lista de erros encontrados.
// Se a lista estiver vazia, o CEP é válido.
func Validate(value string) []error {
	var errs []error

	if value == "" {
		errs = append(errs, ErrCEPEmpty)
		return errs
	}

	// Deve conter apenas números
	for _, char := range value {
		if !unicode.IsDigit(char) {
			errs = append(errs, ErrOnlyNumbers)
			break
		}
	}

	// Deve ter exatamente 8 dígitos
	if len(value) != 8 {
		errs = append(errs, ErrInvalidLength)
	}

	// Não pode ser uma sequência de dígitos repetidos (ex: 00000000)
	if len(value) == 8 {
		firstChar := string(value[0])
		if value == strings.Repeat(firstChar, 8) {
			errs = append(errs, ErrIvalidValue)
		}
	}

	return errs
}
