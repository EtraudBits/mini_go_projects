package client

// Client representa um cliente com dados normalizados.
// Usa tipos nomeados para garantir type safety e auto-documentação.
type Client struct {
	Name  Name  // Nome normalizado com capitalização correta
	CPF   CPF   // CPF validado e normalizado (apenas dígitos)
	Phone Phone // Telefone normalizado
	Email Email // Email validado
}