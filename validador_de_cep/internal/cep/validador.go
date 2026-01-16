package cep

import (
	"strings" // para manipular strings (repetir caracteres)
	"unicode" // para validar se é número
)

// validate recebe cep e retorna uma lista de erros encontrados
// se a lista estiver vazia, o cep é valido.

func Validate(value string) []error { // value é uma string e retorna uma lista de erros
	var errs []error

	// regra 1 - não pode ser vazio
	if value == "" {
		errs = append(errs, ErrCEPEmpty) // append adiciona o erro ao slice erros []error
		return errs                      // se for vazio, retorna o erro e sai da função sem sentido validar o resto.
	}	
	
	// regra 2 - deve conter apenas números
	for _, char := range value { // percorre cada caractere da string
		if !unicode.IsDigit(char) { // se o caractere não for um digito
			errs = append(errs, ErrOnlyNumbers)
			break // não precisa continuar verificando
		}
	}

	// regra 3 - deve ter exatamente 8 dígitos
	if len(value) != 8 { // len conta a quantidade de caracteres na string se não for 8
		errs = append(errs, ErrInvalidLength) // adiciona o erro ao slice
	}

	// regra 4 - não pode ser uma sequencia de dígitos repetidos (ex: 00000000, 11111111, etc)
	if len(value) == 8 { // verifica se o CEP tem 8 caracteres para aplicar a validação
		firstChar := string(value[0]) // pega o primeiro caractere e converte para string
		// cria uma string com o primeiro caractere repetido 8 vezes e compara com o CEP
		// se forem iguais, significa que todos os dígitos são iguais
		if value == strings.Repeat(firstChar, 8) { // compara o CEP com a sequência repetida
			errs = append(errs, ErrIvalidValue) // adiciona erro de valor inválido
		}
	}

	return errs // retorna a lista de erros encontrados
}
