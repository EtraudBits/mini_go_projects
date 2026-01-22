package client

import (
	"strings"
)

// Normalize recebe dados crus de um cliente e retorna um cliente normalizado ou lista de erros.
// Nunca retorna client parcialmente válido sem avisar.
func Normalize(name, cpf, phone, email string) (Client, []error) {
	var errs []error

	normalizedName := normalizeName(name, &errs)
	normalizedCPF := normalizeCPF(cpf, &errs)
	normalizedPhone := normalizePhone(phone, &errs)
	normalizedEmail := normalizeEmail(email, &errs)

	if len(errs) > 0 {
		return Client{}, errs
	}

	client := Client{
		Name:  normalizedName,
		CPF:   normalizedCPF,
		Phone: normalizedPhone,
		Email: normalizedEmail,
	}

	return client, nil
}

// normalizeName normaliza o nome removendo espaços extras e capitalizando cada palavra.
func normalizeName(value string, errs *[]error) Name {
	name := strings.TrimSpace(value)

	if name == "" {
		*errs = append(*errs, ErrNameEmpty)
		return ""
	}

	words := strings.Fields(strings.ToLower(name))

	for i, w := range words {
		words[i] = strings.Title(w) // Nota: strings.Title está depreciado, considere usar cases.Title
	}
	return Name(strings.Join(words, " "))
}

// normalizeCPF valida e normaliza o CPF usando o factory.
func normalizeCPF(value string, errs *[]error) CPF {
	cpf, err := NewCPF(value)
	if err != nil {
		*errs = append(*errs, err)
		return ""
	}
	return cpf
}

// normalizePhone valida e normaliza o telefone usando o factory.
func normalizePhone(value string, errs *[]error) Phone {
	phone, err := NewPhone(value)
	if err != nil {
		*errs = append(*errs, err)
		return ""
	}
	return phone
}

// normalizeEmail valida e normaliza o email usando o factory.
func normalizeEmail(value string, errs *[]error) Email {
	email, err := NewEmail(value)
	if err != nil {
		*errs = append(*errs, err)
		return ""
	}
	return email
}
