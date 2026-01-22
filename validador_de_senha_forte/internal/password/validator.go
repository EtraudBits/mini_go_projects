package password

import "unicode"

// Validate verifica se a senha atende às regras de segurança.
func Validate(password string) error {
	if len(password) < 8 {
		return ErrTooShort
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return ErrNoUppercase
	}
	if !hasLower {
		return ErrNoLowercase
	}
	if !hasNumber {
		return ErrNoNumber
	}
	if !hasSpecial {
		return ErrNoSpecialChar
	}

	return nil
}
