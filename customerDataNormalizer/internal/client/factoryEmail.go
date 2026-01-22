package client

import "strings"

// NewEmail cria um novo Email válido.
// Valida formato, presença de @, parte local e domínio.
func NewEmail(value string) (Email, error) {
	email := strings.TrimSpace(value)
	
	if email == "" {
		return "", ErrEmailInvalid
	}

	if strings.Count(email, "@") != 1 {
		return "", ErrEmailInvalid
	}

	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]

	if local == "" {
		return "", ErrEmailInvalid
	}

	if domain == "" {
		return "", ErrEmailInvalid
	}

	if !strings.Contains(domain, ".") {
		return "", ErrEmailInvalid
	}

	return Email(strings.ToLower(email)), nil
}




