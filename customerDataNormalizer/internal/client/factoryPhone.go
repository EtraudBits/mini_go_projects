package client

import (
	"unicode"
)

// NewPhone cria um novo Phone válido e retorna o tipo Phone | NewPhone creates a new valid Phone and returns the Phone type
// Se algo estiver errado, falha imediatamente | if something is wrong, it fails immediately
// Garante type safety: só é possível criar um Phone se for válido | Ensures type safety: only possible to create a Phone if it's valid
func NewPhone(value string) (Phone, error) {
	
	// regra 1 - Não pode ser vazio / rule 1 - Cannot be empty
	if value == "" {
		return "", ErrPhoneInvalid // se vazio, retorna imediatamente - if empty, returns immediately
	}

	var digits []rune // cria um slice de runas para armazenar os digitos do telefone - creates a slice of runes to store the phone digits
	// rune -> representa caracteres unicode -> vamos guardar apenas números.

	for _, char := range value { // percorre a string, ignora mascaras, guarda só numeros - iterates over the string, ignores masks, stores only numbers
		if unicode.IsDigit(char) { // se o caractere for um digito - if the character is a digit
			digits = append(digits, char) // add o digito ao slice - add the digit to the slice.
		}
	}

	// regra 2 - Deve conter 11 digitos numericos / rule 2 - Must contain 11 numeric digits
	if len(digits) != 11 { // se o tamanho (len) do slice for diferente de 11 - if the length (len) of the slice is different from 11.
		return "", ErrPhoneInvalid // retorna Phone vazio e erro | returns empty Phone and error
	}

	// regra 3 - Deve conter DDD valido (2 digitos) / rule 3 - Must contain valid DDD (2 digits)
	ddd := string(digits[0:2]) // pega os dois primeiros digitos do slice - takes the first two digits from the slice

	validateDDDs := map[string]bool { // cria um mapa (chave/valor) com os DDDs validos do Brasil - creates a map (key/value) with the valid DDDs from Brazil
		// Região Norte
		"68": true, // AC - Acre
		"82": true, // AL - Alagoas
		"96": true, // AP - Amapá
		"92": true, "97": true, // AM - Amazonas
		"71": true, "73": true, "74": true, "75": true, "77": true, // BA - Bahia
		"85": true, "88": true, // CE - Ceará
		"61": true, // DF - Distrito Federal
		"27": true, "28": true, // ES - Espírito Santo
		"62": true, "64": true, // GO - Goiás
		"98": true, "99": true, // MA - Maranhão
		"65": true, "66": true, // MT - Mato Grosso
		"67": true, // MS - Mato Grosso do Sul
		"31": true, "32": true, "33": true, "34": true, "35": true, "37": true, "38": true, // MG - Minas Gerais
		"91": true, "93": true, "94": true, // PA - Pará
		"83": true, // PB - Paraíba
		"41": true, "42": true, "43": true, "44": true, "45": true, "46": true, // PR - Paraná
		"81": true, "87": true, // PE - Pernambuco
		"86": true, "89": true, // PI - Piauí
		"21": true, "22": true, "24": true, // RJ - Rio de Janeiro
		"84": true, // RN - Rio Grande do Norte
		"51": true, "53": true, "54": true, "55": true, // RS - Rio Grande do Sul
		"69": true, // RO - Rondônia
		"95": true, // RR - Roraima
		"47": true, "48": true, "49": true, // SC - Santa Catarina
		"11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, // SP - São Paulo
		"79": true, // SE - Sergipe
		"63": true, // TO - Tocantins
	}

	if !validateDDDs[ddd] { // se o DDD não estiver no mapa - if the DDD is not in the map
		return "", ErrPhoneInvalid // retorna Phone vazio e erro | returns empty Phone and error
	}

	// regra 4 - O terceiro digito deve ser 9 (telefones celulares) / rule 4 - The third digit must be 9 (cell phones)
	if digits[2] != '9' { // se o terceiro digito (índice 2) não for 9 - if the third digit (index 2) is not 9
		return "", ErrPhoneInvalid // retorna Phone vazio e erro | returns empty Phone and error
	}

	return Phone(string(digits)), nil // retorna telefone validado e normalizado como tipo Phone | returns validated and normalized phone as Phone type
}

