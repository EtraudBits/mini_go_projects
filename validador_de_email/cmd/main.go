package main

import (
	"fmt"
	"validador_de_email/internal/email" // importa o pacote email do projeto
)

func main() {
	emailInput := "du84arte@gmail.com" // email a ser validado

	// chama a regra de negocio para validar o email
	err := email.Validate(emailInput) // chama a função Validate do pacote email para validar o email
	//email.Validade(emailIput) -> email = pacote | Validate = função | emailInput = argumento trazido pelo o usuário

	if err != nil { // se houver um erro na validação
		fmt.Println("Email inválido:", err) // imprime o erro retornado pela função Validate
		return
	}

	fmt.Println("Email válido!") // se não houver erro, imprime que o email é válido
}	