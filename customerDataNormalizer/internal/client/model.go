package client

// client represeta um cliente com dados normalizados | client represent a client with normalized data.
// Em sistemas reais, essa struct pode ser persistida em banco | In real systems, this struct can be persisted in a database.
// ou enviada para outros serciços | or sent to other services.

type Client struct { // cria uma struct para organizar os dados do cliente | creates a struct to organize client data.
	Name string
	CPF string
	Phone string
	Email string
}

// aqui no model.go não decide regras | here in model.go it doesn't decide rules.
// ele apenas representa o dominio | it just represents the domain.
// as regras de negocio ficam em outros pacotes | business rules are in other packages.