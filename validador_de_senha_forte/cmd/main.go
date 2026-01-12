package main

import (
	"fmt"

	"validador_de_senha_forte/internal/password" // importa o pacote de validação de senha do go.mod
)

func main() {
	senha := "1enhaaaaa#Y" // Exemplo de senha a ser validada

	err := password.Validate((senha)) // password = pacate importado e Validate = função de validação

	if err != nil { // Se houver um erro, a senha é inválida (Go exige tratamento explícito de erros)
		fmt.Println("Erro", err.Error()) // Exibe a mensagem de erro específica
		return // Encerra a execução
	}

	fmt.Println("Senha válida!") // Senha passou por todas as validações


}