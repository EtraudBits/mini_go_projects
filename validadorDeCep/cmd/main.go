package main

import (
	"fmt"
	"validadorDeCep/internal/cep" // importa o pacote cep do projeto
)

func main() {
	cepValue := "59360000" // input pelo o usuario

	errors := cep.Validate(cepValue) // chama a função Validate do pacote cep para validar o input do usuario

	if len(errors) > 0 { // se houver erros na validação -> se houver mais de zero erros
		fmt.Println("CEP inválido. Erros encontrados:") // imprime a msg de erro
		// percorre a lista de erros e imprime cada um
		for _, err := range errors {
			fmt.Println("-", err.Error()) // err.Error() converte o erro para string
		}
		return // sai do programa
	}

	fmt.Println("CEP válido!") // se não houver erros, imprime que o CEP é válido
}