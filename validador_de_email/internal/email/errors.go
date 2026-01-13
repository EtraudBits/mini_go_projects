package email

import "errors"

var (
	ErrEmptyEmail = errors.New("email não pode ser vazio") // "email cannot be empty"
	ErrIvalidAtSymbol = errors.New("email deve conter apenas um @") // "email must contain only one @")
	ErrInvalidLocalPart = errors.New("email deve conter texto antes do @") // "email must contain text before the @")
	ErrIvalidDomain = errors.New("email deve conter domínio válido") // "email must contain a valid domain"
	ErrIvalidDotDomain = errors.New("domínio deve conter pelo menos um ponto") // "domain must contain at least one dot"
)

/*Um email será considerado válido se:

1 - Não estiver vazio

2 - Contiver exatamente um @

3 - Tiver texto antes do @

4 - Tiver domínio após o @

5 - Domínio contiver pelo menos um ponto (.)

Não começar nem terminar com @ ou .

Por que assim?
Erros são variáveis reutilizáveis
Nome começa com Err → convenção Go
Mensagem clara → reduz dúvida de quem usa a função

*/