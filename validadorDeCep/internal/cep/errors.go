package cep

import "errors"

// cada erro representa uma regra de negocio violada.
// errors.New cria um erro simples com mensagem imutavel.

var (

	ErrCEPEmpty = errors.New("cep não pode ser vazio")
	ErrOnlyNumbers = errors.New("cep deve conter apenas números")
	ErrInvalidLength = errors.New("cep deve ter exatamente 8 dígitos")
	ErrIvalidValue = errors.New("cep invalido")

)

// error é um tipo de interface nativa do GO, representa um erro.

/*
Regras de validação (lógica nua e crua)

Vamos validar um CEP brasileiro considerando:

1 - não pode ser vazio
2 - deve conter apenas números
3 - deve ter exatamente 8 dígitos
4 - não pode ser uma sequência inválida (ex: 00000000)

⚠️ Importante:
Vamos retornar todos os erros de uma vez, não parar no primeiro.
*/