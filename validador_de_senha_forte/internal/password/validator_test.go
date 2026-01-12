package password

import "testing" // pacote de teste do Go

func TestValidate(t *testing.T) { // t = contexto do teste e * = ponteiro (permite modificar o estado do teste)
	err := Validate("Senha84&") // senha válida - regra de negocio atendida
	if err != nil { // se houver erro, a senha é inválida
		t.Errorf("esperava senha válida, mas recebeu erro: %v", err)
	}
}

// testes para cada regra de negócio

// Testa a regra 1 - mínimo de 8 caracteres
func TestValidate_TooShort(t *testing.T) {
	err := Validate ("Bb4*") // senha inválida - menos de 8 caracteres
	if err != ErrTooShort {
		t.Errorf("esperava ErrTooShort, mas recebeu %v", err)
	}
}

// Testa a regra 2 - letra maiúscula
func TestValidate_NoUppercase(t *testing.T) {
	err := Validate ("abc*1246") // senha inválida - sem letra maiúscula
	if err != ErrNoUppercase {
		t.Errorf("esperava ErrNoUppercase, mas recebeu %v", err)
	}
}

// Testa a regra 3 - letra minúscula
func TestValidate_NoLowercase(t *testing.T) {
	err := Validate ("ABCD1234*") // senha inválida - sem letra minúscula
	if err != ErrNoLowercase {
		t.Errorf("esperava ErrNoLowercase, mas recebeu %v", err)
	}
}

// Testa a regra 4 - número
func TestValidate_NoNumber(t *testing.T) {
	err := Validate ("Abcdefg*") // senha inválida - sem número
	if err != ErrNoNumber {
		t.Errorf("esperava ErrNoNumber, mas recebeu %v", err)
	}
}

// Testa a regra 5 - caractere especial
func TestValidate_NoSpecialChar(t *testing.T) {
	err := Validate ("Abcd1234") // senha inválida - sem caractere especial
	if err != ErrNoSpecialChar {
		t.Errorf("esperava ErrNoSpecialChar, mas recebeu %v", err)
	}
}