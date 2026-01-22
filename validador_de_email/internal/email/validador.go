package email

import "strings"

// Validate valida o formato de um endereço de email.
func Validate(email string) error {
	if email == "" {
		return ErrEmptyEmail
	}

	if strings.Count(email, "@") != 1 {
		return ErrIvalidAtSymbol
	}

	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]

	if local == "" {
		return ErrInvalidLocalPart
	}

	if domain == "" {
		return ErrIvalidDomain
	}

	if !strings.Contains(domain, ".") {
		return ErrIvalidDotDomain
	}

	return nil
}

