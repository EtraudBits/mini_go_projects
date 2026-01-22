package main

import (
	"customerDataNormalizer/internal/client"
	"fmt"
)

func main() {
	fmt.Println("=== Customer Data Normalizer ===")
	fmt.Println() // linha em branco | blank line

	// Dados crus do cliente | Raw client data
	rawName := "  DuArTE RodRIGo SanTOS DE oliveIRA  "
	rawCPF := "123.456.789-09" // CPF válido para teste | Valid CPF for testing
	rawPhone := "(84) 92345-6789"
	rawEmail := "DU84ARTE@GMAIL.COM"

	// Normaliza os dados | Normalize the data
	normalizedClient, errs := client.Normalize(rawName, rawCPF, rawPhone, rawEmail)

	// Se houver erros, exibe | If there are errors, display them
	if len(errs) > 0 {
		fmt.Println("❌ Erros encontrados:")
		for _, err := range errs {
			fmt.Printf("  - %v\n", err)
		}
		return
	}

	// Exibe dados normalizados | Display normalized data
	fmt.Println("✅ Cliente normalizado com sucesso!")
	fmt.Printf("\nNome: %s\n", normalizedClient.Name)
	fmt.Printf("CPF: %s\n", normalizedClient.CPF)
	fmt.Printf("Telefone: %s\n", normalizedClient.Phone)
	fmt.Printf("Email: %s\n", normalizedClient.Email)
}
