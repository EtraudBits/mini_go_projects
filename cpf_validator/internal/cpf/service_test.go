package cpf

import "testing" // ativa o runner de testes do go / activates go's test runner

func TestNewCPF(t *testing.T) { //*testing.T -> Ponteiro para o tste em execução / Pointer to the test in execution
	_, err := NewCPF("529.982.247-25") // CPF válido / valid CPF
	if err != nil {
		t.Fatalf("cpf dveri ser valido, error: %v", err) // falha o teste / fails the test
	}
}

func TestNewCPF_IvalidRepeated(t *testing.T) {
	_, err := NewCPF("11111111111") // CPF invalido / ivalid cpf
	if err == nil { // se não der erro / if it doesn't give an error
		t.Error("era esperado erro para cpf com digitos repetidos") // falha o teste / fails the test
	}
}

