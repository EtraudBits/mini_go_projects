package main

import (
	"cpf_validator/internal/cpf" // importa o pacote cpf / imports the cpf package
	"fmt"
)

func main() {
	c, err := cpf.NewCPF("05195805407") // cria um novo cpf / creates a new cpf
	if err != nil {
		fmt.Println("CPF inválido:", err) // cpf inválido / invalid cpf
		return
	}
	fmt.Println("CPF válido:", c) // cpf válido / valid cpf
}