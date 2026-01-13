package email

import "testing" // importa o pacote de teste do Go

// Testes automatizados para a função Validate
func TestValidate(t *testing.T) {
	err := Validate("du84arte@gmail.com") // email válido

	if err != nil { // se houver um erro
		t.Errorf("esperava email válido, mas recebeu erro: %#v", err) // falha no teste -> %#v imprime o valor do erro
	}
}

// testa todos os erros 

// Testa o erro de email vazio
func TestEmpty(t *testing.T) {
	err := Validate("") // email vazio

	if err != ErrEmptyEmail { // verifica se o erro retornado é o esperado
		t.Errorf("esperava ErrEmptyEmail, mas recebeu: %#v", err) // falha no teste
	}
}

// Testa o erro de email sem '@'
func TextIvalidAtSymbol(t *testing.T) {
	err := Validate("du84artegmail.com") // email sem '@'

	if err != ErrIvalidAtSymbol { // verifica se o erro retornado é o esperado
		t.Errorf("esperava ErrIvalidAtSymbol, mas recebeu: %#v", err) // falha no teste
	}
}

// testa o erro de email sem parte local
func TestInvalidLocalPart(t *testing.T) {
	err := Validate("@gmail.com") // email sem parte local

	if err != ErrInvalidLocalPart { // verifica se o erro retornado é o esperado
		t.Errorf("esperava ErrInvalidLocalPart, mas recebeu: %#v", err) // falha no teste
	}
}

// testa o erro de email sem domínio
func TestIvalidDomain(t *testing.T) {
	err := Validate("du84arte@") // email sem domínio

	if err != ErrIvalidDomain { // verifica se o erro retornado é o esperado
		t.Errorf("esperava ErrIvalidDomain, mas recebeu: %#v", err) // falha no teste
	}
}

// testa o erro de domínio sem ponto
func TestIvalidDotDomain(t *testing.T) {
	err := Validate("du84arte@gmailcom") // email com domínio sem ponto

	if err != ErrIvalidDotDomain { // verifica se o erro retornado é o esperado
		t.Errorf("esperava ErrIvalidDotDomain, mas recebeu: %#v", err) // falha no teste
	}
}	