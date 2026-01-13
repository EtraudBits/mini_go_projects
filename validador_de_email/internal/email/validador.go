package email

import "strings" // "strings" é um pacote da biblioteca padrão do Go que fornece funções para manipulação de strings
// Precisamos dividir (Split), contar (Count), verificar conteúdo (Contains), entre outras operações com strings

// variavel que estar no diretorio errors.go com os erros definidos

/* 	1- ErrEmptyEmail = errors.New("email não pode ser vazio") // "email cannot be empty"
2- ErrIvalidAtSymbol = errors.New("email deve conter apenas um @") // "email must contain only one @")
3- ErrInvalidLocalPart = errors.New("email deve conter texto antes do @") // "email must contain text before the @")
4 - ErrIvalidDomain = errors.New("email deve conter domínio válido") // "email must contain a valid domain"
5 - ErrIvalidDotDomain = errors.New("domínio deve conter pelo menos um ponto") // "domain must contain at least one dot"
*/

// lógica do mini-projeto de validação de email (o cérebro do sistema)
func Validate(email string) error { 
	if email == "" { // 1 - Não estiver vazio
		return ErrEmptyEmail // retorna o erro definido em errors.go
	}

	if strings.Count(email, "@") != 1 { // 2 - Contiver exatamente um @ | strings.Count(email, "@") != 1 -> significa que o email deve conter exatamente um caractere '@'}
		return ErrIvalidAtSymbol // retorna o erro definido em errors.go
	}

	parts := strings.Split(email, "@") // divide o email em duas partes: antes e depois do '@'
	local := parts[0] // parte antes do '@'
	domain := parts[1] // parte depois do '@'

	if local == "" { // 3 - Tiver texto antes do @
		return ErrInvalidLocalPart // retorna o erro definido em errors.go
	}

	if domain == "" { // 4 - Tiver domínio após o @
		return ErrIvalidDomain // retorna o erro definido em errors.go
	}

	if !strings.Contains(domain, ".") { // 5 - Domínio contiver pelo menos um ponto (.) | strings.Contains(domain, ".") -> verifica se o domínio contém pelo menos um ponto (.)
	//!strings.Contains significa "não contém". Ou seja, a condição verifica se o domínio NÃO contém um ponto (.).
		return ErrIvalidDotDomain // retorna o erro definido em errors.go
	}

	return nil // se todas as validações passarem, retorna nil (sem erro)

}


