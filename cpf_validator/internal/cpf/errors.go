package cpf

import "errors"

var (
	ErrInvalidLength   = errors.New("cpf deve ter 11 digitos")                            // cpf must have 11 digits
	ErrRepeatedDigits  = errors.New("cpf não pode ter todos os digitos iguais")           // cpf cannot have all identical digits
	ErrInvalidChecksum = errors.New("cpf inválido (digitos verificadores incorretos)")    // invalid cpf (incorrect check digits)
)