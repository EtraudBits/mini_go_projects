package cpf

import "errors"

var (
	Err.InvalidLength = errors.New("cpf deve ter 11 digitos") // cpf must have 11 digits
	Err.RepeatedDigits = errors.New("cpf não pode ter todos os digitos iguais") // cpf cannot have all identical digits
	Err.InvalidChecksum = errors.New("cpf inválido (digitos verificadores incorretos)") // invalid cpf (incorrect check digits
	
)