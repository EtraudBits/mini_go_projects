package client

// Client representa um cliente com dados normalizados | Client represents a client with normalized data.
// Usa tipos nomeados para garantir type safety e auto-documentação | Uses named types to ensure type safety and self-documentation
// Em sistemas reais, essa struct pode ser persistida em banco | In real systems, this struct can be persisted in a database.
// ou enviada para outros serviços | or sent to other services.

type Client struct { // cria uma struct para organizar os dados do cliente | creates a struct to organize client data.
	Name  Name  // Nome normalizado com capitalização correta | Normalized name with correct capitalization
	CPF   CPF   // CPF validado e normalizado (apenas dígitos) | Validated and normalized CPF (digits only)
	Phone Phone // Telefone normalizado | Normalized phone
	Email Email // Email validado | Validated email
}

// aqui no model.go não decide regras | here in model.go it doesn't decide rules.
// ele apenas representa o dominio | it just represents the domain.
// as regras de negocio ficam em outros pacotes | business rules are in other packages.