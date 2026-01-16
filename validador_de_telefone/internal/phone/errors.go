package phone

import "errors"

var (
	// ErrEmptyPhone indica que o telefone esta vazio ou nulo / ErrEmail indicates that the email is empty or null
	ErrEmptyPhone = errors.New("telefopne não pode ser vazio")
	// ErrInvalidDigits indica que o telefone não possui 11 digitos / ErrIvalidDigits indicates that the phone does not have 11 digits
	ErrInvalidDigits = errors.New("telefone deve conter 11 digitos - Ex: 84987654321 ou (84) 98765-4321")
	// ErrinvalidDDD indica que o DDD do telefone é invalido / ErrInvalidDDD indicates that the phone DDD is invalid
	ErrInvalidDDD = errors.New("telefone deve conter DDD valido - Ex: 84, 11, 83 etc")
	// ErrInvalidStart indica que o telefone não começa com 9 após o DDD / ErrInvalidStart indicates that the phone does not start with 9 ofter the DDD
	ErrInvalidStart = errors.New("telefone deve começar com 9 após o DDD - Ex: 9xxxxyyyyy")
)
/* Regras de negócio (contrato claro)

Vamos validar telefone brasileiro no formato simples:
Aceita números com ou sem máscara
"11987654321"
"(11) 98765-4321"

Regras:
1 - Não pode ser vazio
2 - Deve conter 11 dígitos numéricos
3 - Deve conter DDD válido (2 dígitos)	
4 - Deve começar com 9 após o DDD (padrão celular BR)

⚠️ Foco em previsibilidade, não em cobrir todos os casos do planeta.*/