package main

import (
	"fmt"
	"validador_de_telefone/internal/phone" // importa o pacote phone do validador_de_telefone/internal/phone - imports the phone package from validador_de_telefone/internal/phone
)

func main() {

	phoneInput := "(83) 91111-1111" // exemplo de telefone para validar - example of phone to validate

	err := phone.ValidadorPhone(phoneInput) // chama a função Pública ValidadorPhone do pacote phone e verifica a variavel phoneInput - calls the Public function ValidadorPhone from phone package and checks the variable phoneInput

	if err != nil { // se não for nulo (nil) = if not null (nil) 
		fmt.Println("❌ Erros ao validar telefone:") // imprime o cabeçalho dos erros  (x simples = \u2716 codigo unicode) - prints the error header
		fmt.Println(err) // imprime todos os erros na tela - prints all errors on the screen
		return // sai da função main - exits the main function
	}

	fmt.Println("✅ Telefone VÁLIDO!") // se não houver erro, imprime Telefone VÁLIDO! (check verde = \u2705 codigo unicode) - if there is no error, prints valid phone!
} // fim da função main - end of main function


// main só executa a lógica de chamar o validador e trata o erro, nada mais - main only executes the logic of calling the validator and handling the error, nothing more.