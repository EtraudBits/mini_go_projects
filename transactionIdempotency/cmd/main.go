package main

import (
	"fmt"
	"transactionIdempotency/internal/transaction"
)

// Ponto de entrada do aplicativo | Application entry point
func main() {
	repo := transaction.NewRepository()
	service := transaction.NewService(repo)

	t := transaction.Transaction{
		ID: "tx84",
		Customer: "Duarte",
		Amount: 1000000.00,
	}

	err := service.Process(t)
	if err != nil {
		fmt.Printf("❌ Erro: %s\n", err)
		fmt.Printf("   A transação com ID '%s' já foi processada anteriormente.\n", t.ID)
		return
	}

	fmt.Println("✅ Transação processada com sucesso!")
	fmt.Printf("   ID: %s\n", t.ID)
	fmt.Printf("   Cliente: %s\n", t.Customer)
	fmt.Printf("   Valor: R$ %.2f\n", t.Amount)
}

