package password

import "unicode" // unicode permite identificar letras, núemeros e símbolos corretamente.

// validate verifica se a senha atende às regras de segurança.
func Validate(password string) error {
if len(password) < 8 {
	return ErrTooShort // Regra 1 - mínimo de 8 caracteres
}
var hasUpper, hasLower, hasNumber, hasSpecial bool 

for _, char := range password {
	switch {
	case unicode.IsUpper(char):
		hasUpper = true // Regra 2 - letra maiúscula
	case unicode.IsLower(char):
		hasLower = true // Regra 3 - letra minúscula
	case unicode.IsDigit(char):
		hasNumber = true // Regra 4 - número
	case unicode.IsPunct(char) || unicode.IsSymbol(char):
		hasSpecial = true // Regra 5 - caractere especial
	}
} // Verifica se todas as regras foram atendidas

// Retorna erros específicos se alguma regra não for atendida
	if !hasUpper {
		return ErrNoUppercase // Regra 2
	}
	if !hasLower {
		return ErrNoLowercase // Regra 3
	}
	if !hasNumber {
		return ErrNoNumber // Regra 4
	}
	if !hasSpecial {
		return ErrNoSpecialChar // Regra 5
	}
return nil // Senha válida
}