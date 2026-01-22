package client

import "strings" // "strings" é um pacote da biblioteca padrão do Go que fornece funções para manipulação de strings
// Precisamos dividir (Split), contar (Count), verificar conteúdo (Contains), entre outras operações com strings

// NewEmail cria um novo Email válido e retorna o tipo Email | NewEmail creates a new valid Email and returns the Email type
// Se algo estiver errado, falha imediatamente | if something is wrong, it fails immediately
// Garante type safety: só é possível criar um Email se for válido | Ensures type safety: only possible to create an Email if it's valid
func NewEmail(value string) (Email, error) {
	email := strings.TrimSpace(value) // remove espaços em branco no inicio e no fim | removes leading and trailing whitespace
	
	if email == "" { // 1 - Não estiver vazio | must not be empty
		return "", ErrEmailInvalid // retorna Email vazio e erro | returns empty Email and error
	}

	if strings.Count(email, "@") != 1 { // 2 - Contiver exatamente um @ | must contain exactly one @
		return "", ErrEmailInvalid // retorna Email vazio e erro | returns empty Email and error
	}

	parts := strings.Split(email, "@") // divide o email em duas partes: antes e depois do '@' | splits email into two parts: before and after '@'
	local := parts[0] // parte antes do '@' | part before '@'
	domain := parts[1] // parte depois do '@' | part after '@'

	if local == "" { // 3 - Tiver texto antes do @ | must have text before @
		return "", ErrEmailInvalid // retorna Email vazio e erro | returns empty Email and error
	}

	if domain == "" { // 4 - Tiver domínio após o @ | must have domain after @
		return "", ErrEmailInvalid // retorna Email vazio e erro | returns empty Email and error
	}

	if !strings.Contains(domain, ".") { // 5 - Domínio contiver pelo menos um ponto (.) | domain must contain at least one dot (.)
		//!strings.Contains significa "não contém". Ou seja, a condição verifica se o domínio NÃO contém um ponto (.)
		return "", ErrEmailInvalid // retorna Email vazio e erro | returns empty Email and error
	}

	return Email(strings.ToLower(email)), nil // retorna email validado, normalizado e em minúsculas como tipo Email | returns validated, normalized and lowercase email as Email type
}



