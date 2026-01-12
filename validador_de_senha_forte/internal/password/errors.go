/* Regra de negócio (contrato):

Uma senha é válida se:
1 - tiver no mínimo 8 caracteres
2 - contiver letra maiúscula
3 - contiver letra minúscula
4 - contiver número
5 - contiver caractere especial */

package password

import "errors"

// erros de domínio: representam falhas de regra de negócio
var (
	ErrTooShort = errors.New("senha deve ter pelo menos 8 caracteres")
	ErrNoUppercase = errors.New("senha deve conter pelo menos uma letra maiúscula")
	ErrNoLowercase = errors.New("senha deve conter pelo menos uma letra minúscula")
	ErrNoNumber = errors.New("senha deve conter pelo menos um número")
	ErrNoSpecialChar = errors.New("senha deve conter pelo menos um caractere especial")
)